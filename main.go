package main

import (
	"cruds/shortpolling"
	"cruds/longpolling"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rutaaaa del looong
	longpolling.RegisterRoutes(r)

	// Rutaaa del shoort
	shortpolling.RegisterRoutes1(r)

	r.Run(":8080") // Puerto principal
}
