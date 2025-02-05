package domain
import "demo/src/Clients/domain/entities"

type ClientRepository interface {
    Create(client *entities.Client) error
    GetByID(id uint) (*entities.Client, error)
    GetAll() ([]entities.Client, error)
    Update(client *entities.Client) error
    Delete(id uint) error
}