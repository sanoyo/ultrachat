package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type SpaceRepository struct {
	db *sql.DB
}

func NewSpaceRepository(db *sql.DB) *SpaceRepository {
	return &SpaceRepository{db}
}

func (s *SpaceRepository) CreateSpace(ctx context.Context, name string) error {
	query, args, err := sq.Insert("spaces").Columns("name").Values(name).ToSql()
	if err != nil {
		return nil
	}

	_, err = s.db.Exec(query, args...)
	return err
}

// err = s.db.QueryRow(sql, args...).Scan(&model.)
// if err != nil {
// 	return err
// }
