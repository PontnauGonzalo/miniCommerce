package main

import (
	"encoding/json"
	"fmt"
	"github.com/PontnauGonzalo/miniCommerce/commerce"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Adress   string `json:"adress"`
	Age      uint8  `json:"age"`
	LastName string `json:"last_name"`
}

func (user User) Display() {
	data, _ := json.Marshal(user)
	fmt.Printf(string(data))
}

func (user *User) UpdateName(name string) {
	user.Name = name
}

func main() {

	user := User{
		ID:       1,
		Name:     "Gonzalo",
		Adress:   "Calle Falsa 123",
		Age:      26,
		LastName: "Pontnau",
	}

	user.Display()
	user.UpdateName("Gonzalo Pontnau")
	user.Display()

	product1 := commerce.Product{
		ID:    1,
		Name:  "Product 1",
		Type:  commerce.Type{ID: 1, Code: "A", Description: "Air frier"},
		Count: 2,
		Price: 10.5,
		Tags:  []string{"tag1", "tag2"},
	}
	product2 := commerce.Product{
		ID:    2,
		Name:  "Product 2",
		Type:  commerce.Type{ID: 2, Code: "B", Description: "Blender"},
		Count: 1,
		Price: 20.5,
		Tags:  []string{"tag3", "tag4"},
	}

	// Create a new car and add products
	car := commerce.NewCar(1)
	car.AddProduct(product1, product2)

	fmt.Println()
	fmt.Printf("Total products: %d\n", len(car.Products))
	fmt.Printf("Total car: %.2f\n", car.TotalPrice())
}
