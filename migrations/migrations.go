package migrations

import (
	"database/sql"
)

func MigrateUp(db *sql.DB) error {
	// 001_create_users_table.go
	if err := up_001(db); err != nil {
		return err
	}

	// 002_create_games_table.go
	if err := up_002(db); err != nil {
		return err
	}

	// 003_create_game_permissions_table.go
	if err := up_003(db); err != nil {
		return err
	}

	// 004_create_game_resource_usage_table.go
	if err := up_004(db); err != nil {
		return err
	}

	return nil
}

func MigrateDown(db *sql.DB) error {
	// ... reverse order of migrations

	// 004_create_game_resource_usage_table.go
	if err := down_004(db); err != nil {
		return err
	}

	// 003_create_game_permissions_table.go
	if err := down_003(db); err != nil {
		return err
	}

	// 002_create_games_table.go
	if err := down_002(db); err != nil {
		return err
	}

	// 001_create_users_table.go
	if err := down_001(db); err != nil {
		return err
	}

	return nil
}
