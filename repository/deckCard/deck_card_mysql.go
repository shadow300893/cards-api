package deckCard

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	models "github.com/shadow300893/cards-api/models"
	repo "github.com/shadow300893/cards-api/repository"
)

// NewSQLDeckCardRepo returns implement of deck cards repository interface
func NewSQLDeckCardRepo(Conn *sql.DB) repo.DeckCardRepo {
	return &mysqlDeckCardRepo{
		Conn: Conn,
	}
}

type mysqlDeckCardRepo struct {
	Conn *sql.DB
}

func (m *mysqlDeckCardRepo) Create(ctx context.Context, deckCards []*models.DeckCard) (int64, error) {
	var err error
	var (
		placeholders []string
		vals         []interface{}
	)
	for _, data := range deckCards {
		placeholders = append(placeholders, "(?, ?, ?, now(), now())")
		vals = append(vals, data.DeckID, data.CardID, data.IsDrawn)
	}

	query := fmt.Sprintf("INSERT INTO deck_cards(deck_id, card_id, is_drawn, created_at, updated_at) VALUES %s", strings.Join(placeholders, ","))
	stmt, _ := m.Conn.PrepareContext(ctx, query)
	res, err := stmt.ExecContext(ctx, vals...)
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return res.RowsAffected()
}

func (m *mysqlDeckCardRepo) Update(ctx context.Context, deckCard *models.DeckCard) (*models.DeckCard, error) {
	query := "Update deck_cards set is_drawn=?, updated_at = ? where card_id=? and deck_id = ?"

	stmt, _ := m.Conn.PrepareContext(ctx, query)

	_, err := stmt.ExecContext(
		ctx,
		deckCard.IsDrawn,
		deckCard.UpdatedAt,
		deckCard.CardID,
		deckCard.DeckID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return deckCard, nil
}

func (m *mysqlDeckCardRepo) FetchAll(ctx context.Context, deckID string, limit int) ([]*models.Card, error) {

	var rows *sql.Rows
	var err error
	query := "select c.id,c.suit,c.value,c.code from cards c join deck_cards dc on c.id=dc.card_id where dc.deck_id = ? and dc.is_drawn = 0 order by dc.id"

	if limit != 0 {
		query += " limit ?"
		rows, err = m.Conn.QueryContext(ctx, query, deckID, limit)
	} else {
		rows, err = m.Conn.QueryContext(ctx, query, deckID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Card, 0)
	for rows.Next() {
		data := new(models.Card)

		err := rows.Scan(
			&data.ID,
			&data.Suit,
			&data.Value,
			&data.Code,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
