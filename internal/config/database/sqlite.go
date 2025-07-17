package database

import (
	"database/sql"
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"financial-cli/internal/domain/transactions"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func InitDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	dbPath := getDBPath()
	var err error

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao SQLite: %w", err)
	}

	if err := Migrate(db); err != nil {
		return nil, fmt.Errorf("erro ao executar migrations: %w", err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&transactions.Transaction{},
		&category.Category{},
		&register.Register{},
	)
}

func CloseDB() error {
	if sqlDB != nil {
		return sqlDB.Close()
	}
	return nil
}

func getDBPath() string {
	if path := os.Getenv("SQLITE_PATH"); path != "" {
		return path
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Não foi possível determinar o diretório home do usuário: " + err.Error())
	}

	dbPath := filepath.Join(homeDir, "fcli.db")
	return dbPath
}
