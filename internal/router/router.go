package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-ddd-auth/internal/handlers"
	"github.com/sirupsen/logrus"
)

func Routes(router *gin.Engine, pHandler *handlers.User, log *logrus.Logger) http.Handler {
	log.Println("ok")

	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ok", pHandler.Fetch(log))
	//r.Get("/{id:[0-9]+}", pHandler.GetByID(log))
	//r.Post("/", pHandler.Create(log))
	//r.Put("/{id:[0-9]+}", pHandler.Update(log))
	//r.Delete("/{id:[0-9]+}", pHandler.Delete(log))

	return router
}
