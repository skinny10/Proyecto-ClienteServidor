package main

import (
	shortpolling "hello/cruds/ShortPolling"
	"hello/cruds/longpolling"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rutaaaa del looong

	longpolling.RegisterRoutes(r)

	// Rutaaa del shoort
	shortpolling.RegisterRoutes1(r)
	shortpolling.RegisterCarRoutes(r)

	r.Run(":8080") // Puerto principal
}

//primera vez que uso Go, ya se que no esta limpio limpio el codigo
