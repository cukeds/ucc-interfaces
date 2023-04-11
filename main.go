package main

import (
	"fmt"
	"ucc/services"
)

func main() {

	id, err := services.ProductService.GetProductId()
	name, err2 := services.ProductService.GetProductName()
	fmt.Println(id, err)
	fmt.Println(name, err2)
}
