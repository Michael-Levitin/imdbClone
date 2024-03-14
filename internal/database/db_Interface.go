package database

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type CloneDbI interface {
	FindPartsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error)
	RemoveMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	RemoveActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
}
