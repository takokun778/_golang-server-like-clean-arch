package main

import (
	"context"
	"log"
	"template/ent"
)

func main() {
	database := ent.DatabaseConnect()
	defer database.Close()

	ctx := context.TODO()

	if err := database.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
}
