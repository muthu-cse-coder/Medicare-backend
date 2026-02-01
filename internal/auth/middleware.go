// package auth

// import (
// 	"context"
// 	"net/http"
// 	"strings"
// )

// // Middleware authenticates requests using JWT
// func Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Get Authorization header
// 		authHeader := r.Header.Get("Authorization")

// 		// If no auth header, continue without user context
// 		if authHeader == "" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// Check if it's a Bearer token
// 		tokenParts := strings.Split(authHeader, " ")
// 		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := tokenParts[1]

// 		// Validate token
// 		claims, err := ValidateToken(tokenString)
// 		if err != nil {
// 			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Add user info to context
// 		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
// 		ctx = context.WithValue(ctx, "email", claims.Email)
// 		ctx = context.WithValue(ctx, "role", claims.Role)

// 		// Continue with authenticated context
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// // RequireAuth checks if user is authenticated
// func RequireAuth(ctx context.Context) (uint, error) {
// 	userID, ok := ctx.Value("userID").(uint)
// 	if !ok {
// 		return 0, ErrUnauthorized
// 	}
// 	return userID, nil
// }

// // RequireRole checks if user has specific role
// func RequireRole(ctx context.Context, requiredRole string) error {
// 	role, ok := ctx.Value("role").(string)
// 	if !ok {
// 		return ErrUnauthorized
// 	}

// 	if role != requiredRole {
// 		return ErrForbidden
// 	}

// 	return nil
// }

// var (
// 	ErrUnauthorized = NewError("unauthorized", "User is not authenticated")
// 	ErrForbidden    = NewError("forbidden", "User does not have required permissions")
// )

// // Error represents a custom error
// type Error struct {
// 	Code    string
// 	Message string
// }

// func (e *Error) Error() string {
// 	return e.Message
// }

//	func NewError(code, message string) *Error {
//		return &Error{
//			Code:    code,
//			Message: message,
//		}
//	}
// package auth

// import (
// 	"context"
// 	"net/http"
// 	"strings"
// )

// // Middleware authenticates requests using JWT
// func Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")

// 		// If no auth header, continue without user context
// 		if authHeader == "" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// Check Bearer token format
// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
// 			return
// 		}

// 		// Validate token
// 		claims, err := ValidateToken(parts[1])
// 		if err != nil {
// 			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Add user info to context
// 		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
// 		ctx = context.WithValue(ctx, "email", claims.Email)
// 		ctx = context.WithValue(ctx, "role", claims.Role)

//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
package auth

import (
	"context"
	"net/http"
	"strings"
)

// Middleware authenticates requests using JWT
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// If no auth header, continue without user context
		// Also allow empty Bearer string
		if authHeader == "" || authHeader == "Bearer " {
			next.ServeHTTP(w, r)
			return
		}

		// Check Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		// Validate token
		claims, err := ValidateToken(parts[1])
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)
		ctx = context.WithValue(ctx, "role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
