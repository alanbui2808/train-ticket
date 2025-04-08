package main

import (
	"github.com/alanbui/train-ticket/internal/initialize"
)

func main() {
	// gin Engine
	// r := routers.NewRouter()

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	initialize.Run()
}
