package repsitory

import (
	"context"

	"github.com/shadow300893/cards-api/models"
)

// PostRepo interface...
type DeckRepo interface {
	GetByID(ctx context.Context, id string) (*models.Deck, error)
	Create(ctx context.Context, p *models.Deck) (int64, error)
	Update(ctx context.Context, p *models.Deck) (*models.Deck, error)
}

// CardRepo interface...
type CardRepo interface {
	FetchAll(ctx context.Context, inArray []string, shuffle bool) ([]*models.Card, error)
}

// DeckCardRepo interface...
type DeckCardRepo interface {
	Create(ctx context.Context, deckCards []*models.DeckCard) (int64, error)
	Update(ctx context.Context, deckCard *models.DeckCard) (*models.DeckCard, error)
	FetchAll(ctx context.Context, deckID string, limit int) ([]*models.Card, error)
}
