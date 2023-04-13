package conf

import "os"

var instance *Data

type Data struct {
	Port            string      `json:"port,omitempty"`
	GinMode         string      `json:"gin_mode,omitempty"`
	TestScope       bool        `json:"test_scope,omitempty"`
	LoggingLevel    string      `json:"logging_level,omitempty"`
	ConfigurationDB *DBSettings `json:"db_config,omitempty"`
}

type DBSettings struct {
	Username           string
	Password           string
	Host               string
	Name               string
	MaxIdleConnections int
	MaxOpenConnections int
	NetProtocol        string
	ConnectionTimeout  int
}

func GetData() *Data {
	if instance == nil {
		ymlConfig := getYMLNewConfig()

		port := os.Getenv("PORT")
		env := os.Getenv("GIN_MODE")
		test := true

		if len(port) == 0 {
			port = "8080"
		}

		instance = &Data{
			Port:      port,
			GinMode:   env,
			TestScope: test,
			ConfigurationDB: &DBSettings{
				Username:           ymlConfig.GetString("db.username"),
				Password:           os.Getenv("DB_PASS"),
				Host:               os.Getenv("DB_HOST"),
				Name:               ymlConfig.GetString("db.name"),
				MaxIdleConnections: ymlConfig.GetInt("db.max-idle-connections"),
				MaxOpenConnections: ymlConfig.GetInt("db.max-open-connections"),
				NetProtocol:        ymlConfig.GetString("db.net-protocol"),
				ConnectionTimeout:  ymlConfig.GetInt("db.connection-timeout"),
			},
		}
	}

	return instance
}
