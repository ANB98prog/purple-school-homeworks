package repository

import "time"

type Order struct {
	Id     uint
	UserID uint
	Date   time.Time
	Total  float64
	Items  []OrderItem
}

type OrderItem struct {
	Id        uint
	OrderID   uint
	ProductID uint
	Quantity  uint
	Price     float64
}
