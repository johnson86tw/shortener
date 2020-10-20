package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	app := echo.New()

	// middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// service
	accountRepo := accountRepo.NewRepository(pgConn)
	as := accountService.NewAccountService(accountRepo)
	redirectRepo := redirectRepo.NewRepository(pgConn)
	rs := redirectService.NewRedirectService(redirectRepo)
	api.NewAccountHandler(app, as, j)
	api.NewRedirectHandler(app, rs)

	// api
	auth := app.Group("/auth")
	auth.Use(j.AuthRequired)

	{
		auth.GET("/profile", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success",
			})
		})
	}

	// userURLRepo := userURLRepo.NewRepository(pgConn)
	// userURLService := userURLService.NewUserURLService(userURLRepo)
	app.Logger.Fatal(app.Start(serverAddr))
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
