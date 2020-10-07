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

	// config := viper.GetStringMapString("pg")
	// dsn := []string{}
	// for key, val := range config {
	// 	s := key + "=" + val
	// 	dsn = append(dsn, s)
	// }

	// pgConn, err := pgx.Connect(context.Background(), strings.Join(dsn, " "))
	// if err != nil {
	// 	log.Panic(err)
	// }

	// defer pgConn.Close(context.Background())

	// rows, err := pgConn.Query(context.Background(), "SELECT * FROM url_table;")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()

	// rdrts := []*domain.Redirect{}

	// for rows.Next() {
	// 	var id int
	// 	var createdAt interface{}

	// 	rdrt := &domain.Redirect{}
	// 	err := rows.Scan(&id, &rdrt.URL, &rdrt.Code, &createdAt)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}

	// 	rdrts = append(rdrts, rdrt)
	// }

	g := gin.Default()

	redisRepo := _redisRepo.NewRedisRedirectRepository(rdb)
	redirectService := _redirectService.NewRedirectService(redisRepo)
	_redirectHttpDelivery.NewRedirectHandler(g, redirectService)

	g.Run(serverAddr)
}
