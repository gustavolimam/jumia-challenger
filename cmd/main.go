package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/jumia-challenger/internal/api"
	"github.com/jumia-challenger/internal/environment"
	"github.com/jumia-challenger/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	godotenv.Load()

	// Verify if all var envs was filled
	environment.CheckEnvVars()

	log.Info().Msg("Starting the application...")

	go loadWebapp()

	e := loadAPI()

	a := api.Start(e)
	a.RegisterRoutes()

	log.Info().Msgf("HTTP Server on port %v", os.Getenv(environment.Port))
	e.Logger.Fatal(e.Start(":" + os.Getenv(environment.Port)))
}

func loadAPI() *echo.Echo {
	// create an instance of Echo and update the default config
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:" + os.Getenv(environment.WebPort), "http://localhost:3000"},
		AllowMethods: []string{"*"},
	}))

	return e
}

func loadWebapp() {
	web.Run()
}
