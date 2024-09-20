package user

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "select id, username, email from users where email=?"

	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Username, &u.Email)
	if err != nil {
		fmt.Println("error nich : ", err)
		return &User{}, nil
	}

	return &u, nil
}
