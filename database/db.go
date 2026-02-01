// package database

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var DB *gorm.DB

// // Connect to database
// func Connect() error {
// 	dsn := os.Getenv("DATABASE_URL")
// 	if dsn == "" {
// 		return fmt.Errorf("DATABASE_URL environment variable is not set")
// 	}

// 	var err error

// 	// Configure GORM with proper settings
// 	config := &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Info),
// 		NowFunc: func() time.Time {
// 			return time.Now().UTC()
// 		},
// 		PrepareStmt: true,
// 	}

// 	DB, err = gorm.Open(postgres.Open(dsn), config)

// 	if err != nil {
// 		return fmt.Errorf("failed to connect to database: %w", err)
// 	}

// 	// Set connection pool settings
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		return fmt.Errorf("failed to get database instance: %w", err)
// 	}

// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetMaxOpenConns(100)

// 	fmt.Println("✅ Database connected successfully!")
// 	return nil
// }

// // Migrate database tables
// func Migrate(models ...interface{}) error {
// 	// Drop existing tables for fresh migration
// 	err := DB.Migrator().DropTable(models...)
// 	if err != nil {
// 		log.Println("Warning: Could not drop tables:", err)
// 	}

// 	// Create tables with correct structure
// 	err = DB.AutoMigrate(models...)
// 	if err != nil {
// 		return fmt.Errorf("migration failed: %w", err)
// 	}

// 	fmt.Println("✅ Database migration completed!")
// 	return nil
// }

// // GetDB returns database instance
// func GetDB() *gorm.DB {
// 	return DB
// }

// // Close database connection
//
//	func Close() error {
//		sqlDB, err := DB.DB()
//		if err != nil {
//			return err
//		}
//		return sqlDB.Close()
//	}
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Connect establishes PostgreSQL connection
func Connect(connectionString string) error {
	var err error

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("✅ Database connected successfully!")
	return nil
}

// Close closes database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// InitSchema creates database tables
func InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		phone VARCHAR(20),
		role VARCHAR(20) DEFAULT 'patient',
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
	`

	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	log.Println("✅ Database schema initialized!")
	return nil
}
