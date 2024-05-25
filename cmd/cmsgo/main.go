package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrenormanlang/app"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func main() {
	config_toml := flag.String("config", "", "path to the config to be used")
	flag.Parse()

	var app_settings common.AppSettings
	if (*config_toml) != "" {
		log.Info().Msgf("reading config file %s", *config_toml)
		settings, err := common.ReadConfigToml(*config_toml)
		if err != nil {
			log.Error().Msgf("could not read config file: %v", err)
			os.Exit(-1)
		}

		app_settings = settings
	} else {
		log.Info().Msgf("no config file, reading environment variables")
		settings, err := common.LoadSettings()
		if err != nil {
			log.Error().Msgf("could not load settings: %v", err)
			os.Exit(-1)
		}
		app_settings = settings
	}

	db_connection, err := database.MakeSqlConnection(
		app_settings.DatabaseUser,
		app_settings.DatabasePassword,
		app_settings.DatabaseAddress,
		app_settings.DatabasePort,
		app_settings.DatabaseName,
	)
	if err != nil {
		log.Error().Msgf("could not create database connection: %v", err)
		os.Exit(-1)
	}

	var defaultPort = 8080 // Set your default port here

if app_settings.WebserverPort == 0 {
    app_settings.WebserverPort = defaultPort
}

	r := app.SetupRoutes(app_settings, db_connection)
	err = r.Run(fmt.Sprintf(":%d", app_settings.WebserverPort))
	if err != nil {
		log.Error().Msgf("could not run app: %v", err)
		os.Exit(-1)
	}
}
