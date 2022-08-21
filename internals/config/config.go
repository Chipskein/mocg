package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	REMOTE_DIRECTORY      bool
	LOCAL_DIRECTORY       bool
	DEFAULT_DIRECTORY     string
	REMOTE_DIRECTORY_TYPE string //AWS || GDRIVE || none
}

var DefaultConfig = Config{REMOTE_DIRECTORY: false, LOCAL_DIRECTORY: true, DEFAULT_DIRECTORY: "", REMOTE_DIRECTORY_TYPE: "none"}
var UserConfig *Config

func LoadConfigEnviroment() {
	err := godotenv.Load("mocg.cfg")
	if err != nil {
		fmt.Println("Could not load config file using default")
		UserConfig = &DefaultConfig
	}

}
