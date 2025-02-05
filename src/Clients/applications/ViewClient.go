package applications

import (
    "demo/src/Clients/infraestructure/repositories"
    "demo/src/Clients/domain/entities"
)

type GetClient struct {
    db repositories.ClientRepository
}

func NewGetClient(db repositories.ClientRepository) *GetClient {
    return &GetClient{db: db}
}

func (gc *GetClient) Execute() ([]entities.Client, error) {
    clients, err := gc.db.GetAll()
    if err != nil {
        return nil, err
    }
    return clients, nil
}