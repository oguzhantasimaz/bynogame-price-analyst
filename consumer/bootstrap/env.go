package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	KafkaUsername  string `mapstructure:"KAFKA_USERNAME"`
	KafkaPassword  string `mapstructure:"KAFKA_PASSWORD"`
	KafkaBrokerUrl string `mapstructure:"KAFKA_BROKER_URL"`
	CoreServiceUrl string `mapstructure:"CORE_SERVICE_URL"`
	SteamWebApiKey string `mapstructure:"STEAMWEBAPI_KEY"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
