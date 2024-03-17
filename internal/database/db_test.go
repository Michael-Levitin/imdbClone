package database

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	db, err = pgxpool.New(context.Background(), "postgresql://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to database")
	} else {
		log.Info().Msg("connected to database for testing")
	}
	code := m.Run()
	defer db.Close()
	os.Exit(code)
}

func TestCloneDB_RemoveActorsDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *dto.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    *[]dto.Actor
		wantErr error
	}{
		{
			name: "Nil",
			args: args{
				ctx:   context.Background(),
				entry: nil,
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx: context.Background(),
				entry: &dto.Entry{
					Actor: "Sharon Stone",
				},
			},
			want:    &[]dto.Actor{{Name: "Sharon Stone"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.RemoveActorsDB(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("RemoveActorsDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			for i, actor := range *got {
				if actor.Name != (*tt.want)[i].Name {
					t.Errorf("RemoveActorsDB() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestCloneDB_RemoveMoviesDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *dto.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    *[]dto.Movie
		wantErr error
	}{
		{
			name: "Nil",
			args: args{
				ctx:   context.Background(),
				entry: nil,
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx: context.Background(),
				entry: &dto.Entry{
					Movie: "BraveHeart",
				},
			},
			want:    &[]dto.Movie{{Movie: "BraveHeart"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.RemoveMoviesDB(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("RemoveMoviesDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			for i, movie := range *got {
				if movie.Movie != (*tt.want)[i].Movie {
					t.Errorf("RemoveMoviesDB() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestCloneDB_AddActorsDB(t *testing.T) {
	type args struct {
		ctx    context.Context
		actors *[]dto.Actor
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "Nil",
			args: args{
				ctx:    context.Background(),
				actors: nil,
			},
			want:    []int{},
			wantErr: dto.ErrNoData,
		},
		{
			name: "Empty",
			args: args{
				ctx:    context.Background(),
				actors: &[]dto.Actor{},
			},
			want:    []int{},
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx: context.Background(),
				actors: &[]dto.Actor{{
					Name:    "Sharon Stone",
					DOB:     time.Now(),
					Removed: false,
					Created: time.Now(),
				}},
			},
			want:    []int{8},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.AddActorsDB(tt.args.ctx, tt.args.actors)
			if err != tt.wantErr {
				t.Errorf("AddActorsDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddActorsDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneDB_AddMovieDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		movie *dto.Movie
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Nil",
			args: args{
				ctx:   context.Background(),
				movie: nil,
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Empty",
			args: args{
				ctx:   context.Background(),
				movie: &dto.Movie{},
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx: context.Background(),
				movie: &dto.Movie{
					Movie:       "BraveHeart",
					Description: "",
					Release:     time.Now(),
					Rating:      8.5,
					Removed:     false,
					Created:     time.Now(),
				},
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.AddMovieDB(tt.args.ctx, tt.args.movie)
			if err != tt.wantErr {
				t.Errorf("AddMovieDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddMovieDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneDB_AddPartsDB(t *testing.T) {
	type args struct {
		ctx    context.Context
		movie  int
		actors []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "No movie",
			args: args{
				ctx:   context.Background(),
				movie: 0,
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "No actors",
			args: args{
				ctx:    context.Background(),
				movie:  1,
				actors: []int{},
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx:    context.Background(),
				movie:  1,
				actors: []int{1, 2, 3, 4},
			},
			want:    0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.AddPartsDB(tt.args.ctx, tt.args.movie, tt.args.actors)
			if err != tt.wantErr {
				t.Errorf("AddPartsDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddPartsDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneDB_FindActorsDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *dto.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    *[]dto.Actor
		wantErr error
	}{
		{
			name: "nil",
			args: args{
				ctx:   context.Background(),
				entry: nil,
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "No actors",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Actor: "%Tom Cruise%"},
			},
			want: &[]dto.Actor{
				{Name: "Tom Cruise"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.FindActorsDB(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindActorsDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want == nil {
				return
			}

			for i, actor := range *got {
				if actor.Name != (*tt.want)[i].Name {
					t.Errorf("FindActorsDB() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestCloneDB_FindMoviesDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *dto.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    *[]dto.Movie
		wantErr error
	}{
		{
			name: "nil",
			args: args{
				ctx:   context.Background(),
				entry: nil,
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "No movie",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "Add one",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Movie: "%BraveHeart%"},
			},
			want: &[]dto.Movie{
				{Movie: "BraveHeart"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.FindMoviesDB(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindMoviesDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want == nil {
				return
			}

			for i, movie := range *got {
				if movie.Movie != (*tt.want)[i].Movie {
					t.Errorf("FindActorsDB() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestCloneDB_FindPartsDB(t *testing.T) {
	type args struct {
		ctx   context.Context
		entry *dto.Entry
	}
	tests := []struct {
		name    string
		args    args
		want    *[]dto.List
		wantErr error
	}{

		{
			name: "nil",
			args: args{
				ctx:   context.Background(),
				entry: nil,
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "No data",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "find",
			args: args{
				ctx: context.Background(),
				entry: &dto.Entry{
					Actor: "%Tom%",
					Movie: "%Tom%",
				},
			},
			want: &[]dto.List{
				{Movie: "Edge of Tomorrow"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CloneDB{
				db: db,
			}
			got, err := c.FindPartsDB(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindPartsDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			for i, movie := range *got {
				if movie.Movie != (*tt.want)[i].Movie {
					t.Errorf("FindPartsDB() got = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}

func TestNewImdbCloneDB(t *testing.T) {

	tests := []struct {
		name string
		want *CloneDB
	}{
		{
			name: "create",
			want: &CloneDB{db: db},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewImdbCloneDB(db)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImdbCloneDB() = %v, want %v", got, tt.want)
			}
			err := got.db.Ping(context.Background())
			if err != nil {
				t.Errorf("NewImdbCloneDB() = expected no error")
			}
		})
	}
}
