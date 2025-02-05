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

// Verificar si un cliente existe (Short Polling)
func (c *DeleteClientController) CheckExists(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if c.isClientInDB(id) {
        ctx.JSON(http.StatusOK, gin.H{"exists": true})
    } else {
        ctx.JSON(http.StatusOK, gin.H{"exists": false})
    }
}

// Método auxiliar para verificar si un cliente existe en la base de datos
func (c *DeleteClientController) isClientInDB(id int) bool {
    client, err := c.UseCase.DB.GetById(id)
    return err == nil && client != nil
}