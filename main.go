package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"

	_redirectHttpDelivery "github.com/chnejohnson/shortener/redirect/delivery/http"
	_redisRepo "github.com/chnejohnson/shortener/redirect/repository/redis"
	_redirectService "github.com/chnejohnson/shortener/redirect/service"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping().Result(); err != nil {
		panic(err)
	}

	g := gin.Default()

	redisRepo := _redisRepo.NewRedisRedirectRepository(rdb)
	redirectService := _redirectService.NewRedirectService(redisRepo)
	_redirectHttpDelivery.NewRedirectHandler(g, redirectService)

	g.Run(serverAddr)
}
