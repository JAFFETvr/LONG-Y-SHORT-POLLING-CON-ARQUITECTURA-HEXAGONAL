package repositories

import (
    "database/sql"
    "demo/src/Clients/domain/entities"
    "log"
)

// ClientRepository define la interfaz para los métodos del repositorio
type ClientRepository interface {
    Save(client *entities.Client) error
    GetById(id int) (*entities.Client, error)
    GetAll() ([]entities.Client, error)
    DeleteById(id int) error
    EditById(id int, updatedClient *entities.Client) error
}

// clientRepositoryImpl es la implementación de ClientRepository
type clientRepositoryImpl struct {
    DB *sql.DB
}

// NewClientRepository crea una nueva instancia de clientRepositoryImpl
func NewClientRepository(db *sql.DB) ClientRepository {
    return &clientRepositoryImpl{DB: db}
}

func (repo *clientRepositoryImpl) Save(client *entities.Client) error {
    query := "INSERT INTO clientes (name, direccion) VALUES (?, ?)"
    _, err := repo.DB.Exec(query, client.Name, client.Direccion)
    if err != nil {
        log.Printf("[ClientRepository.Save] Error inserting client: %v", err)
        return err
    }
    log.Println("[ClientRepository.Save] Client inserted successfully")
    return nil
}

func (repo *clientRepositoryImpl) GetById(id int) (*entities.Client, error) {
    query := "SELECT id, name, direccion FROM clientes WHERE id = ?"
    row := repo.DB.QueryRow(query, id)
    var client entities.Client
    if err := row.Scan(&client.Id, &client.Name, &client.Direccion); err != nil {
        log.Printf("[ClientRepository.GetById] Error retrieving client with ID %d: %v", id, err)
        return nil, err
    }
    return &client, nil
}

func (repo *clientRepositoryImpl) GetAll() ([]entities.Client, error) {
    query := "SELECT id, name, direccion FROM clientes"
    rows, err := repo.DB.Query(query)
    if err != nil {
        log.Printf("[ClientRepository.GetAll] Error executing query: %v", err)
        return nil, err
    }
    defer rows.Close()

    var clients []entities.Client
    for rows.Next() {
        var client entities.Client
        if err := rows.Scan(&client.Id, &client.Name, &client.Direccion); err != nil {
            log.Printf("[ClientRepository.GetAll] Error scanning row: %v", err)
            return nil, err
        }
        clients = append(clients, client)
    }
    if rows.Err() != nil {
        log.Printf("[ClientRepository.GetAll] Error iterating over rows: %v", rows.Err())
        return nil, rows.Err()
    }
    log.Printf("[ClientRepository.GetAll] Successfully retrieved %d clients", len(clients))
    return clients, nil
}

func (repo *clientRepositoryImpl) DeleteById(id int) error {
    query := "DELETE FROM clientes WHERE id = ?"
    _, err := repo.DB.Exec(query, id)
    if err != nil {
        log.Printf("[ClientRepository.DeleteById] Error deleting client with ID %d: %v", id, err)
        return err
    }
    log.Printf("[ClientRepository.DeleteById] Client with ID %d deleted successfully", id)
    return nil
}

func (repo *clientRepositoryImpl) EditById(id int, updatedClient *entities.Client) error {
    query := "UPDATE clientes SET name = ?, direccion = ? WHERE id = ?"
    _, err := repo.DB.Exec(query, updatedClient.Name, updatedClient.Direccion, id)
    if err != nil {
        log.Printf("[ClientRepository.EditById] Error updating client with ID %d: %v", id, err)
        return err
    }
    log.Printf("[ClientRepository.EditById] Client with ID %d updated successfully", id)
    return nil
}
