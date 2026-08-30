package database

import (
	"fmt"
	"log"
	"os"

	"fiber-auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect membuka koneksi ke database PostgreSQL dan menjalankan auto-migration.
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database: ", err)
	}

	// Auto-migrate akan membuat/menyesuaikan tabel sesuai struct model.
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}

	DB = db
	log.Println("Database terkoneksi dan migrasi selesai")
}
