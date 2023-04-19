package repository

import (
	"challenge/alerts/src/api/alerts/domain"
	"challenge/alerts/src/api/application/common"
	"challenge/alerts/src/api/application/conf"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	timeOut  = time.Second * 10
	myDomain = "Alerts"
	myLayer  = "REPOSITORY"

	GetAllQuery = `SELECT
	Type,
	Description,
	Created_at,
	Country
	FROM Alerts`

	CreateQuery = `INSERT INTO Alerts (
		Type,
		Description,
		Created_at,
		Country
		) VALUES(?,?,?,?)`

	SearchQuery = `SELECT Type, Description, Created_at, Country FROM Alerts WHERE Description LIKE ? OR Country LIKE ?`

	GetByTypeQuery = `SELECT
	Type,
	Description,
	Created_at,
	Country
	FROM Alerts
	WHERE Type = ?`

	GetMetricsQuery = `SELECT
	Country,
	COUNT(*) AS Quantity
	FROM Alerts
	WHERE YEAR(Created_at) = YEAR(CURRENT_DATE())
	AND MONTH(Created_at) = MONTH(CURRENT_DATE())
	GROUP BY Country
	ORDER BY Quantity DESC
	LIMIT 3`
)

type AlertRepository struct {
	dbClient common.DBClient
}

func NewAlertsRepository(cfg *conf.Data, dbClient common.DBClient) domain.AlertsRepository {
	return &AlertRepository{
		dbClient: dbClient,
	}
}

func (repository *AlertRepository) GetAll(ctx context.Context) ([]domain.Alert, error) {
	queryContext, cancelFunc := context.WithTimeout(ctx, timeOut)
	defer cancelFunc()
	db := repository.dbClient.GetConnection()

	rows, queryErr := db.QueryContext(queryContext, GetAllQuery)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
		return nil, fmt.Errorf(common.ErrorLog, myDomain, myLayer, errors.New("Error QueryContext in GetAll"))
	}
	defer rows.Close()

	var alerts []domain.Alert
	for rows.Next() {
		var alert domain.Alert
		err := rows.Scan(
			&alert.Type,
			&alert.Description,
			&alert.CreatedAt,
			&alert.Country,
		)
		if err != nil {
			log.Printf(common.ErrorLog, myDomain, myLayer, err)
			return nil, errors.New("Error rows.Scan in GetAll")
		}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		log.Printf(common.ErrorLog, myDomain, myLayer, rows.Err())
		return nil, errors.New("Error rows.Err in GetAll")
	}
	return alerts, nil
}

func (repository *AlertRepository) Create(ctx context.Context, alert domain.Alert) (*domain.Alert, error) {
	queryContext, cancelFunc := context.WithTimeout(ctx, timeOut)
	defer cancelFunc()
	db := repository.dbClient.GetConnection()

	_, err := db.ExecContext(queryContext, CreateQuery,
		alert.Type,
		alert.Description,
		alert.CreatedAt,
		alert.Country)

	if err != nil {
		log.Printf(common.ErrorLog, myDomain, myLayer, err)
		return nil, errors.New("Error Create alert")
	}
	return &alert, nil
}

func (repository *AlertRepository) Search(ctx context.Context, query domain.AlertSearch) ([]domain.Alert, error) {
	db := repository.dbClient.GetConnection()

	likeDescription := "%" + query.Input + "%"
	likeCountry := "%" + query.Input + "%"
	rows, queryErr := db.Query(SearchQuery, likeDescription, likeCountry)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
		return nil, errors.New("Error executing query in Search")
	}
	defer rows.Close()

	var alerts []domain.Alert
	for rows.Next() {
		var alert domain.Alert
		err := rows.Scan(
			&alert.Type,
			&alert.Description,
			&alert.CreatedAt,
			&alert.Country,
		)
		if err != nil {
			log.Printf(common.ErrorLog, myDomain, myLayer, err)
			return nil, errors.New("Error rows.Scan in Search")
		}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		log.Printf(common.ErrorLog, myDomain, myLayer, rows.Err())
		return nil, errors.New("Error rows.Err in Search")
	}
	return alerts, nil
}

func (repository *AlertRepository) GetAlertsByType(ctx context.Context, typeInput string) ([]domain.Alert, error) {
	queryContext, cancelFunc := context.WithTimeout(ctx, timeOut)
	defer cancelFunc()
	db := repository.dbClient.GetConnection()

	rows, queryErr := db.QueryContext(queryContext, GetByTypeQuery, typeInput)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
		return nil, errors.New("Error executing query in GetAlertsByType")
	}
	defer rows.Close()

	var alerts []domain.Alert
	for rows.Next() {
		var alert domain.Alert
		err := rows.Scan(
			&alert.Type,
			&alert.Description,
			&alert.CreatedAt,
			&alert.Country,
		)
		if err != nil {
			log.Printf(common.ErrorLog, myDomain, myLayer, err)
			return nil, errors.New("Error rows.Scan in GetAlertsByType")
		}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		log.Printf(common.ErrorLog, myDomain, myLayer, rows.Err())
		return nil, errors.New("Error rows.Err() in GetAlertsByType")
	}
	return alerts, nil
}

func (repository *AlertRepository) GetMetrics(ctx context.Context) ([]domain.Metrics, error) {
	queryContext, cancelFunc := context.WithTimeout(ctx, timeOut)
	defer cancelFunc()
	db := repository.dbClient.GetConnection()

	rows, queryErr := db.QueryContext(queryContext, GetMetricsQuery)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, nil
		}
		log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
		return nil, errors.New("Error executing query in GetMetrics")
	}
	defer rows.Close()

	var metrics []domain.Metrics
	for rows.Next() {
		var metric domain.Metrics
		err := rows.Scan(
			&metric.Country,
			&metric.Quantity,
		)
		if err != nil {
			log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
			return nil, errors.New("Error rows.Scan in GetMetrics")
		}

		metrics = append(metrics, metric)
	}

	if rows.Err() != nil {
		log.Printf(common.ErrorLog, myDomain, myLayer, queryErr)
		return nil, errors.New("Error rows.Err() in GetMetrics")
	}
	return metrics, nil
}
