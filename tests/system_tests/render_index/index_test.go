package render_index

import (
	"context"
	_"database/sql"
	"embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/andrenormanlang/app"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"
)

//go:generate ../../../migrations ./migrations

//go:embed migrations/*.sql
var embedMigrations embed.FS

var (
	dbName    = "cmsgo"
	address   = "localhost"
	port      = 3336
)

var app_settings = common.AppSettings{
	DatabaseAddress:  "localhost",
	DatabasePort:     3336,
	DatabaseUser:     "root",
	DatabasePassword: "",
	DatabaseName:     "cmsgo",
	WebserverPort:    8080,
}

func runDatabaseServer() {
	pro := createTestDatabase()
	engine := sqle.NewDefault(pro)
	engine.Analyzer.Catalog.MySQLDb.AddRootAccount()

	session := memory.NewSession(sql.NewBaseSession(), pro)
	ctx := sql.NewContext(context.Background(), sql.WithSession(session))
	ctx.SetCurrentDatabase("test")

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}
	s, err := server.NewServer(config, engine, memory.NewSessionBuilder(pro), nil)
	if err != nil {
		panic(err)
	}
	if err = s.Start(); err != nil {
		panic(err)
	}
}

func createTestDatabase() *memory.DbProvider {
	db := memory.NewDatabase(dbName)
	db.BaseDatabase.EnablePrimaryKeyIndexes()

	pro := memory.NewDBProvider(db)
	return pro
}

func waitForDB(app_settings common.AppSettings) (database.SqlDatabase, error) {
	for range (400) {
		database, err := database.MakeSqlConnection(
			app_settings.DatabaseUser,
			app_settings.DatabasePassword,
			app_settings.DatabaseAddress,
			app_settings.DatabasePort,
			app_settings.DatabaseName)
		if err == nil {
			return database, nil
		}

		time.Sleep(25 * time.Millisecond)
	}
	return database.SqlDatabase{}, fmt.Errorf("Database did not start")
}

func TestIndexPagePing(t *testing.T) {

	go runDatabaseServer()

	database, err := waitForDB(app_settings)
	require.Nil(t, err)
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		require.Nil(t, err)
	}

	if err := goose.Up(database.Connection, "migrations"); err != nil {
		panic(err)
	}

	r := app.SetupRoutes(app_settings, database)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestIndexPagePostExists(t *testing.T) {

	go runDatabaseServer()

	database, err := waitForDB(app_settings)
	require.Nil(t, err)
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("mysql"); err != nil {
		require.Nil(t, err)
	}

	if err := goose.Up(database.Connection, "migrations"); err != nil {
		panic(err)
	}

	r := app.SetupRoutes(app_settings, database)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	require.Contains(t, w.Body.String(), "My Very First Post")
}

