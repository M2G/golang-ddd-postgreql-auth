package handlers

/*
import (
	"encoding/json"
	"fmt"
	"github.com/golang-ddd-postgreql-auth/internal/repositories"
	"github.com/golang-ddd-postgreql-auth/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/golang-ddd-postgreql-auth/internal/drivers"
	models "github.com/golang-ddd-postgreql-auth/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NewPostHandler
func NewPostHandler(db *drivers.DB) *Post {
	return &Post{
		repo: repositories.NewSQLPostRepo(db.SQL),
	}
}

// Post
type Post struct {
	repo store.Store
}

// Fetch all post data
func (p *Post) Fetch(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, _ := p.repo.Fetch(r.Context(), 5)

		log.Info("Successfully fetch")
		respondwithJSON(w, http.StatusOK, payload)
	}
}

// Create a new post
func (p *Post) Create(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//nolint
		post := models.Post{}

		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			log.Error("an error occurred while decoding response body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := p.repo.Create(r.Context(), &post)

		if id == -1 {
			log.Error("Server Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Error("Server Error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debugf("Successfully Created %v", id)
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}
}

// Update a post by id
func (p *Post) Update(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//nolint
		postID, error := getPostIDFromRequest(r)

		if error != nil {
			log.Error("couldnt get postID")
			http.Error(w, error.Error(), http.StatusUnprocessableEntity)
			return
		}

		data := models.Post{ID: *postID}

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			log.Error("an error occurred while decoding response body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload, err := p.repo.Update(r.Context(), &data)

		if err != nil {
			log.Error("an error occurred while update")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debug("Successfully Updated")
		respondwithJSON(w, http.StatusOK, payload)
	}
}

// GetByID returns a post details
func (p *Post) GetByID(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		postID, err := getPostIDFromRequest(r)

		if err != nil {
			log.Error("couldnt get postID")
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		payload, err := p.repo.GetByID(r.Context(), *postID)

		if err != nil {
			log.Error("Content not found")
			http.Error(w, err.Error(), http.StatusNoContent)
			return
		}

		respondwithJSON(w, http.StatusOK, payload)
	}
}

// Delete a post
func (p *Post) Delete(log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		postID, err := getPostIDFromRequest(r)
		if err != nil {
			log.Error("couldnt get postID")
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		payload, err := p.repo.Delete(r.Context(), *postID)

		fmt.Println(payload)

		if err != nil {
			log.Error("an error occurred while delete")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debug("Delete Successfully")
		respondwithJSON(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
	}
}

func getPostIDFromRequest(request *http.Request) (*int64, error) {
	if request == nil {
		return nil, errors.New("empty request")
	}

	id := chi.URLParam(request, "id")
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.WithStack(
			errors.Wrapf(err, "cannot parse postID as %s", id),
		)
	}

	return &postID, nil
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
*/
