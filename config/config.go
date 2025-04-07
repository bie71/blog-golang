package config

import "github.com/spf13/viper"

type APP struct {
	APP_PORT  string `json:"app_port"`
	APP_ENV   string `json:"app_env"`
	JwtKey    string `json:"jwt_key"`
	JwtIssuer string `json:"jwt_issuer"`
}

type PsqlDB struct {
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBDriver   string `json:"db_driver"`
	DBMaxIdle  int    `json:"db_max_idle"`
	DBMaxOpen  int    `json:"db_max_open"`
	DBMaxLife  int    `json:"db_max_life"`
}

type Config struct {
	APP APP
	DB  PsqlDB
}

func NewConfig() *Config {
	return &Config{
		APP: APP{
			APP_PORT:  viper.GetString("APP_PORT"),
			APP_ENV:   viper.GetString("APP_ENV"),
			JwtKey:    viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer: viper.GetString("JWT_ISSUER"),
		},
		DB: PsqlDB{
			DBHost:     viper.GetString("DB_HOST"),
			DBPort:     viper.GetString("DB_PORT"),
			DBUser:     viper.GetString("DB_USER"),
			DBPassword: viper.GetString("DB_PASSWORD"),
			DBName:     viper.GetString("DB_NAME"),
			DBMaxIdle:  viper.GetInt("DB_MAX_IDLE"),
			DBMaxOpen:  viper.GetInt("DB_MAX_OPEN"),
			DBMaxLife:  viper.GetInt("DB_MAX_LIFE"),
		},
	}
}
