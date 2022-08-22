package deck

import (
	"context"
	"database/sql"

	models "github.com/shadow300893/cards-api/models"
	dRepo "github.com/shadow300893/cards-api/repository"
)

// NewSQLDeckRepo returns implement of deck repository interface
func NewSQLDeckRepo(Conn *sql.DB) dRepo.DeckRepo {
	return &mysqlDeckRepo{
		Conn: Conn,
	}
}

type mysqlDeckRepo struct {
	Conn *sql.DB
}

func (m *mysqlDeckRepo) GetByID(ctx context.Context, id string) (*models.Deck, error) {

	payload := &models.Deck{}

	err := m.Conn.QueryRow("select id, shuffled, remaining from decks WHERE id = ?", id).Scan(&payload.ID, &payload.Shuffled, &payload.Remaining)

	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (m *mysqlDeckRepo) Create(ctx context.Context, d *models.Deck) (int64, error) {
	query := "Insert into decks (id, shuffled, total, remaining, created_at, updated_at) values (?,?,?,?, now(), now())"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, d.ID, d.Shuffled, d.Total, d.Total)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}

func (m *mysqlDeckRepo) Update(ctx context.Context, deck *models.Deck) (*models.Deck, error) {
	query := "Update decks set remaining=?, updated_at = ? where id=?"

	stmt, _ := m.Conn.PrepareContext(ctx, query)

	_, err := stmt.ExecContext(
		ctx,
		deck.Remaining,
		deck.UpdatedAt,
		deck.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return deck, nil
}
