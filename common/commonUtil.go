package common

import (
	"customer/logger"
	"fmt"
	"os"

	"github.com/google/uuid"
	config "github.com/spf13/viper"
)

var log = logger.NewLogger()

func NewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

func MapErrorCode(transID string, code string, message string) (response ResponseBean) {
	response.Code = code
	response.Msg = message
	response.TransID = transID
	return
}

func LoadConfigFile(transID string) error {
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Args[1]
	}

	// API Start
	log.Info(transID, fmt.Sprintf("Server start running on %s environment configuration", env))
	config.SetConfigName(env)
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	err := config.ReadInConfig()
	if err != nil {
		errMsg := fmt.Sprintf("Read config file %s.yml occur error: %s", env, err.Error())
		log.Error(transID, errMsg)
		return err
	}

	return err
}
