package models

import "time"

type Deck struct {
	ID        string    `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Total     int8      `json:"-"`
	Remaining int8      `json:"remaining"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type DeckResponse struct {
	ID        string  `json:"deck_id"`
	Shuffled  bool    `json:"shuffled"`
	Remaining int8    `json:"remaining"`
	Card      []*Card `json:"cards"`
}
