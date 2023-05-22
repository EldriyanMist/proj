package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectToDB устанавливает соединение с базой данных PostgreSQL
func ConnectToDB(host, port, user, password, dbname string) (*sql.DB, error) {
	// Формируем строку подключения
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// Возвращаем работающее соединение к базе данных
	return db, nil
}
