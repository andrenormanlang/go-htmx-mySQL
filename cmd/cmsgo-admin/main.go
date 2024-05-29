package main

import (
	"flag"
	"fmt"
	"os"

	admin_app "github.com/andrenormanlang/admin-app"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func main() {
	// sets zerolog as the main logger
	// in this APP
	common.SetupLogger()

	config_toml := flag.String("config", "", "path to config toml file")
	flag.Parse()

	var app_settings common.AppSettings
	if *config_toml != "" {
		log.Info().Msgf("reading config file %s", *config_toml)
		settings, err := common.ReadConfigToml(*config_toml)
		if err != nil {
			log.Error().Msgf("could not read toml: %v", err)
			os.Exit(-2)
		}
		app_settings = settings
	} else {
		log.Info().Msgf("no config passed, reading environment variable settings")
		settings, err := common.LoadSettings()
		if err != nil {
			log.Fatal().Msgf("could not load app settings: %v", err)
			os.Exit(-1)
		}
		app_settings = settings
	}

	database, err := database.MakeSqlConnection(
		app_settings.DatabaseUser,
		app_settings.DatabasePassword,
		app_settings.DatabaseAddress,
		app_settings.DatabasePort,
		app_settings.DatabaseName,
	)
	if err != nil {
		log.Fatal().Msgf("could not create database: %v", err)
		os.Exit(-1)
	}

	r := admin_app.SetupRoutes(app_settings, database)
	err = r.Run(fmt.Sprintf(":%d", app_settings.WebserverPort))
	if err != nil {
		log.Error().Msgf("could not run app: %v", err)
		os.Exit(-1)
	}
}
