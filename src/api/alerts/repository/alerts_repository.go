package repository

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/application/conf"
	"context"
	"database/sql"
	"errors"
	"time"
)

const (
	timeOut = time.Second * 10

	GetAllQuery = `SELECT
	Type,
	Description,
	Created_at,
	Country
	FROM Alerts`
)

type AlertRepository struct {
	dbClient common.DBClient
}

func NewAlertsRepository(cfg *conf.Data, dbClient common.DBClient) domain.AlertsRepository {
	return &AlertRepository{
		dbClient: dbClient,
	}
}

func (repository *AlertRepository) GetAll(ctx context.Context) ([]domain.Alerts, error) {
	queryContext, cancelFunc := context.WithTimeout(ctx, timeOut)
	defer cancelFunc()
	db := repository.dbClient.GetConnection()

	rows, queryErr := db.QueryContext(queryContext, GetAllQuery)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, queryErr
	}
	defer rows.Close()

	var alerts []domain.Alerts
	for rows.Next() {
		var alert domain.Alerts
		err := rows.Scan(
			&alert.Type,
			&alert.Description,
			&alert.CreatedAt,
			&alert.Country,
		)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return alerts, nil
}
