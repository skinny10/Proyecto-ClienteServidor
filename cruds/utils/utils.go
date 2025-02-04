package utils

import (
	"sync"
)

// Item que sirvee de para todos los demas archivos
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Estructuras del shortt y del long
var (
	DataCrud1 = make(map[int]Item) // Datos para CRUD 1
	DataCrud2 = make(map[int]Item) // Datos para CRUD 2
	mutexCrud sync.Mutex
)

// GetAllItems aqui obtenemos los elemtos del crud que yo quiera
func GetAllItems(data map[int]Item) []Item {
	mutexCrud.Lock()
	defer mutexCrud.Unlock()

	items := []Item{}
	for _, v := range data {
		items = append(items, v)
	}
	return items
}

// CreateItem agregamos un elemnto nuevo a un crud
func CreateItem(data map[int]Item, item Item) {
	mutexCrud.Lock()
	defer mutexCrud.Unlock()
	data[item.ID] = item
}

// UpdateItem actualiza un elemento de un CRUD
func UpdateItem(data map[int]Item, id int, item Item) {
	mutexCrud.Lock()
	defer mutexCrud.Unlock()
	if _, exists := data[id]; exists {
		data[id] = item
	}
}

// DeleteItem eliminamos el elemento que yo quiera de un crud
func DeleteItem(data map[int]Item, id int) {
	mutexCrud.Lock()
	defer mutexCrud.Unlock()
	delete(data, id)
}
