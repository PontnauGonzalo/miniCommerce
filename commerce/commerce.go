package commerce

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func (product *Product) TotalPrice() float64 {
	return float64(product.Count) * product.Price
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

// AddProduct adds a product to the car
func (car *Car) AddProduct(products ...Product) {
	car.Products = append(car.Products, products...)
}

// TotalPrice returns the total price of the car
func (car Car) TotalPrice() float64 {
	var total float64
	for _, product := range car.Products {
		total += product.TotalPrice()
	}
	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
