package services

import (
	"errors"
	client "ucc/clients/product_client"
)

// Definimos la estructura que guarda la informacion del servicio, y es la que usamos para implementar la interfaz
type productService struct {
	productClient client.ProductClientInterface
}

// Definimos la interfaz de servicio
type productServiceInterface interface {
	GetProductId() (int, error)
	GetProductName() (string, error)
}

// ProductService Variable global de tipo productServiceInterface
var (
	ProductService productServiceInterface
)

// initProductService crea una nueva instancia del productService struct.
// Toma un productClient como parametro y lo guarda en el campo productClient.
// Devuelve una instancia de productServiceInterface
func initProductService(productClient client.ProductClientInterface) productServiceInterface {
	service := new(productService)
	service.productClient = productClient
	return service
}

// Usamos esta funcion para que cuando se inicialize el paquete se guarde nuestra interfaz en la variable global
func init() {
	ProductService = initProductService(client.ProductClient)
}

// GetProductId busca una ID del cliente y la analiza para ver si devuelve un error o no.
// Si la ID es -1, devuelve un error indicando que la ID no fue encontrada.
// Sino devuelve la ID
func (s *productService) GetProductId() (int, error) {

	id := s.productClient.GetProductId()
	if id == -1 {
		return id, errors.New("no se encontro la id")
	}
	return id, nil

}

// GetProductName GetProductId busca un Nombre del cliente y lo analiza para ver si devuelve un error o no.
// Si el nombre esta vacio devuelve un error.
// Sino devuelve el nombre
func (s *productService) GetProductName() (string, error) {
	name := s.productClient.GetProductName()
	if len(name) <= 0 {
		return name, errors.New("el nombre esta vacio")
	}
	return name, nil
}
