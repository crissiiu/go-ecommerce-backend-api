package main

import "github.com/crisiiu/go-ecommerce-backend-api/internal/routers"

func main() {
	router := routers.NewRouter()
	router.Run(":8002") // listens on 0.0.0.0:8080 by default
}
