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

	accountRepo "github.com/chnejohnson/shortener/service/account/repository/postgres"
	accountService "github.com/chnejohnson/shortener/service/account/service"

	redirectRepo "github.com/chnejohnson/shortener/service/redirect/repository/postgres"
	redirectService "github.com/chnejohnson/shortener/service/redirect/service"

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

	engine := gin.Default()

	// service
	accountRepo := accountRepo.NewRepository(pgConn)
	as := accountService.NewAccountService(accountRepo)
	redirectRepo := redirectRepo.NewRepository(pgConn)
	rs := redirectService.NewRedirectService(redirectRepo)
	api.NewAccountHandler(engine, as, j)
	api.NewRedirectHandler(engine, rs)

	// api
	authorized := engine.Group("/login")
	authorized.Use(j.AuthRequired)

	{
		authorized.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}

	// userURLRepo := userURLRepo.NewRepository(pgConn)
	// userURLService := userURLService.NewUserURLService(userURLRepo)

	engine.Run(serverAddr)
}

// compose redis
// redisRepo := _redisRepo.NewRedisRedirectRepository(rdb)
// rdbService := _redirectService.NewRedirectService(redisRepo)
// _redirectHttpDelivery.NewRedirectHandler(g, rdbService)

// id, err := uuid.Parse("4425ff13-354f-4e45-897f-ac76476305d5")
// if err != nil {
// 	logrus.Error(err)
// }

// urls, err := userURLService.FetchAll(id)
// if err != nil {
// 	logrus.Error(err)
// }

// for _, u := range urls {
// 	fmt.Println(*u)
// }
