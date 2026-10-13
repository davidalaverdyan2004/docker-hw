package main

import (
	"context"
	"employees/db"
	"employees/httpserver"
	"fmt"
	"os"
)

func main() {

	ctx := context.Background()
	connString := os.Getenv("CONN_STRING")

	pool, err := db.Connect(ctx, connString)
	if err != nil {
		fmt.Println("err connect", err)
		return
	}

	defer pool.Close()

	err = db.CreateTable(ctx, pool)
	if err != nil {
		fmt.Println("Create table:", err)
		return
	}

	fmt.Println("сервер запущен на :5050")
	if err := httpserver.StartHttpServer(pool); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
