package infraestructure

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL() *MySQL {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("[MySQL] Error cargando el archivo .env: %v", err)
	}

	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatalf("[MySQL] La variable MYSQL_DSN no está definida en el archivo .env")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[MySQL] Error al conectar con la base de datos: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("[MySQL] Error al hacer ping a la base de datos: %v", err)
	}

	log.Println("[MySQL] Conexión establecida correctamente")
	return &MySQL{DB: db}
}

func (m *MySQL) Close() {
	if err := m.DB.Close(); err != nil {
		log.Printf("[MySQL] Error al cerrar la conexión: %v", err)
	} else {
		log.Println("[MySQL] Conexión cerrada correctamente")
	}
}
