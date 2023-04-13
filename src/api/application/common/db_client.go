package common

import (
	"challenge/alerts/src/api/application/conf"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DBClient interface {
	GetConnection() *sql.DB
}

type dbClient struct {
	dataSource *sql.DB
}

var databaseClient *dbClient

func NewDBClient(dbSettings *conf.DBSettings) DBClient {
	if databaseClient == nil {
		dataSource, dbError := sql.Open("mysql", getDSNConnection(dbSettings))
		if dbError != nil {
			panic(dbError)
		}

		dataSource.SetMaxIdleConns(dbSettings.MaxIdleConnections)
		dataSource.SetConnMaxLifetime(time.Second)
		dataSource.SetMaxOpenConns(dbSettings.MaxOpenConnections)

		if dbError = dataSource.Ping(); dbError != nil {
			fmt.Printf("error when trying to connect database. Error: %v", dbError)
			panic(dbError)
		}

		databaseClient = &dbClient{
			dataSource: dataSource,
		}
	}

	return databaseClient
}

func (r *dbClient) GetConnection() *sql.DB {
	return r.dataSource
}

func getDSNConnection(dbSettings *conf.DBSettings) string {
	mySQLConfig := mysql.NewConfig()
	mySQLConfig.Addr = dbSettings.Host
	mySQLConfig.User = dbSettings.Username
	mySQLConfig.Passwd = dbSettings.Password
	mySQLConfig.Net = dbSettings.NetProtocol
	mySQLConfig.DBName = dbSettings.Name
	mySQLConfig.ParseTime = true
	return mySQLConfig.FormatDSN()
}
