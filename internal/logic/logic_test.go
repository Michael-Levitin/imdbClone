package logic

import (
	"context"
	"reflect"
	"testing"

	"github.com/Michael-Levitin/imdbClone/internal/database"
	"github.com/Michael-Levitin/imdbClone/internal/dto"
)

var c = CloneLogic{CloneDB: database.NewCloneDBMock()}

func TestCloneLogic_AddActors(t *testing.T) {
	type args struct {
		ctx    context.Context
		actors *[]dto.Actor
	}
	var tests = []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "nil data",
			args: args{
				ctx:    context.Background(),
				actors: nil,
			},
			want:    []int{},
			wantErr: dto.ErrNoData,
		},
		{
			name: "empty",
			args: args{
				ctx:    context.Background(),
				actors: &[]dto.Actor{},
			},
			want:    []int{},
			wantErr: nil,
		},
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				actors: &[]dto.Actor{{
					Name: "Brad",
				}},
			},
			want:    []int{1},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.AddActors(tt.args.ctx, tt.args.actors)
			if err != tt.wantErr {
				t.Errorf("AddActors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddActors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_AddMParts(t *testing.T) {
	type args struct {
		ctx  context.Context
		imdb *dto.Imdb
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "empty",
			args: args{
				ctx:  context.Background(),
				imdb: &dto.Imdb{},
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				imdb: &dto.Imdb{
					Movie:  dto.Movie{Movie: "Brave"},
					Actors: []dto.Actor{{Name: "Brad"}},
				},
			},
			want:    2,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.AddMParts(tt.args.ctx, tt.args.imdb)
			if err != tt.wantErr {
				t.Errorf("AddMParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddMParts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_AddMovie(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				movie: &dto.Movie{},
			},
			want:    0,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				movie: &dto.Movie{
					Movie: "Brave",
				},
			},
			want:    1,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.AddMovie(tt.args.ctx, tt.args.movie)
			if err != tt.wantErr {
				t.Errorf("AddMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_FindActors(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Actor: "Pit"},
			},
			want:    &[]dto.Actor{{Name: "Pitt"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.FindActors(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindActors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindActors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_FindMovies(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Movie: "Bra"},
			},
			want:    &[]dto.Movie{{Movie: "Brave"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.FindMovies(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_FindParts(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Movie: "Bra"},
			},
			want:    &[]dto.List{{Movie: "Brave"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.FindParts(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("FindParts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindParts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_RemoveActors(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Actor: "Pit"},
			},
			want:    &[]dto.Actor{{Name: "Pitt"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RemoveActors(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("RemoveActors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveActors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloneLogic_RemoveMovies(t *testing.T) {
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
			name: "empty",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{},
			},
			want:    nil,
			wantErr: dto.ErrNoData,
		},
		{
			name: "OK",
			args: args{
				ctx:   context.Background(),
				entry: &dto.Entry{Actor: "Bra"},
			},
			want:    &[]dto.Movie{{Movie: "Brave"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.RemoveMovies(tt.args.ctx, tt.args.entry)
			if err != tt.wantErr {
				t.Errorf("RemoveMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
