package logic

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type CloneLogicI interface {
	FindParts(ctx context.Context, entry *dto.Entry) (*[]dto.List, error)
	RemoveMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error)
	RemoveActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error)
}
