package applications

import (
    "demo/src/Clients/infraestructure/repositories"
)

type DeleteClient struct {
    db repositories.ClientRepository
}

func NewDeleteClient(db repositories.ClientRepository) *DeleteClient {
    return &DeleteClient{db: db}
}

func (dc *DeleteClient) Execute(id int) error {
    err := dc.db.DeleteById(id)
    if err != nil {
        return err
    }
    return nil
}
