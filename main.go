package main

import (
	"demo/src/products/application"
	"demo/src/products/infraestructure"
	"demo/src/products/infraestructure/repositories"

	clients_infraestructure "demo/src/Clients/infraestructure"
	clients_repositories "demo/src/Clients/infraestructure/repositories"
	clients_application "demo/src/Clients/applications"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Iniciar la conexión a MySQL
	mysql := infraestructure.NewMySQL()
	defer mysql.Close()

	// Repositorios de productos
	productRepo := repositories.NewProductRepository(mysql.DB)

	// Casos de uso de productos
	createProduct := application.NewCreateProduct(*productRepo)
	getProducts := application.NewGetProducts(*productRepo)
	updateProduct := application.NewUpdateProduct(*productRepo)
	deleteProduct := application.NewDeleteProduct(*productRepo)

	// Controladores de productos
	createProductController := infraestructure.NewCreateProductController(createProduct)
	getProductsController := infraestructure.NewGetProductsController(getProducts)
	updateProductController := infraestructure.NewUpdateProductController(updateProduct)
	deleteProductController := infraestructure.NewDeleteProductController(deleteProduct)

	// Repositorios de clientes
	clientRepo := clients_repositories.NewClientRepository(mysql.DB)

	// Casos de uso de clientes
	createClient := clients_application.NewCreateClient(clientRepo)
	getClients := clients_application.NewGetClient(clientRepo)
	updateClient := clients_application.NewUpdateClient(clientRepo)
	deleteClient := clients_application.NewDeleteClient(clientRepo)

	// Controladores de clientes
	createClientController := clients_infraestructure.NewCreateClientController(createClient)
	getClientsController := clients_infraestructure.NewGetClientsController(getClients)
	updateClientController := clients_infraestructure.NewUpdateClientController(updateClient)
	deleteClientController := clients_infraestructure.NewDeleteClientController(deleteClient)

	// Crear el PollingUseCase (no necesita repositorio de productos)
	pollingUseCase := application.NewPollingUseCase(*productRepo) // Se pasa el repositorio de productos aquí
	pollingController := infraestructure.NewPollingController(pollingUseCase)

	// Configuración de las rutas
	router := gin.Default()

	// Rutas de productos
	productRoutes := infraestructure.NewProductRoutes(
		createProductController,
		getProductsController,
		updateProductController,
		deleteProductController,
		pollingController, // Controlador de polling
	)
	productRoutes.SetupRoutes(router)

	// Rutas de clientes
	clientsRoutes := clients_infraestructure.NewClientRoutes(
		createClientController,
		getClientsController,
		updateClientController,
		deleteClientController,
	)
	clientsRoutes.SetupRoutes(router)

	// Iniciar el servidor
	log.Println("[Main] Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
