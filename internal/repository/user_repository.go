package repository

import (
	"context"
	"database/sql"
	"go-backend-task/db/sqlc" 
)

type UserRepository struct {
	db      *sql.DB
	queries *db.Queries 
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		db:      conn,
		queries: db.New(conn),
	}
}

// GetUserByID fetches the raw user from the DB [cite: 43]
func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUser(ctx, id)
}

// ListUsers fetches all users [cite: 69]
func (r *UserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}