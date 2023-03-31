package repository

import (
	"context"
	"database/sql"

	"github.com/sanoyo/ultrachat/models"
)

type SpaceRepository struct {
	db *sql.DB
}

func NewSpaceRepository(db *sql.DB) *SpaceRepository {
	return &SpaceRepository{db}
}

func (r *SpaceRepository) Create(ctx context.Context, space *models.Space) error {
	return space.Insert(ctx, r.db)
}

func (r *SpaceRepository) Update(ctx context.Context, space *models.Space) error {
	return space.Update(ctx, r.db)
}

func (r *SpaceRepository) Delete(ctx context.Context, space *models.Space) error {
	return space.Delete(ctx, r.db)
}
