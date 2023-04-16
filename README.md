# challenge-alerts


# Setting up Development Environment for a challenge-alerts

This README provides guidelines on how to configure the development environment for a challenge-alerts.


## Prerequisites
Before setting up your development environment, make sure you have the following prerequisites installed on your system:

- Golang: Download and install Golang from the official website https://golang.org/dl/
- Git: Download and install Git from the official website https://git-scm.com/downloads
- Mysql: Download and install MySql from the official website https://dev.mysql.com/downloads/mysql/

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
