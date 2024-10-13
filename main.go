package commerce

import (
	"encoding/json"
	"fmt"
	"github.com/PontnauGonzalo/miniCommerce"
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
}
