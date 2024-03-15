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
	_findActorQuery = `
SELECT a.name
FROM actors a
WHERE a.name ILIKE @actor
    AND removed = false
ORDER BY a.name;
`
	_findMovieQuery = `
SELECT m.movie
FROM movies m
WHERE m.movie ILIKE @movie
  AND removed = false
ORDER BY m.movie;
`
	_addActorsQueryHead = `
INSERT INTO actors (name, dob, removed, created_at)
VALUES
`
	_addActorsQueryTail = `
ON CONFLICT(name)
    DO UPDATE SET dob        = excluded.dob,
                  removed    = excluded.removed,
                  created_at = excluded.created_at
RETURNING id;
`
	_addMovieQuery = `
INSERT INTO movies (movie, description, release, rating, removed, created_at)
VALUES (@movie, @description, @release, @rating, @removed, @created_at)
ON CONFLICT(movie)
    DO UPDATE SET description = excluded.description,
                  release     = excluded.release,
                  rating      = excluded.rating,
                  removed     = excluded.removed,
                  created_at  = excluded.created_at
RETURNING id;
`
	_addPartsQuery = `
INSERT INTO parts (movie_id, actor_id)
SELECT m.movie_id, actor_id
FROM (select @movie::int as movie_id) m
         CROSS JOIN unnest(@actors::int[]) actor_id
ON CONFLICT(movie_id, actor_id)
    DO NOTHING;`
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

func (c CloneDB) FindActorsDB(ctx context.Context, entry *dto.Entry) (*[]dto.Actor, error) {
	rows, err := c.db.Query(ctx, _findActorQuery,
		pgx.NamedArgs{"actor": entry.Actor})
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("FindActorsDB could not get list %+v", entry))
		return &[]dto.Actor{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Actor])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.Actor{}, err
	}

	return &list, nil
}

func (c CloneDB) FindMoviesDB(ctx context.Context, entry *dto.Entry) (*[]dto.Movie, error) {
	rows, err := c.db.Query(ctx, _findMovieQuery,
		pgx.NamedArgs{"movie": entry.Movie})
	if err != nil {
		log.Debug().Err(err).Msg(fmt.Sprintf("FindMoviesDB could not get list %+v", entry))
		return &[]dto.Movie{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Movie])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return &[]dto.Movie{}, err
	}

	return &list, nil
}

func (c CloneDB) AddActorsDB(ctx context.Context, actors *[]dto.Actor) ([]int, error) {
	rows, err := c.db.Query(ctx, _addActorsQueryHead+dto.ActorsToString(actors)+_addActorsQueryTail)
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("AddActorsDB could not add actors"))
		return []int{}, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Id])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return []int{}, err
	}

	ids := make([]int, len(list))
	for i, id := range list {
		ids[i] = id.Id
	}

	return ids, nil
}

func (c CloneDB) AddMovieDB(ctx context.Context, movie *dto.Movie) (int, error) {
	rows, err := c.db.Query(ctx, _addMovieQuery,
		pgx.NamedArgs{"movie": movie.Movie,
			"description": movie.Description,
			"release":     movie.Release,
			"rating":      movie.Rating,
			"removed":     movie.Removed,
			"created_at":  movie.Created,
		})

	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("AddMovieDB could not add movie"))
		return 0, err
	}

	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.Id])
	if err != nil {
		log.Trace().Err(err).Msg(fmt.Sprintf("CollectRows error"))
		return 0, err
	}

	return list[0].Id, nil
}

func (c CloneDB) AddPartsDB(ctx context.Context, movie int, actors []int) (int, error) {
	cmd, err := c.db.Exec(ctx, _addPartsQuery, pgx.NamedArgs{"movie": movie, "actors": actors})
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("AddPartsDB could not add parts"))
		return 0, err
	}
	return int(cmd.RowsAffected()), nil
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
