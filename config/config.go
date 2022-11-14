package config

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Host                 string
		Port                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
	}
	Server struct {
		Port string
	}
	Api struct {
		Host        string
		ServiceId   string `mapstructure:"service_id"`
		Key         string
		SecretKey   string `mapstructure:"secret_key"`
		AccessKeyId string `mapstructure:"access_key_id"`
	}
}

// global variable
var C config

func ReadConfig() {
	Config := &C

	profile := os.Getenv("PROFILE")
	fmt.Println("Current Profile: ", profile)

	if profile == "prod" {
		viper.SetConfigName("config")
	} else if profile == "dev" {
		viper.SetConfigName("config.dev")
	} else {
		viper.SetConfigName("config.local")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // local file
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

	spew.Dump(C)
}
