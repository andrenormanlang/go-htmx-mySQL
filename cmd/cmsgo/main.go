package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type AppSettings struct {
	databaseAddress  string
	databasePort     int
	databaseUser     string
	databasePassword string
}

type Post struct {
	Title   string
	Content string
}

type Database struct {
	address    string
	port       int
	user       string
	connection *sql.DB
}

func (db Database) getPosts() ([]Post, error) {
	rows, err := db.connection.Query("SELECT title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allPosts := []Post{}

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Title, &post.Content); err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}

func loadSettings() AppSettings {
	databaseAddress := os.Getenv("GOCMS_DATABASE_ADDRESS")
	if len(databaseAddress) == 0 {
		log.Fatalf("GOCMS_DATABASE_ADDRESS is not defined\n")
	}

	databaseUser := os.Getenv("GOCMS_DATABASE_USER")
	if len(databaseUser) == 0 {
		log.Fatalf("GOCMS_DATABASE_USER is not defined\n")
	}
	databasePassword := os.Getenv("GOCMS_DATABASE_PASSWORD")
	if len(databasePassword) == 0 {
		log.Fatalf("GOCMS_DATABASE_PASSWORD is not defined\n")
	}

	databasePortStr := os.Getenv("GOCMS_DATABASE_PORT")
	if len(databasePortStr) == 0 {
		log.Fatalf("GOCMS_DATABASE_PORT is not defined")
	}

	databasePort, err := strconv.Atoi(databasePortStr)
	if err != nil {
		log.Fatalf("GOCMS_DATABASE_PORT is not a valid integer: %v\n", err)
	}

	return AppSettings{
		databaseUser:     databaseUser,
		databaseAddress:  databaseAddress,
		databasePort:     databasePort,
		databasePassword: databasePassword,
	}
}

func makeHomeHandler(settings AppSettings, db Database) func(*gin.Context) {
	return func(c *gin.Context) {
		posts, err := db.getPosts()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("could not get posts: %v", err))
			return
		}

		c.HTML(http.StatusOK, "index", gin.H{"posts": posts})
	}
}

func makeDatabase(settings AppSettings) Database {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/cms-and-go",
		settings.databaseUser, settings.databasePassword, settings.databaseAddress, settings.databasePort)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return Database{
		address:    settings.databaseAddress,
		port:       settings.databasePort,
		user:       settings.databaseUser,
		connection: db,
	}
}

func main() {
	

	appSettings := loadSettings()
	dbConnection := makeDatabase(appSettings)

	r := gin.Default()
	r.MaxMultipartMemory = 1
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", makeHomeHandler(appSettings, dbConnection))

	r.POST("/contact-send", func(c *gin.Context) {
		c.Request.ParseForm()
		name := c.Request.FormValue("name")
		email := c.Request.FormValue("email")
		message := c.Request.FormValue("message")

		_, err := mail.ParseAddress(email)
		if err != nil {
			c.HTML(http.StatusOK, "contact-failure.html", gin.H{
				"email": email,
				"error": "Invalid email",
			})
			return
		}

		if len(name) > 50 {
			c.HTML(http.StatusOK, "contact-failure.html", gin.H{
				"email": email,
				"error": "Invalid name",
			})
			return
		}

		if len(message) > 500 {
			c.HTML(http.StatusOK, "contact-failure.html", gin.H{
				"email": email,
				"error": "Message is too long",
			})
			return
		}

		c.HTML(http.StatusOK, "contact-success.html", gin.H{
			"name":  name,
			"email": email,
		})
	})

	r.Static("/static", "./static")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
