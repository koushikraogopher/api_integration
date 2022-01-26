package main

import (
	"fmt"

	"api_integration/routers"
)

func main() {

	// Loading routes
	router := routers.NewRouter()

	// Start listening and serving requests
	router.Run(":8080")

	fmt.Println("router:", router)
}
