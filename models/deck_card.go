package models

import "time"

type DeckCard struct {
	ID        string    `json:"-"`
	DeckID    string    `json:"-"`
	CardID    string    `json:"-"`
	IsDrawn   bool      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
