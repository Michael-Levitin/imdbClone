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
	if len(entry.Actor) < 5 && len(entry.Movie) < 5 {
		return &[]dto.List{}, fmt.Errorf("query too short")
	}
	return c.CloneDB.FindPartsDB(ctx, entry)
}

func (c CloneLogic) FindActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: FindActors recieved %+v\n", entry))
	if len(entry.Actor) < 5 && len(entry.Movie) < 5 {
		return &[]dto.Actor{}, fmt.Errorf("query too short")
	}
	return c.CloneDB.FindActorsDB(ctx, entry)
}

func (c CloneLogic) RemoveMovies(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: RemoveMovies recieved %+v\n", entry))
	if len(entry.Actor) < 5 && len(entry.Movie) < 5 {
		return &[]dto.Movie{}, fmt.Errorf("query too short")
	}
	return c.CloneDB.RemoveMoviesDB(ctx, entry)
}

func (c CloneLogic) RemoveActors(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	log.Debug().Msg(fmt.Sprintf("Logic: RemoveActors recieved %+v\n", entry))
	if len(entry.Actor) < 5 && len(entry.Movie) < 5 {
		return &[]dto.Actor{}, fmt.Errorf("query too short")
	}
	return c.CloneDB.RemoveActorsDB(ctx, entry)
}
