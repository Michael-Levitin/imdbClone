package database

import (
	"context"
	"fmt"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

const (
	_findPartsQuery = `
SELECT m.movie as movie, m.rating, m.release, a.name, a.dob
FROM movies m
         JOIN parts p on m.id = p.movie_id
         JOIN actors a on a.id = p.actor_id
WHERE m.movie ILIKE @movie
  AND a.name ILIKE @actor
  AND m.removed = false
  AND a.removed = false
ORDER BY m.rating DESC, a.name;
`
	_removeMoviesQuery = `
UPDATE movies
SET removed = true
WHERE movie ILIKE @movie
	AND removed = false
RETURNING movie;
`
	_removeActorsQuery = `
UPDATE actors
SET removed = true
WHERE name ILIKE @actor
	AND removed = false
RETURNING name;
`
)

type CloneDB struct {
	db *pgxpool.Pool
}

func NewImdbCloneDB(db *pgxpool.Pool) *CloneDB {
	return &CloneDB{db: db}
}

func (c CloneDB) FindPartsDB(ctx context.Context, entry *dto.Entry) (*[]dto.List, error) {
	rows, err := c.db.Query(ctx, _findPartsQuery,
		pgx.NamedArgs{"movie": entry.Movie, "actor": entry.Actor})
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("FindPartsDB could not get list %+v", entry))
		return &[]dto.List{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.List])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.List{}, err
	}

	return &list, nil
}

func (c CloneDB) RemoveMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	rows, err := c.db.Query(ctx, _removeMoviesQuery,
		pgx.NamedArgs{"movie": entry.Movie})
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("RemoveMoviesDB could not get list %+v", entry))
		return &[]dto.Movie{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Movie])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.Movie{}, err
	}

	return &list, nil
}

func (c CloneDB) RemoveActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	rows, err := c.db.Query(ctx, _removeActorsQuery,
		pgx.NamedArgs{"actor": entry.Actor})
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("RemoveActorsDB could not get list %+v", entry))
		return &[]dto.Actor{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Actor])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.Actor{}, err
	}

	return &list, nil
}
