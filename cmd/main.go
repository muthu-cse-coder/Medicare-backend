package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"medicare-backend/database"
// 	"medicare-backend/graph"
// 	"medicare-backend/internal/auth"
// 	"medicare-backend/internal/auth/middleware"
// 	models "medicare-backend/internal/model"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/joho/godotenv"
// )

// const defaultPort = "8080"

// func main() {
// 	// Load environment variables
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Warning: .env file not found")
// 	}

// 	// Initialize JWT
// 	if err := auth.InitJWT(); err != nil {
// 		log.Fatal("Failed to initialize JWT:", err)
// 	}

// 	// Connect to database
// 	if err := database.Connect(); err != nil {
// 		log.Fatal("Database connection failed:", err)
// 	}
// 	defer database.Close()

// 	// Run migrations
// 	if err := database.Migrate(&models.User{}); err != nil { // ← Should use internal/models
// 		log.Fatal("Migration failed:", err)
// 	}

// 	// Get port from environment
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	// Create GraphQL resolver
// 	resolver := graph.NewResolver()

// 	// Create GraphQL server
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
// 		Resolvers: resolver,
// 	}))

// 	// Setup routes
// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

// 	// GraphQL endpoint with auth middleware and CORS
// 	http.Handle("/query",
// 		middleware.CORS(
// 			auth.Middleware(srv),
// 		),
// 	)

// 	// Health check endpoint
// 	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	})

// 	// Start server
// 	log.Printf("🚀 connect to http://localhost:%s/ for GraphQL playground", port)

// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal("Server failed to start:", err)
// 	}
// }
