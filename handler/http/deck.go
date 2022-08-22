package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/shadow300893/cards-api/driver"
	"github.com/shadow300893/cards-api/models"
	repository "github.com/shadow300893/cards-api/repository"
	card "github.com/shadow300893/cards-api/repository/card"
	deck "github.com/shadow300893/cards-api/repository/deck"
	deckCard "github.com/shadow300893/cards-api/repository/deckCard"
)

// NewDeckHandler ...
func NewDeckHandler(db *driver.DB) *Deck {
	return &Deck{
		repo:      deck.NewSQLDeckRepo(db.SQL),
		cRepo:     card.NewSQLCardRepo(db.SQL),
		dCardRepo: deckCard.NewSQLDeckCardRepo(db.SQL),
	}
}

// Deck ...
type Deck struct {
	repo      repository.DeckRepo
	cRepo     repository.CardRepo
	dCardRepo repository.DeckCardRepo
}

// Create a new deck
func (d *Deck) CreateDeck(w http.ResponseWriter, r *http.Request) {
	shuffled, _ := strconv.ParseBool(r.URL.Query().Get("shuffled"))
	cards := r.URL.Query().Get("cards")
	cardsCount := int8(52)
	var cardsSlice []string
	var cardsData []*models.Card

	//if cards query param present, fetch requested cards else fetch all cards
	if cards != "" {
		cardsSlice = strings.Split(cards, ",")
		cardsCount = int8(len(cardsSlice))
		cardsData, _ = d.cRepo.FetchAll(r.Context(), cardsSlice, false)
		if len(cardsData) == 0 {
			respondWithError(w, http.StatusNotFound, "Selected cards not found")
			return
		}
		if int8(len(cardsData)) != cardsCount {
			cardsCount = int8(len(cardsData))
		}
	} else {
		cardsData, _ = d.cRepo.FetchAll(r.Context(), cardsSlice, shuffled)
	}

	//create a new deck
	uuid := uuid.New().String()
	deck := models.Deck{
		ID:        uuid,
		Shuffled:  shuffled,
		Total:     cardsCount,
		Remaining: cardsCount,
	}

	_, err := d.repo.Create(r.Context(), &deck)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	// store cards requested for the deck
	var deckCards []*models.DeckCard
	for _, card := range cardsData {
		data := models.DeckCard{
			CardID:    card.ID,
			DeckID:    uuid,
			IsDrawn:   false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		deckCards = append(deckCards, &data)
	}
	_, err = d.dCardRepo.Create(r.Context(), deckCards)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, deck)
}

// GetByID returns a deck details
func (d *Deck) OpenDeck(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	//fetch deck by deck id
	deck, err := d.repo.GetByID(r.Context(), string(id))

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Deck not found")
		return
	}

	//fetch all not drawn cards for a deck
	cards, err := d.dCardRepo.FetchAll(r.Context(), deck.ID, 0)
	deckResponse := models.DeckResponse{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
		Card:      cards,
	}

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Deck not found")
		return
	}

	respondwithJSON(w, http.StatusOK, deckResponse)
}

// Draw cards from a deck
func (d *Deck) DrawCards(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	//fetch a deck by id
	deck, err := d.repo.GetByID(r.Context(), string(id))
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	if count == 0 {
		count = 1
	}

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Deck not found")
		return
	}

	if deck.Remaining == 0 {
		respondWithError(w, http.StatusNoContent, "Deck is empty")
		return
	}

	if deck.Remaining < int8(count) {
		respondWithError(w, http.StatusOK, fmt.Sprintf("Deck doesn't have %d cards", count))
		return
	}

	//fetch count no of cards (defaults to 1) from deck
	cardsData, err := d.dCardRepo.FetchAll(r.Context(), deck.ID, count)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Cards not found in deck")
		return
	}

	//mark above fetched cards as drawn
	for _, card := range cardsData {
		data := models.DeckCard{
			DeckID:    deck.ID,
			CardID:    card.ID,
			IsDrawn:   true,
			UpdatedAt: time.Now(),
		}
		_, err = d.dCardRepo.Update(r.Context(), &data)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error while drawing cards")
			return
		}
	}

	// update remaining cards value in deck
	deckData := models.Deck{
		ID:        deck.ID,
		Remaining: deck.Remaining - int8(count),
		UpdatedAt: time.Now(),
	}
	_, err = d.repo.Update(r.Context(), &deckData)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error while drawing cards")
		return
	}

	respondwithJSON(w, http.StatusOK, map[string]interface{}{"cards": cardsData})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
