package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT      string
	GRPC_PORT      string
	DB_HOST        string
	DB_PORT        int
	DB_USER        string
	DB_NAME        string
	DB_PASSWORD    string
	ACCESS_TOKEN   string
	REFRESH_TOKEN  string
	KAFKA_TOPIC    string
	KAFKA_GROUP_ID string
	KAFKA_BROKERS  string
}

func coalasce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}

func Load() Config {
	cfg := Config{}
	cfg.HTTP_PORT = cast.ToString(coalasce("HTTP_PORT", ":8081"))
	cfg.GRPC_PORT = cast.ToString(coalasce("GRPC_PORT", "car_sevice:50051"))
	cfg.DB_HOST = cast.ToString(coalasce("DB_HOST", "car_sevice"))
	cfg.DB_PORT = cast.ToInt(coalasce("DB_PORT", 5432))
	cfg.DB_USER = cast.ToString(coalasce("DB_USER", "postgres"))
	cfg.DB_NAME = cast.ToString(coalasce("DB_NAME", "auth"))
	cfg.DB_PASSWORD = cast.ToString(coalasce("DB_PASSWORD", "1918"))
	cfg.ACCESS_TOKEN = cast.ToString(coalasce("ACCESS_TOKEN", "my_secret_key"))
	cfg.REFRESH_TOKEN = cast.ToString(coalasce("REFRESH_TOKEN", "my_secret_key"))

	cfg.KAFKA_TOPIC = cast.ToString(coalasce("KAFKA_TOPIC", "booking_topic"))
	cfg.KAFKA_GROUP_ID = cast.ToString(coalasce("KAFKA_GROUP_ID", "booking_group"))
	cfg.KAFKA_BROKERS = cast.ToString(coalasce("KAFKA_BROKERS", "kafka:9092"))

	return cfg
}
