package repository

import (
	goerr "errors"
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product *Product) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id uint) error
	GetById(id uint) (*Product, error)
	GetAll() ([]*Product, error)
}

type pgProductRepository struct {
	*gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &pgProductRepository{db}
}

var _ ProductRepository = (*pgProductRepository)(nil)

func (repo *pgProductRepository) Create(product *Product) (*Product, error) {
	result := repo.DB.Create(product)
	if err := result.Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *pgProductRepository) Update(product *Product) (*Product, error) {
	result := repo.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *pgProductRepository) Delete(id uint) error {
	result := repo.DB.Delete(&Product{}, id)
	return result.Error
}

func (repo *pgProductRepository) GetById(id uint) (*Product, error) {
	var p Product
	result := repo.DB.First(&p, id)
	if result.Error != nil {
		if goerr.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &errors.ItemNotFound{Message: fmt.Sprintf("product with id %v not found", id)}
		}
		return nil, result.Error
	}

	return &p, nil
}

func (repo *pgProductRepository) GetAll() ([]*Product, error) {
	var products []*Product
	result := repo.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
