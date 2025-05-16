package endpoint_tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrenormanlang/go-htmx-mySQL/app"
	"github.com/andrenormanlang/go-htmx-mySQL/common"
	"github.com/andrenormanlang/go-htmx-mySQL/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestIndexSuccess(t *testing.T) {
	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "cms-and-go",
		WebserverPort:    8080,
	}

	database_mock := mocks.DatabaseMock{
		GetPostsHandler: func(limit int, offset int) ([]common.Post, error) {
			return []common.Post{
				{
					Title:   "TestPost",
					Content: "TestContent",
					Excerpt: "TestExcerpt",
					Id:      0,
				},
			}, nil
		},
	}

	r := app.SetupRoutes(app_settings, database_mock)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/page/0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "TestPost")
	assert.Contains(t, w.Body.String(), "TestExcerpt")
}

func TestIndexFailToGetPosts(t *testing.T) {
	app_settings := common.AppSettings{
		DatabaseAddress:  "localhost",
		DatabasePort:     3006,
		DatabaseUser:     "root",
		DatabasePassword: "root",
		DatabaseName:     "cms-and-go",
		WebserverPort:    8080,
	}

	database_mock := mocks.DatabaseMock{
		GetPostsHandler: func(limit int, offset int) ([]common.Post, error) {
			return []common.Post{}, fmt.Errorf("invalid")
		},
	}

	r := app.SetupRoutes(app_settings, database_mock)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
