package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockProductClient implementamos la estructura con Mock
type mockProductClient struct {
	mock.Mock
}

// Definimos las funciones con la interfaz Mock
func (c *mockProductClient) GetProductId() int {
	ret := c.Called() // Guardamos en una variable todos los argumentos que pasemos en el test.
	return ret.Int(0) // Devuelve el primer return
}

func (c *mockProductClient) GetProductName() string {
	ret := c.Called()
	return ret.String(0)
}

func TestGetProductId(t *testing.T) {
	// Creamos una instancia del cliente mock
	mockClient := new(mockProductClient)

	// Le definimos que nos tiene que devolver cuando llamamos al cliente
	mockClient.On("GetProductId").Return(1)

	// Creamos una instancia del servicio con el cliente mock como parametro
	service := initProductService(mockClient)

	id, err := service.GetProductId()

	assert.Equal(t, 1, id)
	assert.Nil(t, err)
}

func TestGetProductIdWithError(t *testing.T) {
	mockClient := new(mockProductClient)
	mockClient.On("GetProductId").Return(-1)
	service := initProductService(mockClient)

	id, err := service.GetProductId()

	assert.Equal(t, -1, id)
	assert.NotNil(t, err)
}

func TestGetProductName(t *testing.T) {
	mockClient := new(mockProductClient)
	mockClient.On("GetProductName").Return("Juan")
	service := initProductService(mockClient)

	name, err := service.GetProductName()

	assert.Equal(t, "Juan", name)
	assert.Nil(t, err)
}

func TestGetProductNameWithError(t *testing.T) {
	mockClient := new(mockProductClient)
	mockClient.On("GetProductName").Return("")
	service := initProductService(mockClient)

	name, err := service.GetProductName()

	assert.Equal(t, "", name)
	assert.NotNil(t, err)
}
