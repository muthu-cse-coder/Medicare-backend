package database

import (
	"context"
	"database/sql"
	"errors"
	models "medicare-backend/internal/model"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrDatabaseError      = errors.New("database error")
)

type UserRepository struct{}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser inserts a new user into database
func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password, phone, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	err := DB.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Phone,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail retrieves user by email
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, name, email, password, phone, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	err := DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, name, email, password, phone, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	err := DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// CheckEmailExists checks if email already exists
func (r *UserRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	err := DB.QueryRowContext(ctx, query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UpdateUser updates user information
func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, phone = $3, role = $4, updated_at = NOW()
		WHERE id = $5
	`

	result, err := DB.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Phone,
		user.Role,
		user.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

// DeleteUser deletes user by ID
func (r *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

// GetAllUsers retrieves all users with pagination
func (r *UserRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]*models.User, error) {
	query := `
		SELECT id, name, email, phone, role, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// CountUsers counts total users
func (r *UserRepository) CountUsers(ctx context.Context) (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM users`

	err := DB.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
