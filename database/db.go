// package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// var DB *sql.DB

// // Connect establishes PostgreSQL connection
// func Connect(connectionString string) error {
// 	var err error

// 	DB, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		return fmt.Errorf("failed to open database: %w", err)
// 	}

// 	// Test connection
// 	if err = DB.Ping(); err != nil {
// 		return fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	// Set connection pool settings
// 	DB.SetMaxOpenConns(25)
// 	DB.SetMaxIdleConns(5)

// 	log.Println("✅ Database connected successfully!")
// 	return nil
// }

// // Close closes database connection
// func Close() error {
// 	if DB != nil {
// 		return DB.Close()
// 	}
// 	return nil
// }

// // InitSchema creates database tables
// func InitSchema() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(255) NOT NULL,
// 		email VARCHAR(255) UNIQUE NOT NULL,
// 		password VARCHAR(255) NOT NULL,
// 		phone VARCHAR(20),
// 		role VARCHAR(20) DEFAULT 'patient',
// 		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
// 		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);

// 	CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
// 	CREATE INDEX IF NOT EXISTS idx_users_role ON users(role);
// 	`

// 	_, err := DB.Exec(query)
// 	if err != nil {
// 		return fmt.Errorf("failed to create schema: %w", err)
// 	}

//		log.Println("✅ Database schema initialized!")
//		return nil
//	}
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	var err error

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	if host == "" {
		return fmt.Errorf("PGHOST not set (Railway env issue)")
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("✅ Database connected successfully!")
	return nil
}
