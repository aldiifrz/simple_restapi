package repository

import (
	"errors"
	"simple_restapi/internal/product/entity"
)

type ProductRepository interface {
	GetAll() []entity.Product
	Add(product entity.Product) entity.Product
	Update(product entity.Product) error
	Delete(id int) error
	FindByID(id int) (*entity.Product, error) // Add FindByID method
}

type InMemoryProductRepository struct {
	products []entity.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []entity.Product{},
	}
}

func (repo *InMemoryProductRepository) GetAll() []entity.Product {
	return repo.products
}

func (repo *InMemoryProductRepository) Add(product entity.Product) entity.Product {
	product.ID = len(repo.products) + 1 // Assign an ID
	repo.products = append(repo.products, product)
	return product
}

func (repo *InMemoryProductRepository) Update(product entity.Product) error {
	for i, p := range repo.products {
		if p.ID == product.ID {
			repo.products[i] = product
			return nil
		}
	}
	return errors.New("product not found")
}

func (repo *InMemoryProductRepository) Delete(id int) error {
	for i, p := range repo.products {
		if p.ID == id {
			repo.products = append(repo.products[:i], repo.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

// New method to find a product by id
func (repo *InMemoryProductRepository) FindByID(id int) (*entity.Product, error) {
	for _, p := range repo.products {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}
