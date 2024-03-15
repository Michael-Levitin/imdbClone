package database

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type CloneDbI interface {
	FindPartsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error)
	FindActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
	FindMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	AddActorsDB(ctx context.Context, actors *[]dto.Actor) ([]int, error)
	AddMovieDB(ctx context.Context, movie *dto.Movie) (int, error)
	AddPartsDB(ctx context.Context, movie int, actors []int) (int, error)
	RemoveMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	RemoveActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
}
