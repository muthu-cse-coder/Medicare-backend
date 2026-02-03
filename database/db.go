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

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	// Local development-க்கு .env load செய்யலாம்
	_ = godotenv.Load()

	// Railway DATABASE_URL முதலில் பார்க்கும்
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		// Local testing fallback
		host := os.Getenv("PGHOST")
		port := os.Getenv("PGPORT")
		user := os.Getenv("PGUSER")
		password := os.Getenv("PGPASSWORD")
		dbname := os.Getenv("PGDATABASE")

		if host == "" || port == "" || user == "" || password == "" || dbname == "" {
			return fmt.Errorf("database connection info not set in environment")
		}

		connectionString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
	}

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Database connected successfully!")
	return nil
}
