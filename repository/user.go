package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/sanoyo/ultrachat/graph/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (s *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query, args, err := sq.Select("users").Columns("user_id, name, email").Where(sq.Eq{"email": email}).ToSql()
	if err != nil {
		return nil, err
	}

	var user model.User
	err = s.db.QueryRow(query, args...).Scan(&user)
	if err != nil {
		return nil, err
	}

	return &user, err
}
