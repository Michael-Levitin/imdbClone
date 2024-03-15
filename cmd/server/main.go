package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Michael-Levitin/imdbClone/config"
	"github.com/Michael-Levitin/imdbClone/internal/database"
	"github.com/Michael-Levitin/imdbClone/internal/delivery"
	"github.com/Michael-Levitin/imdbClone/internal/logic"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// загружаем конфиг
	config.Init()
	sc := config.New()
	//logger := zerolog.New(os.Stdout)
	zerolog.SetGlobalLevel(sc.LogLevel)

	// подключаемся к базе данных
	dbAdrr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.DbUsername, sc.DbPassword, sc.DbHost, sc.DbPort, sc.DbName)
	db, err := pgxpool.New(context.Background(), dbAdrr)
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to database")
	}
	log.Info().Msg("connected to database")
	defer db.Close()

	cloneDB := database.NewImdbCloneDB(db)                 // подключаем бд
	cloneLogic := logic.NewImdbCloneLogic(cloneDB)         // подключаем бд к логике...
	cloneServer := delivery.NewImdbCloneServer(cloneLogic) // ... а логику в delivery

	http.HandleFunc("/findParts", cloneServer.FindParts)
	http.HandleFunc("/findActors", cloneServer.FindActors)
	http.HandleFunc("/findMovies", cloneServer.FindMovies)
	http.HandleFunc("/addActors", cloneServer.AddActors)
	http.HandleFunc("/addMovie", cloneServer.AddMovie)
	http.HandleFunc("/addParts", cloneServer.AddParts)
	http.HandleFunc("/removeMovies", cloneServer.RemoveMovies)
	http.HandleFunc("/removeActors", cloneServer.RemoveActors)

	log.Info().Msg("server is running...")
	err = http.ListenAndServe(":8080", nil)
	log.Fatal().Err(err).Msg("http server crashed")
}
