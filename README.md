
# Setting up Development Environment for a challenge-alerts

This README provides guidelines on how to configure the development environment for a challenge-alerts.


## Prerequisites
Before setting up your development environment, make sure you have the following prerequisites installed on your system:

- Golang: Download and install Golang from the official website https://golang.org/dl/
- Git: Download and install Git from the official website https://git-scm.com/downloads
- Mysql: Download and install MySql from the official website https://dev.mysql.com/downloads/mysql/


## Creating a Database and Table Using SQL
To create a database and table using SQL, follow these steps:

1. Open the MySQL and open a new script SQL

2. To create a new database, run the following command:

```sql
CREATE DATABASE challenge
```

```sql
USE challenge;
```

```sql
CREATE TABLE Alerts (
    Alert_id int NOT NULL AUTO_INCREMENT,
    Type varchar(128) NOT NULL,
    Description text,
    Created_at datetime,
    Country varchar(128) ,
    PRIMARY KEY (Alert_id)

);
```
if you prefer set data could run the next command

```sql
INSERT INTO challenge.Alerts (`Type`,Description,Created_at,Country) VALUES
     ('Red','Incremental data access','2023-04-12 00:00:00','Colombia'),
     ('Green','Default Config file Loaded','2023-04-13 00:00:00','Argentina'),
     ('test1','description1','2023-04-12 00:00:00','Chile'),
     ('test2','description2','2023-04-13 10:00:00','Mexico'),
     ('test2','description2','2023-04-13 12:00:00','Mexico'),
     ('Blue','this is a description','2023-03-27 12:00:00','Colombia'),
     ('Red','be carefull with this alert','2023-04-11 16:00:00','Mexico'),
     ('Yellow','use this alert to be carefull','2023-03-22 16:00:00','Argentina'),
     ('Orange','alert orange to set ','2023-04-09 00:00:00','Chile');
```


 ## Cloning the Repository
 To clone the repository, run the following command:

 ```bash
git clone https://github.com/dferrin-meli/challenge-alerts.git
```

## Configuring Environment Variables

challenge-alerts requires some environment variables to be set up before it can be used.

### DB_HOST
The DB_HOST environment variable should be set with the IP from your database is. could be _localhost_

 ```bash
export DB_HOST=localhost
```

### DB_PASS
The DB_PASS environment variable should be set with the access database password.

 ```bash
export DB_PASS=mypassword123
```

### CONF_DIR
The CONF_DIR environment variable should be set to the directory where is the .yml file

 ```bash
export CONF_DIR=/Users/your-user-name/go/src/alerts-challenge/conf
```

## Building and Running the Project
To build the project, navigate to the directory where the _main.go_ (alerts-challenge/src/api) file is and run the following command

 ```bash
go build
```

To run the project, run the following command:

 ```bash
./<executable-name>
```

## Send Request by postman

Example table for executing requests to an endpoint.
The table includes the necessary details for making HTTP requests to an endpoint,
such as the request method, endpoint URL, and description.

| Request Method | Endpoint URL  | Description |
| ------ | ------ | ------ |
| GET | localhost:8080/ping | Check health API |
| GET | localhost:8080/challenge/alerts | List all alerts |
| POST | localhost:8080/challenge/alerts?type=Purple&description=alert purple to set &created_at=2023-04-15T00:00:00Z&country=colombia | Create Alert |
| GET | localhost:8080/challenge/alerts/search?input=de | Search Alerts by description or country with the given input|
| GET | localhost:8080/challenge/alerts/search-by-type?type=Red | Search Alerts by Type|
| GET | localhost:8080/challenge/alerts/metrics | Get alerts metrics (top 3 country with more alerts last month)|
