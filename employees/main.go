package main

import (
	"employees/httpserver"
	"fmt"
	"os"
)

func main() {

	fmt.Println("сервер запущен на :5050")
	if err := httpserver.StartHttpServer(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
