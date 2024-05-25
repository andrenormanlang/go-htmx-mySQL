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


	flag.Parse()

	var app_settings common.AppSettings
	

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

	var defaultPort = 8081 // Set your default port here

if app_settings.WebserverPort == 0 {
    app_settings.WebserverPort = defaultPort
}

	err = r.Run(fmt.Sprintf(":%d", app_settings.WebserverPort))
	if err != nil {
		log.Error().Msgf("could not run app: %v", err)
		os.Exit(-1)
	}
}
