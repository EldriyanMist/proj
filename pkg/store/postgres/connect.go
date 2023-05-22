package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Cоединение с PostgreSQL
func ConnectToDB(host, port, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Подключение к БД
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Проверка
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
