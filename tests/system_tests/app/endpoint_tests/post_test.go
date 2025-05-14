package endpoint_tests

import (
	_ "database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrenormanlang/go-htmx-mySQL/app"
	"github.com/andrenormanlang/go-htmx-mySQL/tests/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestPostExists(t *testing.T) {
	// Setup test database with migrations
	app_settings, database, err := helpers.SetupTestDatabase()
	require.Nil(t, err)

	r := app.SetupRoutes(app_settings, database)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/post/1", nil)
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}
