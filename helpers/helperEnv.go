package helpers

import (
	"log"

	"github.com/spf13/viper"
)

//GetKeyValue is a function for get a value from a key from .env file
func GetKeyValue(key, defaultValue string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		return defaultValue
	}
	
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion %v", key)
	}

	return value
}
