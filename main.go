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
	_accountService "github.com/chnejohnson/shortener/account/service"
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
	rs := _redirectService.NewRedirectService(pgRepo)
	_redirectHttpDelivery.NewRedirectHandler(g, rs)

	accountRepo := _accountRepo.NewRepository(pgConn)
	as := _accountService.NewAccountService(accountRepo)

	// try
	acc := &domain.Account{}
	acc.Name = "Howard"
	acc.Email = "howard@gmail.com"
	acc.Password = "23"

	err = as.Create(acc)
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("Succeed to sign up")
	}

	err = as.Login("howard@gmail.com", "23")
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("Succeed to login")
	}

	g.Run(serverAddr)
}
