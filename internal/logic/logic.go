package logic

import (
	"context"
	"fmt"

	"github.com/Michael-Levitin/imdbClone/internal/database"
	"github.com/Michael-Levitin/imdbClone/internal/dto"
	"github.com/rs/zerolog/log"
)

type CloneLogic struct {
	CloneDB database.CloneDbI
}

// NewImdbCloneLogic подключаем интерфейс БД в новую логику
func NewImdbCloneLogic(CloneDB *database.CloneDB) *CloneLogic {
	return &CloneLogic{CloneDB: CloneDB}
}

func (c CloneLogic) FindParts(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: FindParts recieved %+v\n", entry))
	return c.CloneDB.FindPartsDB(ctx, entry)
}

func (c CloneLogic) FindActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: FindActors recieved %+v\n", entry))
	return c.CloneDB.FindActorsDB(ctx, entry)
}

func (c CloneLogic) FindMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: FindMovies recieved %+v\n", entry))
	return c.CloneDB.FindMoviesDB(ctx, entry)
}

func (c CloneLogic) AddActors(ctx context.Context, actors *[]dto.Actor) (*[]dto.Id, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: AddActors recieved %+v\n", actors))
	return c.CloneDB.AddActorsDB(ctx, actors)
}

func (c CloneLogic) AddMovie(ctx context.Context, movie *dto.Movie) (*dto.Id, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: AddMovie recieved %+v\n", movie))
	return c.CloneDB.AddMovieDB(ctx, movie)
}

func (c CloneLogic) RemoveMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: RemoveMovies recieved %+v\n", entry))
	return c.CloneDB.RemoveMoviesDB(ctx, entry)
}

func (c CloneLogic) RemoveActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: RemoveActors recieved %+v\n", entry))
	return c.CloneDB.RemoveActorsDB(ctx, entry)
}
