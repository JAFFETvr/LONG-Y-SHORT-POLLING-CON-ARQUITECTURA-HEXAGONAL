package infraestructure

import (
    "github.com/gin-gonic/gin"
    "demo/src/Clients/applications"
    "demo/src/Clients/domain/entities"
    "net/http"
)

type CreateClientController struct {
    UseCase *applications.CreateClient
}

func NewCreateClientController(useCase *applications.CreateClient) *CreateClientController {
    return &CreateClientController{UseCase: useCase}
}

func (c *CreateClientController) Handle(ctx *gin.Context) {
    var client entities.Client
    if err := ctx.ShouldBindJSON(&client); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.UseCase.Execute(&client); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, client)
}
