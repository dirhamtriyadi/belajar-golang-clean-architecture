package sqlite

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang_clean_architecture")
	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTO_INCREMENT,
			name TEXT,
			price REAL
			);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
