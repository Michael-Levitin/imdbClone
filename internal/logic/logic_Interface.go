package logic

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type CloneLogicI interface {
	FindMoviesActors(ctx context.Context, item *dto.Entry) ([]*dto.List, error)
	RemoveMovies(ctx context.Context, item *dto.Entry) ([]*dto.List, error)
	RemoveActors(ctx context.Context, info *dto.Entry) ([]*dto.List, error)
}
