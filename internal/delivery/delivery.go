package delivery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Michael-Levitin/imdbClone/internal/dto"
	"github.com/Michael-Levitin/imdbClone/internal/logic"
	"github.com/rs/zerolog/log"
)

type CloneServer struct {
	logic logic.CloneLogicI
}

func NewImdbCloneServer(logic logic.CloneLogicI) *CloneServer {
	return &CloneServer{logic: logic}
}

func (c CloneServer) FindMoviesActors(w http.ResponseWriter, r *http.Request) {
	entry, err := getParam(r)
	if err != nil {
		log.Warn().Err(err).Msg("error reading parameters")
		fmt.Fprintln(w, err)
		return
	}

	list, err := c.logic.FindMoviesActors(context.Background(), entry)
	if err != nil {
		log.Warn().Err(err).Msg("error executing c.logic.FindMoviesActors")
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (c CloneServer) RemoveMovies(w http.ResponseWriter, r *http.Request) {
	entry, err := getParam(r)
	if err != nil {
		fmt.Fprintln(w, "error reading parameters: ", err)
		return
	}

	list, err := c.logic.GoodUpdate(context.Background(), entry)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (c CloneServer) RemoveActors(w http.ResponseWriter, r *http.Request) {
	entry, err := getParam(r)
	if err != nil {
		fmt.Fprintln(w, "error reading parameters: ", err)
		return
	}

	list, err := c.logic.GoodRemove(context.Background(), entry)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func getParam(r *http.Request) (*dto.Entry, error) {
	var err error
	if err = r.ParseForm(); err != nil {
		return nil, fmt.Errorf("ParseForm() err: %v", err)
	}

	queryParams := r.URL.Query()
	entry := dto.Entry{
		Actor: queryParams.Get("movie"),
		Movie: queryParams.Get("actor"),
	}

	return &entry, nil
}
