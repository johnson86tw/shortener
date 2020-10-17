package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_accountRepo "github.com/chnejohnson/shortener/service/account/repository/postgres"
	_accountService "github.com/chnejohnson/shortener/service/account/service"

	api "github.com/chnejohnson/shortener/api"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		gin.SetMode("debug")
		log.Println("Service run on develop mode")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	redisAddr := viper.GetString("redis.address")
	serverAddr := viper.GetString("server.address")
	pgConfig := viper.GetStringMapString("pg")
	jwtSecret := viper.GetString("jwt.secret")

	j := &api.JWT{JWTSecret: []byte(jwtSecret)}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping().Result(); err != nil {
		logrus.Warn("Cannot connect to redis")
	}

	dsn := []string{}
	for key, val := range pgConfig {
		s := key + "=" + val
		dsn = append(dsn, s)
	}

	pgConn, err := pgx.Connect(context.Background(), strings.Join(dsn, " "))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := pgConn.Close(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	accountRepo := _accountRepo.NewRepository(pgConn)
	as := _accountService.NewAccountService(accountRepo)

	// pgRepo := _redirectRepo.NewRepository(pgConn)
	// redirectService := _redirectService.NewRedirectService(pgRepo)

	// api
	engine := gin.Default()
	api.NewAccountHandler(engine, as, j)

	authorized := engine.Group("/auth")
	authorized.Use(j.AuthRequired)

	{
		authorized.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	// compose redis
	// redisRepo := _redisRepo.NewRedisRedirectRepository(rdb)
	// rdbService := _redirectService.NewRedirectService(redisRepo)
	// _redirectHttpDelivery.NewRedirectHandler(g, rdbService)

	engine.Run(serverAddr)
}
