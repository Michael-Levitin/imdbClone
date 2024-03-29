package dto

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ErrNoData = fmt.Errorf("not data")

type Entry struct {
	Actor string `json:"actor,omitempty"`
	Movie string `json:"movie,omitempty"`
}

type List struct {
	Movie   string    `json:"movie" db:"movie"`
	Rating  float32   `json:"rating" db:"rating"`
	Release time.Time `json:"release" db:"release"`
	Name    string    `json:"name" db:"name"`
	DOB     time.Time `json:"dob" db:"dob"`
}

type Movie struct {
	Id          int       `json:"id" db:"id"`
	Movie       string    `json:"movie" db:"movie"`
	Description string    `json:"description" db:"description"`
	Release     time.Time `json:"release" db:"release"`
	Rating      float32   `json:"rating" db:"rating"`
	Removed     bool      `json:"removed" db:"removed"`
	Created     time.Time `json:"created_at" db:"created_at"`
}

type Actor struct {
	Id      int       `json:"id" db:"id"`
	Name    string    `json:"name" db:"name"`
	DOB     time.Time `json:"dob" db:"dob"`
	Removed bool      `json:"removed" db:"removed"`
	Created time.Time `json:"created_at" db:"created_at"`
}

func ActorsToString(actors *[]Actor) string {
	var jsonStr strings.Builder

	for i, actor := range *actors {
		jsonStr.WriteString("('" + actor.Name + "', ")
		jsonStr.WriteString("'" + actor.DOB.Format("2006-01-02 15:04:05") + "', ")
		jsonStr.WriteString(strconv.FormatBool(actor.Removed) + ", ")
		jsonStr.WriteString("'" + actor.Created.Format("2006-01-02 15:04:05") + "')")

		if i < len(*actors)-1 {
			jsonStr.WriteString(",\n")
		}
	}

	defer jsonStr.Reset()
	return jsonStr.String()
}

type Id struct {
	Id int `json:"id" db:"id"`
}

type Imdb struct {
	Movie  Movie
	Actors []Actor
}
