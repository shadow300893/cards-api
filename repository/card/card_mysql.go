package deck

import (
	"context"
	"database/sql"
	"strings"

	"github.com/shadow300893/cards-api/models"
	cRepo "github.com/shadow300893/cards-api/repository"
)

// NewSQLCardRepo returns implement of card repository interface
func NewSQLCardRepo(Conn *sql.DB) cRepo.CardRepo {
	return &mysqlCardRepo{
		Conn: Conn,
	}
}

type mysqlCardRepo struct {
	Conn *sql.DB
}

func (m *mysqlCardRepo) FetchAll(ctx context.Context, inArray []string, shuffle bool) ([]*models.Card, error) {

	var rows *sql.Rows
	var err error
	query := "select * from cards"
	if len(inArray) > 0 {
		args := make([]interface{}, len(inArray))
		for i, code := range inArray {
			args[i] = code
		}

		query += " where code in (?" + strings.Repeat(`,?`, len(inArray)-1) + ")"

		if shuffle {
			query += " order by RAND()"
		}
		rows, err = m.Conn.QueryContext(ctx, query, args...)
	} else {
		if shuffle {
			query += " order by RAND()"
		}
		rows, err = m.Conn.QueryContext(ctx, query)
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
