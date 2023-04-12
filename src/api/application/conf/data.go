package conf

import "os"

var instance *Data

type Data struct {
	Port         string `json:"port,omitempty"`
	GinMode      string `json:"gin_mode,omitempty"`
	TestScope    bool   `json:"test_scope,omitempty"`
	LoggingLevel string `json:"logging_level,omitempty"`
}

func GetData() *Data {
	if instance == nil {
		// ymlConfig := getYMLNewConfig()

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
		}
	}

	return instance
}
