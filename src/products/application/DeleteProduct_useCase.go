package application

import (
	"demo/src/products/infraestructure/repositories"
)

type DeleteProduct struct {
	db repositories.ProductRepository
}

func NewDeleteProduct(db repositories.ProductRepository) *DeleteProduct {
	return &DeleteProduct{db: db}
}

func (dp *DeleteProduct) Execute(id int) error {
	err := dp.db.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
