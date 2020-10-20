package main

import (
	"context"
	"log"
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
	userURLRepo "github.com/chnejohnson/shortener/service/user_url/repository/postgres"
	userURLService "github.com/chnejohnson/shortener/service/user_url/service"
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

	// web framework
	app := echo.New()

	// middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	// basic
	accountRepo := accountRepo.NewRepository(pgConn)
	as := accountService.NewAccountService(accountRepo)
	redirectRepo := redirectRepo.NewRepository(pgConn)
	rs := redirectService.NewRedirectService(redirectRepo)

	api.NewAccountHandler(app, as, j)
	api.NewRedirectHandler(app, rs)

	// auth
	auth := app.Group("/auth")
	auth.Use(j.AuthRequired)
	{
		userURLRepo := userURLRepo.NewRepository(pgConn)
		us := userURLService.NewUserURLService(userURLRepo)
		api.NewUserURLHandler(auth, us)
	}

	app.Logger.Fatal(app.Start(serverAddr))
}
