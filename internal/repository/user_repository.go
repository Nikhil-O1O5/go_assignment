package repository

import (
	"context"
	"database/sql"
	"go-backend-task/db/sqlc" 
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		queries: db.New(conn),
	}
}

func (r *UserRepository) Create(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *UserRepository) GetByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUser(ctx, id)
}

func (r *UserRepository) Update(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) ListAll(ctx context.Context) ([]db.User, error) {
	return r.queries.ListUsers(ctx)
}