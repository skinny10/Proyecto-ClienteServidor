package shortpolling

import (
	"net/http"
	"strconv"

	"cruds/utils" // aqui no se importa el reverendo paqueteeeeeee
	"github.com/gin-gonic/gin"
)

// Registro de las benditas rutas
func RegisterRoutes1(router *gin.Engine) {
	router.GET("/crud1/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.GetAllItems(utils.DataCrud1))
	})

	router.POST("/crud1/items", func(c *gin.Context) {
		var item utils.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			utils.CreateItem(utils.DataCrud1, item)
			c.JSON(http.StatusCreated, item)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.PUT("/crud1/items/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var item utils.Item
		if err := c.ShouldBindJSON(&item); err == nil {
			utils.UpdateItem(utils.DataCrud1, id, item)
			c.JSON(http.StatusOK, item)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.DELETE("/crud1/items/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		utils.DeleteItem(utils.DataCrud1, id)
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
	})
}
