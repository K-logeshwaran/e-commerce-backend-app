package db

import (
	"encoding/json"
	"io"
	"time"
)

const ConStr string = "postgresql://postgres:printhello003@localhost:5432/ecom?sslmode=disable"

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  int       `json:"category_id"`
	ImageURL    string    `json:"image_url"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Customer struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Cart struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
}

func (p *Product) ToJson(w io.Writer) {
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func (p *Customer) ToJson(w io.Writer) {
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}
func (p *Order) ToJson(w io.Writer) {
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func Tostruct(v interface{}, b io.ReadCloser) error {
	err := json.NewDecoder(b).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
