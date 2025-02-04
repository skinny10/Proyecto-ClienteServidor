package longpolling

import (
	"net/http"
	"sync"

	"cruds/utils" // Importar paquete utils
	"github.com/gin-gonic/gin"
)

var (
	subscribers []chan []utils.Item
	mutex       sync.Mutex
)

// notifica a los clientes de long polling
func notifySubscribers() {
	mutex.Lock()
	defer mutex.Unlock()

	items := utils.GetAllItems(utils.DataCrud1)
	for _, sub := range subscribers {
		sub <- items
	}
	subscribers = nil
}

// Maneja las suscripciones de long polling
func longPollingHandler(c *gin.Context) {
	ch := make(chan []utils.Item)

	mutex.Lock()
	subscribers = append(subscribers, ch)
	mutex.Unlock()

	items := <-ch
	c.JSON(http.StatusOK, items)
}

// Registra las rutas de long polling
func RegisterRoutes(router *gin.Engine) {
	router.GET("/longpolling/items", longPollingHandler)
}
