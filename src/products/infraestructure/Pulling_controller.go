package infraestructure

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PollingController struct {
	PollingUseCase *application.PollingUseCase
}

func NewPollingController(pu *application.PollingUseCase) *PollingController {
	return &PollingController{PollingUseCase: pu}
}

// Short Polling: Devuelve los productos actuales
func (pc *PollingController) ShortPollingHandler(c *gin.Context) {
	// Llamar a ShortPolling y manejar el error
	data, err := pc.PollingUseCase.ShortPolling()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// Long Polling: Espera hasta que haya cambios o timeout
func (pc *PollingController) LongPollingHandler(c *gin.Context) {
	data, updated := pc.PollingUseCase.LongPolling()
	if updated {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusRequestTimeout, data)
	}
}

// NotifyDataUpdate: Notifica cambios en los datos
func (pc *PollingController) NotifyDataUpdateHandler(c *gin.Context) {
	pc.PollingUseCase.NotifyDataUpdate()
	c.JSON(http.StatusOK, gin.H{"message": "Datos actualizados"})
}
