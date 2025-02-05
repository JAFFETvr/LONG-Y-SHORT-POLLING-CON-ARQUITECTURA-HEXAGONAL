package application

import (
	"demo/src/products/infraestructure/repositories"
	"time"
)

type PollingUseCase struct {
	productRepo     repositories.ProductRepository
	lastUpdatedTime time.Time
	waitingClients  []chan string
}

func NewPollingUseCase(productRepo repositories.ProductRepository) *PollingUseCase {
	return &PollingUseCase{
		productRepo:    productRepo,
		lastUpdatedTime: time.Now(),
		waitingClients:  make([]chan string, 0),
	}
}

// Short Polling: Retorna los productos actuales
func (p *PollingUseCase) ShortPolling() (map[string]interface{}, error) {
	// Obtener los productos actuales desde el repositorio
	products, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Devolver los productos en el formato deseado
	return map[string]interface{}{
		"message": "Productos actuales",
		"products": products,
		"time": time.Now().Format(time.RFC3339),
	}, nil
}

// Long Polling: Espera hasta que haya cambios o timeout
func (p *PollingUseCase) LongPolling() (map[string]string, bool) {
	updateChannel := make(chan string)
	p.waitingClients = append(p.waitingClients, updateChannel)

	select {
	case <-updateChannel:
		return map[string]string{
			"message": "Datos actualizados en Long Polling",
			"time":    time.Now().Format(time.RFC3339),
		}, true
	case <-time.After(30 * time.Second):
		return map[string]string{
			"message": "No hubo cambios en los datos",
		}, false
	}
}

// NotifyDataUpdate: Notifica cambios en los datos
func (p *PollingUseCase) NotifyDataUpdate() {
	// Notificar a todos los clientes en espera
	for _, ch := range p.waitingClients {
		close(ch)
	}
	p.waitingClients = nil
}
