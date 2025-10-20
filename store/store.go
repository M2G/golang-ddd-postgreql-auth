package store

import (
	"context"

	"github.com/golang-ddd-postgresql-auth/models"
)

// Store explain...
type Store interface {
	Fetch(ctx context.Context, args ...interface{}) ([]*models.User, error)
	//GetByID(ctx context.Context, id int64) (*models.User, error)
	//Create(ctx context.Context, p *models.User) (int64, error)
	//Update(ctx context.Context, p *models.User) (*models.User, error)
	//Delete(ctx context.Context, id int64) (bool, error)
}
