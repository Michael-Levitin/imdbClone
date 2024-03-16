package database

import (
	"context"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

type cloneDBMock struct {
}

func NewCloneDBMock() *cloneDBMock {
	return &cloneDBMock{}
}

func (c cloneDBMock) FindPartsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	if entry == nil || entry.Movie == "" {
		return nil, dto.ErrNoData
	}
	if entry.Movie == "Bra" {
		return &[]dto.List{{Movie: "Brave"}}, nil
	}
	return nil, nil
}

func (c cloneDBMock) FindActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	if entry == nil || entry.Actor == "" {
		return nil, dto.ErrNoData
	}
	if entry.Actor == "Pit" {
		return &[]dto.Actor{{Name: "Pitt"}}, nil
	}
	return nil, nil
}

func (c cloneDBMock) FindMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	if entry == nil || entry.Movie == "" {
		return nil, dto.ErrNoData
	}
	if entry.Movie == "Bra" {
		return &[]dto.Movie{{Movie: "Brave"}}, nil
	}
	return nil, nil
}

func (c cloneDBMock) AddActorsDB(ctx context.Context, actors *[]dto.Actor) ([]int, error) {
	if actors == nil {
		return []int{}, dto.ErrNoData
	}
	if len(*actors) == 0 {
		return []int{}, nil
	}
	if len(*actors) == 1 && (*actors)[0].Name == "Brad" {
		return []int{1}, nil
	}

	return []int{}, nil
}

func (c cloneDBMock) AddMovieDB(ctx context.Context, movie *dto.Movie) (int, error) {
	if movie == nil || movie.Movie == "" {
		return 0, dto.ErrNoData
	}
	if movie.Movie == "Brave" {
		return 1, nil
	}
	return 0, nil
}

func (c cloneDBMock) AddPartsDB(ctx context.Context, movie int, actors []int) (int, error) {
	if movie == 0 || actors == nil {
		return 0, dto.ErrNoData
	}
	if len(actors) == 0 {
		return 0, nil
	}
	if len(actors) == 1 && (actors)[0] == 1 {
		return 2, nil
	}

	return 0, nil
}

func (c cloneDBMock) RemoveMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	if entry == nil || entry.Actor == "" {
		return nil, dto.ErrNoData
	}
	if entry.Actor == "Bra" {
		return &[]dto.Movie{{Movie: "Brave"}}, nil
	}
	return nil, nil
}

func (c cloneDBMock) RemoveActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	if entry == nil || entry.Actor == "" {
		return nil, dto.ErrNoData
	}
	if entry.Actor == "Pit" {
		return &[]dto.Actor{{Name: "Pitt"}}, nil
	}
	return nil, nil
}
