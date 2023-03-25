package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/topritchett/game-manager/games"
	"github.com/topritchett/game-manager/migrations"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db/gameserver?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = migrations.MigrateUp(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration successful!")
	games.CreateGame("game1")
}
