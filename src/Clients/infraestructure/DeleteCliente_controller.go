package infraestructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"demo/src/Clients/applications"
)

type DeleteClientController struct {
	UseCase *applications.DeleteClient
}

func NewDeleteClientController(useCase *applications.DeleteClient) *DeleteClientController {
	return &DeleteClientController{UseCase: useCase}
}

func (c *DeleteClientController) Handle(ctx *gin.Context) {
	// Obtener el ID del parámetro de la URL
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.UseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cliente eliminado correctamente"})
}
