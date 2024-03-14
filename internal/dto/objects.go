package dto

import "time"

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
	Movie string `json:"movie" db:"movie"`
}

type Actor struct {
	Name string `json:"name" db:"name"`
}
