package user

import (
	"context"
	"database/sql"
	"server/helper"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "insert into users(username, email, password) values (?, ?, ?)"

	res, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	helper.PanicIfError(err, "error saat query context")

	id, err := res.LastInsertId()
	helper.PanicIfError(err, "error get last insert id")

	user.Id = int64(id)

	return user, nil
}
