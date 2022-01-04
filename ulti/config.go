package ulti

import (
	"database/sql"

	"github.com/spf13/viper"
)

// config DB variables
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func Init() (db *sql.DB) {
	config, _ := LoadConfig(".")
	db, _ = sql.Open(config.DBDriver, config.DBSource)
	return db
}