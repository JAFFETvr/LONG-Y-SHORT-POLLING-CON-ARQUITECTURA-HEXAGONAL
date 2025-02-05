package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	useCase_up *application.UpdateProduct
}

func NewUpdateProductController(useCase_up *application.UpdateProduct) *UpdateProductController {
	return &UpdateProductController{useCase_up: useCase_up}
}

func (upc *UpdateProductController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updatedProduct entities.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	err = upc.useCase_up.Execute(id, &updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": updatedProduct})
}
