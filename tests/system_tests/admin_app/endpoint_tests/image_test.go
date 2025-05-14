package admin_endpoint_tests

import (
	_ "database/sql"
	"encoding/json"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	admin_app "github.com/andrenormanlang/go-htmx-mySQL/admin-app"
	"github.com/andrenormanlang/go-htmx-mySQL/tests/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/test-go/testify/require"
)

func TestImageUpload(t *testing.T) {
	// Setup test database with migrations
	app_settings, database, err := helpers.SetupTestDatabase()
	require.Nil(t, err)
	defer func() {
		err := helpers.CleanupTestDatabase(app_settings)
		require.Nil(t, err)
	}()

	// Multipart image form creation
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	go func() {
		defer writer.Close()

		// Create the image part
		image_part, err := helpers.CreateFormImagePart(writer, "file", "test.png", "image/png")
		require.Nil(t, err)
		err = png.Encode(image_part, helpers.CreateImage())
		require.Nil(t, err)

		// Create the alt part
		text_part, err := helpers.CreateTextFormHeader(writer, "alt")
		require.Nil(t, err)
		_, err = text_part.Write([]byte("test alt"))
		require.Nil(t, err)
	}()

	// Execute multiform request
	post_recorder := httptest.NewRecorder()
	r := admin_app.SetupRoutes(app_settings, database)
	require.Nil(t, err)

	req, _ := http.NewRequest("POST", "/images", pr)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	r.ServeHTTP(post_recorder, req)

	require.Equal(t, http.StatusOK, post_recorder.Code)

	// Make sure returned an ID
	var image_id_response admin_app.ImageIdResponse
	err = json.Unmarshal(post_recorder.Body.Bytes(), &image_id_response)
	require.Nil(t, err)
}
