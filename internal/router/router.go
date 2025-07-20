package router

import (
  "net/http"

  "github.com/go-chi/chi"
	"github.com/go-mysql-crud/internal/handlers"
	"github.com/sirupsen/logrus"
)

// A completely separate router for posts routes
func PostRouter(pHandler *handlers.Post, log *logrus.Logger) http.Handler { // nolint
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch(log))
	r.Get("/{id:[0-9]+}", pHandler.GetByID(log))
	r.Post("/", pHandler.Create(log))
	r.Put("/{id:[0-9]+}", pHandler.Update(log))
	r.Delete("/{id:[0-9]+}", pHandler.Delete(log))

	return r
}
