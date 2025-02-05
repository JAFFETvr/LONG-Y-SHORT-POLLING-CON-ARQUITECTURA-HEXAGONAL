package infraestructure

import (
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	CreateProductController *CreateProductController
	GetProductsController   *GetProductsController
	UpdateProductController *UpdateProductController
	DeleteProductController *DeleteProductController
	PollingController       *PollingController
}

func NewProductRoutes(
	cpc *CreateProductController,
	gpc *GetProductsController,
	upc *UpdateProductController,
	dpc *DeleteProductController,
	pc *PollingController,
) *ProductRoutes {
	return &ProductRoutes{
		CreateProductController: cpc,
		GetProductsController:   gpc,
		UpdateProductController: upc,
		DeleteProductController: dpc,
		PollingController:       pc,
	}
}

func (pr *ProductRoutes) SetupRoutes(router *gin.Engine) {
	router.POST("/products", pr.CreateProductController.Execute)
	router.GET("/products", pr.GetProductsController.Execute)
	router.PUT("/products/:id", pr.UpdateProductController.Execute)
	router.DELETE("/products/:id", pr.DeleteProductController.Execute)

	// Rutas de Short y Long Polling
	router.GET("/short-polling", pr.PollingController.ShortPollingHandler)
	router.GET("/long-polling", pr.PollingController.LongPollingHandler) // Si necesitas long-polling, implementarlo también
	router.POST("/notify-update", pr.PollingController.NotifyDataUpdateHandler) // Si necesitas notificación de actualización
}
