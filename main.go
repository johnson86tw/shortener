package main

import (
	"context"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/chnejohnson/shortener/domain"
	_redirectHttpDelivery "github.com/chnejohnson/shortener/redirect/delivery/http"
	_redirectService "github.com/chnejohnson/shortener/redirect/service"

	_accountRepo "github.com/chnejohnson/shortener/account/repository/postgres"
	_redirectRepo "github.com/chnejohnson/shortener/redirect/repository/postgres"
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

	// Setting gin and middleware
	g := gin.Default()

	// compose redis
	// redisRepo := _redisRepo.NewRedisRedirectRepository(rdb)
	// rdbService := _redirectService.NewRedirectService(redisRepo)
	// _redirectHttpDelivery.NewRedirectHandler(g, rdbService)

	// compose postgres
	pgRepo := _redirectRepo.NewRepository(pgConn)
	pgService := _redirectService.NewRedirectService(pgRepo)
	_redirectHttpDelivery.NewRedirectHandler(g, pgService)

	accountRepo := _accountRepo.NewRepository(pgConn)

	account := &domain.Account{}
	account.Email = "grace@gmail.com"
	account.Password = "1234"
	account.Name = "grace"

	err = accountRepo.Create(account)
	if err != nil {
		logrus.Error(err)
	}

	password, err := accountRepo.Find("grace@gmail.com")
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.WithField("password", password).Info("Get Password!")
	}

	g.Run(serverAddr)
}
