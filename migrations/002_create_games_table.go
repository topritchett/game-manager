package migrations

import (
	"database/sql"
)

func up_002(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS games (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			price FLOAT NOT NULL,
			created_by INTEGER NOT NULL,
			permission_level INTEGER NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (created_by) REFERENCES users(id)
		)
    `)
	return err
}

func down_002(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS games")
	return err
}
