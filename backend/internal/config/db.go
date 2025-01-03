package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql" // エイリアスを設定
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormMysql "gorm.io/driver/mysql" // GORMのMySQLドライバーもエイリアス化
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// DSN (Data Source Name) を構築
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// GORMでDB接続
	database, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{}) // エイリアスを使用
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// 標準SQL DB接続
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to get generic database object:", err)
	}

	// マイグレーションの実行
	runMigrations(sqlDB, dbName)

	DB = database
	log.Println("Database connection established.")
}

func runMigrations(db *sql.DB, dbName string) {
	driver, err := migrateMysql.WithInstance(db, &migrateMysql.Config{}) // エイリアスを使用
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations", // マイグレーションファイルのパス
		dbName,
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully.")
}
