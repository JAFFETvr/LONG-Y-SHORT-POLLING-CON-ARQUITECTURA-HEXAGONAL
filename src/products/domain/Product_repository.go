package domain

import "demo/src/products/domain/entities"

type IProduct interface {
	GetAllProducts()([]entities.Product,error)
	SaveProductWithParams(product *entities.Product)
	DeleteById(id int) error
	UpdateById(id int, updatedProduct *entities.Product) error
}
