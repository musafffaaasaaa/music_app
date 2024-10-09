package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv загружает переменные окружения из .env файла
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}
}

// GetEnv получает значение переменной окружения или возвращает значение по умолчанию, если переменная не задана.
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
