package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/shadow300893/cards-api/driver"
	handler "github.com/shadow300893/cards-api/handler/http"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbName := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	dHandler := handler.NewDeckHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/decks", deckRouter(dHandler))
	})

	fmt.Println("Server listen at :8000")
	http.ListenAndServe(":8000", r)
}

// A completely separate router for deck routes
func deckRouter(dHandler *handler.Deck) http.Handler {
	r := chi.NewRouter()
	r.Get("/create", dHandler.CreateDeck)
	r.Get("/{id}", dHandler.OpenDeck)
	r.Get("/{id}/cards/draw", dHandler.DrawCards)

	return r
}
