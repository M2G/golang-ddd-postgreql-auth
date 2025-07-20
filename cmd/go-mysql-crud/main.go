package main

import (
	"fmt"
	"github.com/go-mysql-crud/internal/handlers"
	"log"
	"net/http"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-mysql-crud/internal/drivers"
  "github.com/go-mysql-crud/internal/router"
	logrus "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func RunAPI(c *cli.Context) error {

	log := logrus.New() // nolint
	log.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})

	dbName := "test_db"
	dbPass := "docker"
	dbHost := "db"
	dbPort := "3306"

	connection, err := drivers.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := handlers.NewPostHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", router.PostRouter(pHandler, log))
	})

	log.Info("Server listen at :8181")
	serverErr := http.ListenAndServe(":8181", r)

	if serverErr != nil {
		log.Println("Error starting server")
		return nil
	}

	fmt.Println("Started server on - 127.0.0.1:8080")

	return nil // nolint
}

func main() {
	app := &cli.App{
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Usage:    "API CRUD using MySQL",
		Commands: []*cli.Command{
			{
				Name:    "api",
				Aliases: []string{"a"},
				Usage:   "a cli application for the api crud.",
				Action:  RunAPI,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("An error occurred: %s", err)
		return
	}
}

/*

// Example for test mysq

package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/davecgh/go-spew/spew"
)

func main() {
    d, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
    if err != nil {
        log.Fatal(err)
    }

    err = d.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("connected")

    rows, err := d.Query("SELECT id, title, content FROM test_tb")
    if err != nil {
        log.Fatal(err)
    }
    posts := []*Post{}
    for rows.Next() {
        post := Post{}
        err := rows.Scan(&post.ID, &post.Title, &post.Content)
        if err != nil {
            continue
        }
        posts = append(posts, &post)
    }

    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    spew.Dump(posts)

    err = d.Close()
    if err != nil {
        log.Fatal(err)
    }
}

type Post struct {
    ID      int64  `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}*/
