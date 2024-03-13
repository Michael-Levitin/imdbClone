package database

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CloneDB struct {
	db *pgxpool.Pool
}

func NewImdbCloneDB(db *pgxpool.Pool) *CloneDB {
	return &CloneDB{db: db}
}

func (h CloneDB) FindMoviesActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	//TODO implement me
	panic("implement me")
}

func (h CloneDB) RemoveMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	//TODO implement me
	panic("implement me")
}

func (h CloneDB) RemoveActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	//TODO implement me
	panic("implement me")
}
