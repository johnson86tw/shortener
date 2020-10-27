package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	accountRepo "github.com/chnejohnson/shortener/service/account/repository/postgres"
	accountService "github.com/chnejohnson/shortener/service/account/service"

	redirectRepo "github.com/chnejohnson/shortener/service/redirect/repository/postgres"
	redirectService "github.com/chnejohnson/shortener/service/redirect/service"

	userURLRepo "github.com/chnejohnson/shortener/service/user_url/repository/postgres"
	userURLService "github.com/chnejohnson/shortener/service/user_url/service"

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
	serverAddr := viper.GetString("server.address")
	pgConfig := viper.GetStringMapString("pg")
	jwtSecret := viper.GetString("jwt.secret")

	j := &api.JWT{JWTSecret: []byte(jwtSecret)}

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
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// service
	accountRepo := accountRepo.NewRepository(pgConn)
	as := accountService.NewAccountService(accountRepo)
	userURLRepo := userURLRepo.NewRepository(pgConn)
	us := userURLService.NewUserURLService(userURLRepo)
	redirectRepo := redirectRepo.NewRepository(pgConn)
	rs := redirectService.NewRedirectService(redirectRepo, userURLRepo)

	// api
	app := e.Group("/api")

	api.NewAccountHandler(app, as, j)
	api.NewRedirectHandler(app, rs)

	// auth
	auth := app.Group("/auth")
	auth.Use(j.AuthRequired)
	{
		api.NewUserURLHandler(auth, us)
	}

	e.Logger.Fatal(e.Start(serverAddr))
}

func writeRoutesFile(app *echo.Echo) error {
	data, err := json.MarshalIndent(app.Routes(), "", "  ")
	if err != nil {
		return err
	}
	ioutil.WriteFile("routes.json", data, 0644)
	return nil
}
