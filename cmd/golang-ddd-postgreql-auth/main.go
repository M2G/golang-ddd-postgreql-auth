package main

import (
  "fmt"
  "golang-ddd-postgresql-auth/internal/drivers"

  _ "github.com/gin-gonic/gin"
  // "github.com/golang-ddd-postgresql-auth/internal/handlers"
  "context"
  "log"
  _ "net/http"
  "os"
  "time"

  nested "github.com/antonfisher/nested-logrus-formatter"
  //"github.com/go-chi/chi/middleware"
  //"golang-ddd-postgresql-auth/internal/router"
  logrus "github.com/sirupsen/logrus"
  "github.com/urfave/cli/v3"
)

func RunAPI(ctx context.Context, cmd *cli.Command) error {
  log := logrus.New()
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

  log.Info(connection)

  /*
  	r := chi.NewRouter()
  	r.Use(middleware.RequestID)
  	r.Use(middleware.RealIP)
  	r.Use(middleware.StripSlashes)
  	r.Use(middleware.Recoverer)
  	r.Use(middleware.Logger)
  */
  /*
  	pHandler := handlers.NewPostHandler(connection)

  	r.Route("/", func(rt chi.Router) {
  		rt.Mount("/posts", router.PostRouter(pHandler, log))
  	})

  	log.Info("Server listen at :8181")
  	serverErr := http.ListenAndServe(":8181", r)

  	if serverErr != nil {
  		log.Println("Error starting server")
  		return nil
  	}*/

  fmt.Println("Started server on - 127.0.0.1:8080")

  return nil // nolint
}

func main() {
  app := &cli.Command{
    Version: "v1.0.0",
    Usage:   "API CRUD using MySQL",
    Action: func(ctx context.Context, cmd *cli.Command) error {
      fmt.Printf(time.Now().String())
      return nil
    },
    Commands: []*cli.Command{
      {
        Name:    "api",
        Aliases: []string{"a"},
        Usage:   "a cli application for the api crud.",
        Action:  RunAPI,
      },
    },
  }

  if err := app.Run(context.Background(), os.Args); err != nil {
    log.Fatalf("An error occurred: %s", err)
    return
  }
}

/*
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("connected")

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
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
