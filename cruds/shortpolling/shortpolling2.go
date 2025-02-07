package shortpolling

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// Car estructura del CRUD de carros
type Car struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

// Base de datos en memoria y mutex para concurrencia
var (
	carData  = make(map[int]Car)
	carMutex sync.Mutex
)

// RegisterCarRoutes registra las rutas del CRUD de carros
func RegisterCarRoutes(r *gin.Engine) {
	sp := r.Group("/shortpolling/cars")
	{
		sp.GET("/", GetAllCars)
		sp.POST("/", CreateCar)
		sp.DELETE("/:id", DeleteCar)
	}
}

// GetAllCars obtiene todos los carros (Short Polling)
func GetAllCars(c *gin.Context) {
	carMutex.Lock()
	defer carMutex.Unlock()

	cars := []Car{}
	for _, car := range carData {
		cars = append(cars, car)
	}
	c.JSON(http.StatusOK, cars)
}

// CreateCar agrega un nuevo carro
func CreateCar(c *gin.Context) {
	var newCar Car
	if err := c.BindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	carMutex.Lock()
	carData[newCar.ID] = newCar
	carMutex.Unlock()

	c.JSON(http.StatusCreated, newCar)
}

// DeleteCar elimina un carro por ID
func DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	carMutex.Lock()
	delete(carData, id)
	carMutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Carro eliminado"})
}
