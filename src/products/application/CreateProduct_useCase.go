package application
//holaaaa
import (
	"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/repositories"
)

type CreateProduct struct {
	db repositories.ProductRepository
}

func NewCreateProduct(db repositories.ProductRepository) CreateProduct {
	return CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(product *entities.Product) error {
	return cp.db.Save(product)
}
