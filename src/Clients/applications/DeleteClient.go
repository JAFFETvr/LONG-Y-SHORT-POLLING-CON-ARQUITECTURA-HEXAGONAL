package applications

import (
	"demo/src/Clients/infraestructure/repositories"
)

type DeleteClient struct {
	DB repositories.ClientRepository
}

func NewDeleteClient(db repositories.ClientRepository) *DeleteClient {
	return &DeleteClient{DB: db}
}

func (dc *DeleteClient) Execute(id int) error {
	err := dc.DB.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
