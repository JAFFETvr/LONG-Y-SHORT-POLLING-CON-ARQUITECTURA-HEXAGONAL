package applications

import (
    "demo/src/Clients/infraestructure/repositories"
    "demo/src/Clients/domain/entities"
)

type CreateClient struct {
    db repositories.ClientRepository
}

func NewCreateClient(db repositories.ClientRepository) *CreateClient {
    return &CreateClient{db: db}
}

func (cc *CreateClient) Execute(client *entities.Client) error {
    err := cc.db.Save(client)
    if err != nil {
        return err
    }
    return nil
}