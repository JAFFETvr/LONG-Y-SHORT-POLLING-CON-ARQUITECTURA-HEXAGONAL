package repositories

import (
	"database/sql"
	"demo/src/products/domain/entities"
	"log"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) Save(product *entities.Product) error {
	query := "INSERT INTO productos (name, price) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, product.Name, product.Price)
	if err != nil {
		log.Printf("[ProductRepository.Save] Error inserting product: %v", err)
		return err
	}
	log.Println("[ProductRepository.Save] Product inserted successfully")
	return nil
}

func (repo *ProductRepository) GetAll() ([]entities.Product, error) {
	query := "SELECT id, name, price FROM productos"
	rows, err := repo.DB.Query(query)
	if err != nil {
		log.Printf("[ProductRepository.GetAll] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			log.Printf("[ProductRepository.GetAll] Error scanning row: %v", err)
			return nil, err
		}
		products = append(products, product)
	}
	if rows.Err() != nil {
		log.Printf("[ProductRepository.GetAll] Error iterating over rows: %v", rows.Err())
		return nil, rows.Err()
	}
	log.Printf("[ProductRepository.GetAll] Successfully retrieved %d products", len(products))
	return products, nil
}

func (repo *ProductRepository) DeleteById(id int) error {
	query := "DELETE FROM productos WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		log.Printf("[ProductRepository.DeleteById] Error deleting product with ID %d: %v", id, err)
		return err
	}
	log.Printf("[ProductRepository.DeleteById] Product with ID %d deleted successfully", id)
	return nil
}

func (repo *ProductRepository) EditById(id int, updatedProduct *entities.Product) error {
	query := "UPDATE productos SET name = ?, price = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, updatedProduct.Name, updatedProduct.Price, id)
	if err != nil {
		log.Printf("[ProductRepository.EditById] Error updating product with ID %d: %v", id, err)
		return err
	}
	log.Printf("[ProductRepository.EditById] Product with ID %d updated successfully", id)
	return nil
}
