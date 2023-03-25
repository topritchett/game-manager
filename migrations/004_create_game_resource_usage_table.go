package migrations

import (
	"database/sql"
)

func up_004(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS game_resource_usage (
			id SERIAL PRIMARY KEY,
			game_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			cpu_usage FLOAT NOT NULL,
			memory_usage FLOAT NOT NULL,
			network_usage FLOAT NOT NULL,
			timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
			FOREIGN KEY (game_id) REFERENCES games(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
    `)
	return err
}

func down_004(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS game_resource_usage")
	return err
}
