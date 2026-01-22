package repository

import "gorm.io/gorm"

type OrderRepository interface {
	Create(order *Order) (*Order, error)
	GetById(id uint) (*Order, error)
	GetUserOrders(userId uint) ([]Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

var _ OrderRepository = (*orderRepository)(nil)

// Create - создает заказ
func (repo *orderRepository) Create(order *Order) (*Order, error) {
	result := repo.db.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

// GetById - возвращает заказ по идентификатору
func (repo *orderRepository) GetById(id uint) (*Order, error) {
	order := &Order{}
	result := repo.db.First(&order, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

// GetUserOrders - возвращает все заказы пользователя
func (repo *orderRepository) GetUserOrders(userId uint) ([]Order, error) {
	var orders []Order
	result := repo.db.Find(&orders, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
