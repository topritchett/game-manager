package migrations

import (
	"database/sql"
)

func up_003(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS game_permissions (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			game_id INTEGER NOT NULL,
			permission_level INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (game_id) REFERENCES games(id)
		)
    `)
	return err
}

func down_003(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS game_permissions")
	return err
}
