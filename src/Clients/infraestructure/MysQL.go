package infraestructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLGORM struct {
	DB *gorm.DB
}

func NewMySQLGORM() *MySQLGORM {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[MySQLGORM] Error cargando el archivo .env: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("[MySQLGORM] La variable DATABASE_URL no está definida en el archivo .env")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[MySQLGORM] Error al conectar con la base de datos: %v", err)
	}

	log.Println("[MySQLGORM] Conexión establecida correctamente")
	return &MySQLGORM{DB: db}
}

func (m *MySQLGORM) Close() {
	sqlDB, err := m.DB.DB()
	if err != nil {
		log.Printf("[MySQLGORM] Error al obtener la conexión SQL: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("[MySQLGORM] Error al cerrar la conexión: %v", err)
	} else {
		log.Println("[MySQLGORM] Conexión cerrada correctamente")
	}
}
