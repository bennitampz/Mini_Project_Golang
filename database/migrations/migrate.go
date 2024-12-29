package migrations

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"tokoku_go/config"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.DB

	migrationDir := "database/migrations"

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("Error reading migrations directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".up.sql") {
			continue
		}

		filePath := filepath.Join(migrationDir, file.Name())
		sqlBytes, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading migration file %s: %v", file.Name(), err)
		}

		sql := string(sqlBytes)
		if err := executeMigration(db, sql); err != nil {
			log.Fatalf("Error executing migration %s: %v", file.Name(), err)
		}

		log.Printf("Migration %s executed successfully", file.Name())
	}
}

func executeMigration(db *gorm.DB, sql string) error {
	result := db.Exec(sql)
	if result.Error != nil {
		return fmt.Errorf("error executing SQL: %w", result.Error)
	}
	return nil
}
