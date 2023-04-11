package product_client

import (
	"math/rand"
	"time"
)

// Definimos la estructura productClient sin campos que usamos para implementar la interfaz
type productClient struct{}

// ProductClientInterface Definimos la interfaz que expone los metodos GetProductId y GetProductName
type ProductClientInterface interface {
	GetProductId() int
	GetProductName() string
}

// ProductClient Variable global de tipo ProductClientInterface
var (
	ProductClient ProductClientInterface
)

// Usamos esta funcion para que cuando se inicialize el paquete se guarde nuestra interfaz en la variable global
func init() {
	ProductClient = &productClient{}
}

// GetProductId Devuelve un número aleatorio que puede ser -1 o 0
func (s *productClient) GetProductId() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) - 1
}

// GetProductName Devuelve una cadena aleatoria que puede ser vacía o "Hello"
func (s *productClient) GetProductName() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return ""
	}
	return "Hello"
}
