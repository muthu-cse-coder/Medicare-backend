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
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Initialize JWT
	if err := auth.Init(cfg.JWT.Secret); err != nil {
		log.Fatal("Failed to initialize JWT:", err)
	}

	// Connect to database using Railway DATABASE_URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = cfg.GetDatabaseURL() // fallback
	}

	if err := database.Connect(dbURL); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Initialize database schema
	if err := database.InitSchema(); err != nil {
		log.Fatal("Failed to initialize schema:", err)
	}

	// Create GraphQL server
	resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	}))

	// Setup routes
	mux := http.NewServeMux()

	// GraphQL Playground only in development
	if cfg.Server.Environment == "development" {
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		log.Println("🎮 GraphQL Playground available at /")
	}

	// GraphQL endpoint with middleware
	graphqlHandler := middleware.Logger(
		middleware.CORS(cfg.Server.FrontendURL)(
			auth.Middleware(srv),
		),
	)
	mux.Handle("/query", graphqlHandler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// PORT from Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}

	log.Printf("🚀 Server starting on port %s", port)
	log.Println("📊 GraphQL endpoint: /query")
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
