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

func (s *SpaceRepository) CreateSpace(ctx context.Context, name string) (int64, error) {
	query, args, err := sq.Insert("spaces").Columns("name").Values(name).ToSql()
	if err != nil {
		return 0, nil
	}

	result, err := s.db.Exec(query, args...)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return id, err
}

// err = s.db.QueryRow(sql, args...).Scan(&model.)
// if err != nil {
// 	return err
// }
