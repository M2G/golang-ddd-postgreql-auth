package repositories

/*
import (
	"context"
	"database/sql"
	"github.com/go-mysql-crud/store"

	models "github.com/go-mysql-crud/models"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(conn *sql.DB) store.Store {
	return &mysqlPostRepo{
		conn: conn,
	}
}

type mysqlPostRepo struct {
	conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		if err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		); err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
	query := "SELECT id, title, content FROM test_tb LIMIT ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := "SELECT id, title, content FROM test_tb WHERE id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	//nolint
	payload := &models.Post{}

	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "INSERT test_tb SET title=?, content=?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlPostRepo) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	query := "UPDATE test_tb SET title=?, content=? WHERE id=?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM test_tb WHERE id=?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
*/
