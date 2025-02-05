package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type ClientRoutes struct {
	CreateClientController *CreateClientController
	GetClientsController   *GetClientsController
	UpdateClientController *UpdateClientController
	DeleteClientController *DeleteClientController
}

func NewClientRoutes(
	createCtrl *CreateClientController,
	getCtrl *GetClientsController,
	updateCtrl *UpdateClientController,
	delCtrl *DeleteClientController,
) *ClientRoutes {
	return &ClientRoutes{
		CreateClientController: createCtrl,
		GetClientsController:   getCtrl,
		UpdateClientController: updateCtrl,
		DeleteClientController: delCtrl,
	}
}

func (cr *ClientRoutes) SetupRoutes(router *gin.Engine) {
	clients := router.Group("/clients")
	{
		clients.POST("/", cr.CreateClientController.Handle)  
		clients.GET("/", cr.GetClientsController.Handle)     
		clients.PUT("/:id", cr.UpdateClientController.Handle) 
		clients.DELETE("/:id", cr.DeleteClientController.Handle) 
	}
}
