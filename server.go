// package main

// import (
// 	"log"
// 	"net/http"

// 	"medicare-backend/config"
// 	"medicare-backend/database"
// 	"medicare-backend/graph"
// 	"medicare-backend/internal/auth"

// 	"medicare-backend/internal/auth/middleware"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// )

// func main() {
// 	// Load configuration
// 	cfg, err := config.Load()
// 	if err != nil {
// 		log.Fatal("Failed to load config:", err)
// 	}

// 	// Initialize JWT
// 	if err := auth.Init(cfg.JWT.Secret); err != nil {
// 		log.Fatal("Failed to initialize JWT:", err)
// 	}

// 	// Connect to database
// 	if err := database.Connect(cfg.GetDatabaseURL()); err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// 	defer database.Close()

// 	// Initialize database schema
// 	if err := database.InitSchema(); err != nil {
// 		log.Fatal("Failed to initialize schema:", err)
// 	}

// 	// Create GraphQL server
// 	resolver := graph.NewResolver()
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
// 		Resolvers: resolver,
// 	}))

// 	// Setup routes
// 	mux := http.NewServeMux()

// 	// GraphQL Playground (only in development)
// 	if cfg.Server.Environment == "development" {
// 		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 		log.Println("🎮 GraphQL Playground: http://localhost:" + cfg.Server.Port + "/")
// 	}

// 	// GraphQL endpoint with middleware chain
// 	graphqlHandler := middleware.Logger(
// 		middleware.CORS(cfg.Server.FrontendURL)(
// 			auth.Middleware(srv),
// 		),
// 	)
// 	mux.Handle("/query", graphqlHandler)

// 	// Health check endpoint
// 	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"status":"healthy"}`))
// 	})

// 	// Start server
// 	log.Println("🚀 Server starting on http://localhost:" + cfg.Server.Port)
// 	log.Println("📊 GraphQL endpoint: http://localhost:" + cfg.Server.Port + "/query")

//		if err := http.ListenAndServe(":"+cfg.Server.Port, mux); err != nil {
//			log.Fatal("Server failed to start:", err)
//		}
//	}
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"medicare-backend/config"
	"medicare-backend/database"
	"medicare-backend/graph"
	"medicare-backend/internal/auth"
	"medicare-backend/internal/auth/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

func main() {
	// Load local .env (optional, safe if exists)
	_ = godotenv.Load()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize JWT
	if err := auth.Init(cfg.JWT.Secret); err != nil {
		log.Fatal("Failed to initialize JWT:", err)
	}

	// Connect to Database
	if err := database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Initialize schema
	// if err := database.InitSchema(); err != nil {
	// 	log.Fatal("Schema init failed:", err)
	// }

	// GraphQL server
	resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	mux := http.NewServeMux()

	// Playground only in development
	if cfg.Server.Environment == "development" {
		mux.Handle("/", playground.Handler("GraphQL Playground", "/query"))
		log.Println("🎮 GraphQL Playground available at /")
	}

	// GraphQL endpoint
	graphqlHandler := middleware.Logger(
		middleware.CORS(cfg.Server.FrontendURL)(
			auth.Middleware(srv),
		),
	)
	mux.Handle("/query", graphqlHandler)

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// Test DB connection
	mux.HandleFunc("/test-db", func(w http.ResponseWriter, r *http.Request) {
		var count int
		err := database.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("DB Error: %v", err)))
			return
		}
		w.Write([]byte(fmt.Sprintf("Users in DB: %d", count)))
	})

	// PORT handling: Railway override automatically
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
		if port == "" {
			port = "8080" // fallback for local
		}
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Println("📊 GraphQL endpoint: /query")
	log.Println("🔥 About to start HTTP server...")

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
