package helpers

import (
	_ "database/sql"
	"embed"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"

	"github.com/andrenormanlang/go-htmx-mySQL/common"
	"github.com/andrenormanlang/go-htmx-mySQL/database"
)

//go:generate ../../../migrations ./migrations

//go:embed migrations/*.sql
var EmbedMigrations embed.FS

func WaitForDb(app_settings common.AppSettings) (database.SqlDatabase, error) {

	for range 400 {
		database, err := database.MakeMysqlConnection(app_settings.DatabaseUser, app_settings.DatabasePassword, app_settings.DatabaseAddress, app_settings.DatabasePort, app_settings.DatabaseName)

		if err == nil {
			return database, nil
		}

		time.Sleep(25 * time.Millisecond)
	}

	return database.SqlDatabase{}, fmt.Errorf("database did not start")
}

// SetupTestDatabase initializes a test database with all migrations applied.
// This should be used in all tests that need a database connection.
func SetupTestDatabase() (common.AppSettings, database.SqlDatabase, error) {
	app_settings := GetAppSettings()

	// First connect to MySQL without a specific database
	rootDb, err := database.MakeMysqlConnection(
		app_settings.DatabaseUser,
		app_settings.DatabasePassword,
		app_settings.DatabaseAddress,
		app_settings.DatabasePort,
		"",
	)
	if err != nil {
		return app_settings, database.SqlDatabase{}, fmt.Errorf("failed to connect to root database: %v", err)
	}
	defer rootDb.Connection.Close()

	// Drop the database if it exists and create it fresh
	_, err = rootDb.Connection.Exec("DROP DATABASE IF EXISTS `" + app_settings.DatabaseName + "`")
	if err != nil {
		return app_settings, database.SqlDatabase{}, fmt.Errorf("failed to drop database: %v", err)
	}

	_, err = rootDb.Connection.Exec("CREATE DATABASE `" + app_settings.DatabaseName + "`")
	if err != nil {
		return app_settings, database.SqlDatabase{}, fmt.Errorf("failed to create database: %v", err)
	}

	// Now connect to the new database
	database, err := WaitForDb(app_settings)
	if err != nil {
		return app_settings, database, fmt.Errorf("failed to connect to new database: %v", err)
	}

	// Initialize goose
	goose.SetBaseFS(EmbedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		return app_settings, database, fmt.Errorf("failed to set dialect: %v", err)
	}

	// Run migrations - let Goose handle table creation and versioning
	if err := goose.Up(database.Connection, "migrations"); err != nil {
		return app_settings, database, fmt.Errorf("failed to run migrations: %v", err)
	}

	return app_settings, database, nil
}

// CleanupTestDatabase drops the test database after tests complete
func CleanupTestDatabase(app_settings common.AppSettings) error {
	// Connect to MySQL without selecting a database
	db, err := database.MakeMysqlConnection(
		app_settings.DatabaseUser,
		app_settings.DatabasePassword,
		app_settings.DatabaseAddress,
		app_settings.DatabasePort,
		"",
	)
	if err != nil {
		return fmt.Errorf("failed to connect to database for cleanup: %v", err)
	}
	defer db.Connection.Close()

	// Drop the test database
	_, err = db.Connection.Exec("DROP DATABASE IF EXISTS `" + app_settings.DatabaseName + "`")
	if err != nil {
		return fmt.Errorf("failed to drop test database: %v", err)
	}

	return nil
}

// GetAppSettings gets the settings for the http servers
// taking into account a unique port. Very hacky way to
// get a unique port: manually have to pass a new number
// for every test...
// TODO : Find a way to assign a unique port at compile
//
//	time
func GetAppSettings() common.AppSettings {
	// Generate a unique database name using timestamp and a random suffix
	timestamp := time.Now().UnixNano()
	dbName := fmt.Sprintf("cms_and_go_test_%d", timestamp)

	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3306,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     dbName,
		WebserverPort:    8080,
		CacheEnabled:     false,
	}

	return app_settings
}
