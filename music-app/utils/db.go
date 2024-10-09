package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// Объявляем глобальную переменную для соединения
var conn *pgx.Conn

// ConnectDB - функция для подключения к базе данных
func ConnectDB() {
	var err error
	// Получаем URL подключения из переменных окружения
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_NAME", "musicdb"),
	)

	// Подключаемся к базе данных
	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v\n", err)
	}
	fmt.Println("Успешно подключено к базе данных!")
}

// GetDB - функция для получения соединения с базой данных
func GetDB() *pgx.Conn {
	return conn
}

// CloseDB - функция для закрытия соединения с базой данных
func CloseDB() {
	if conn != nil {
		err := conn.Close(context.Background())
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных: %v\n", err)
		}
	}
}
