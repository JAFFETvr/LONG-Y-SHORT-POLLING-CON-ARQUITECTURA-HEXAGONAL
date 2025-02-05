package infraestructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"demo/src/Clients/applications"
	"demo/src/Clients/domain/entities"
)

type UpdateClientController struct {
	UseCase *applications.UpdateClient
}

func NewUpdateClientController(useCase *applications.UpdateClient) *UpdateClientController {
	return &UpdateClientController{UseCase: useCase}
}

func (c *UpdateClientController) Handle(ctx *gin.Context) {
	// Obtener el ID del cliente desde los parámetros de la URL
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updatedClient entities.Client
	if err := ctx.ShouldBindJSON(&updatedClient); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UseCase.Execute(id, &updatedClient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente actualizado correctamente"})
}
