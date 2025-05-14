package main

import (
	"flag"
	"fmt"
	"os"
	// "strconv"

	admin_app "github.com/andrenormanlang/go-htmx-mySQL/admin-app"
	"github.com/andrenormanlang/go-htmx-mySQL/common"
	"github.com/andrenormanlang/go-htmx-mySQL/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func main() {
	// sets zerolog as the main logger
	// in this APP
	common.SetupLogger()

	config_toml := flag.String("config", os.Getenv("CMSGO_CONFIG"), "path to config toml file")
	flag.Parse()

	var app_settings common.AppSettings
	if *config_toml == "" {
		log.Error().Msgf("config not specified (--config)")
		os.Exit(-1)
	}

	log.Info().Msgf("reading config file %s", *config_toml)
	settings, err := common.ReadConfigToml(*config_toml)
	if err != nil {
		log.Error().Msgf("could not read toml: %v", err)
		os.Exit(-2)
	}
	app_settings = settings

	database, err := database.MakeMysqlConnection(
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

	// Use PORT environment variable if set, otherwise use WebserverPort from config
	// port := app_settings.WebserverPort
	// if portEnv := os.Getenv("PORT"); portEnv != "" {
	// 	if p, err := strconv.Atoi(portEnv); err == nil {
	// 		port = p
	// 		log.Info().Msgf("Using PORT environment variable: %d", port)
	// 	}
	// }

	// r := admin_app.SetupRoutes(app_settings, database)
	// err = r.Run(fmt.Sprintf(":%d", port))
	// if err != nil {
	// 	log.Fatal().Msgf("could not start server: %v", err)
	// 	os.Exit(-1)
	// }
	r := admin_app.SetupRoutes(app_settings, database)
	err = r.Run(fmt.Sprintf(":%d", app_settings.AdminPort))
	if err != nil {
		log.Error().Msgf("could not run app: %v", err)
		os.Exit(-1)
	}

	// package main

	// import (
	// 	"flag"
	// 	"fmt"
	// 	"os"

	// 	admin_app "github.com/andrenormanlang/go-htmx-mySQL/admin-app"
	// 	"github.com/andrenormanlang/go-htmx-mySQL/common"
	// 	"github.com/andrenormanlang/go-htmx-mySQL/database"
	// 	"github.com/andrenormanlang/go-htmx-mySQL/plugins"
	// 	_ "github.com/go-sql-driver/mysql"
	// 	"github.com/rs/zerolog/log"
	// 	lua "github.com/yuin/gopher-lua"
	// )

	// func loadShortcodeHandlers(shortcodes []common.Shortcode) (map[string]*lua.LState, error) {
	// 	shortcode_handlers := make(map[string]*lua.LState, 0)
	// 	for _, shortcode := range shortcodes {
	// 		// Read the LUA state
	// 		state := lua.NewState()
	// 		err := state.DoFile(shortcode.Plugin)
	// 		// TODO : check that the function HandleShortcode(args)
	// 		//        exists and returns the correct type
	// 		if err != nil {
	// 			return map[string]*lua.LState{}, fmt.Errorf("could not load shortcode %s: %v", shortcode.Name, err)
	// 		}
	// 		shortcode_handlers[shortcode.Name] = state
	// 	}

	// 	return shortcode_handlers, nil
	// }

	// func main() {
	// 	// sets zerolog as the main logger
	// 	// in this APP
	// 	common.SetupLogger()

	// 	config_toml := flag.String("config", os.Getenv("CMSGO_CONFIG"), "path to config toml file")
	// 	flag.Parse()

	// 	var app_settings common.AppSettings
	// 	if *config_toml == "" {
	// 		log.Error().Msgf("config not specified (--config)")
	// 		os.Exit(-1)
	// 	}

	// 	log.Info().Msgf("reading config file %s", *config_toml)
	// 	settings, err := common.ReadConfigToml(*config_toml)
	// 	if err != nil {
	// 		log.Error().Msgf("could not read toml: %v", err)
	// 		os.Exit(-2)
	// 	}
	// 	app_settings = settings

	// 	database, err := database.MakeSqlConnection(
	// 		app_settings.DatabaseUser,
	// 		app_settings.DatabasePassword,
	// 		app_settings.DatabaseAddress,
	// 		app_settings.DatabasePort,
	// 		app_settings.DatabaseName,
	// 	)
	// 	if err != nil {
	// 		log.Fatal().Msgf("could not create database: %v", err)
	// 		os.Exit(-1)
	// 	}

	// 	shortcode_handlers, err := loadShortcodeHandlers(app_settings.Shortcodes)
	// 	if err != nil {
	// 		log.Error().Msgf("%v", err)
	// 		os.Exit(-1)
	// 	}

	// 	// TODO: we probably want to refact loadShortcodeHandler
	// 	// TODO: into loadPluginHandlers instead
	// 	post_hook := plugins.PostHook{}
	// 	image_plugin := plugins.Plugin{
	// 		ScriptName: "img",
	// 		Id: "img-plugin",
	// 	}
	// 	post_hook.Register(image_plugin)

	// 	// img, _ := shortcode_handlers["img"]
	// 	hooks_map := map[string]plugins.Hook{
	// 		"add_post": post_hook,
	// 	}

	// 	r := admin_app.SetupRoutes(app_settings, database, shortcode_handlers, hooks_map)
	// 	err = r.Run(fmt.Sprintf(":%d", app_settings.WebserverPort))
	// 	if err != nil {
	// 		log.Error().Msgf("could not run app: %v", err)
	// 		os.Exit(-1)
	// 	}
	// }
}
