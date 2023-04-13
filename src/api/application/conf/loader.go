package conf

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

func LoadYMLConfiguration() {
	// TODO logs
	confPath := confPath()
	fmt.Printf("CONFPATH IS: %+v", confPath)      // log info
	_ = getDefaultConfigFile("default", confPath) // set default value.
	fmt.Print("Default Config file Loaded")       // log info
	envScope := fmt.Sprintf("%s_%s", "fury", os.Getenv("SCOPE"))
	fmt.Printf("ENVSCOPE IS: %+v", envScope) // log info
	absPath := path.Join(confPath, asYaml(envScope))
	fmt.Printf("absPATH IS: %+v", absPath) // log info
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("config file not found %v", envScope) // log Warn
	}
	_ = mergeConfigFiles(envScope, confPath)

	fmt.Printf("conf loaded: %+v", viper.GetViper()) // log info
}

func getDefaultConfigFile(configName string, confPath string) error {
	viper.SetConfigName(configName) // set configName's value as default.
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error loading conf. err= %+v", err) // log error
		return err
	}
	return nil
}

func mergeConfigFiles(file string, confPath string) error {
	fmt.Printf("Merging file : %+v", file) // log info
	viper.SetConfigName(file)
	viper.AddConfigPath(confPath)
	err := viper.MergeInConfig()
	if err != nil {
		fmt.Printf("error merging conf. err= %+v", err) // log error
		return err
	}
	fmt.Printf("Finished merging file : %+v", file) // log info
	return nil
}

func confPath() string {
	confDir := os.Getenv("CONF_DIR")
	if len(confDir) == 0 {
		fmt.Print("CONF_DIR env variable was not set") // log info
		base, _ := os.Getwd()
		confDir = path.Join(base, "conf")
		fmt.Printf("Using %s as CONF_DIR", confDir) // log info
	}
	fmt.Printf("reading conf from: %s", confDir) // log info
	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		fmt.Printf("Not able to find configuration dir. Please check your CONF_DIR environment variable %v", err) // log panic
		panic("Not able to find configuration dir. Please check your CONF_DIR environment variable")
	}
	fmt.Printf("finished read conf from: %s", confDir) // log info
	return confDir
}

func asYaml(file string) string {
	return fmt.Sprintf("%s.yml", file)
}
