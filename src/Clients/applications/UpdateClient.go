package applications

import (
    "demo/src/Clients/infraestructure/repositories"
    "demo/src/Clients/domain/entities"
)

type UpdateClient struct {
    db repositories.ClientRepository
}

func NewUpdateClient(db repositories.ClientRepository) *UpdateClient {
    return &UpdateClient{db: db}
}

func (uc *UpdateClient) Execute(id int, updatedClient *entities.Client) error {
    err := uc.db.EditById(id, updatedClient)
    if err != nil {
        return err
    }
    return nil
}
