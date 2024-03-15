package logic

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type CloneLogicI interface {
	FindParts(ctx context.Context, entry *dto.Entry) (*[]dto.List, error)
	FindActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
	FindMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	AddActors(ctx context.Context, actors *[]dto.Actor) (*[]dto.Id, error)
	AddMovie(ctx context.Context, movie *dto.Movie) (*dto.Id, error)
	RemoveMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	RemoveActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
}
