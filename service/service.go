package service

import (
	"context"
	"strconv"

	"medicare-backend/database"
	"medicare-backend/graph/model"
	"medicare-backend/internal/auth"
	models "medicare-backend/internal/model"

	"medicare-backend/internal/utils"
)

type AuthService struct {
	userRepo *database.UserRepository
}

// NewAuthService creates a new auth service instance
func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: database.NewUserRepository(),
	}
}

// Register registers a new user
func (s *AuthService) Register(ctx context.Context, input model.RegisterInput) (*model.AuthPayload, error) {
	// Validate inputs
	if err := s.validateRegisterInput(input); err != nil {
		return nil, err
	}

	// Sanitize inputs
	input.Name = utils.SanitizeString(input.Name)
	input.Email = utils.SanitizeString(input.Email)

	// Check if email already exists
	exists, err := s.userRepo.CheckEmailExists(ctx, input.Email)
	if err != nil {
		return nil, utils.NewDatabaseError("failed to check email")
	}
	if exists {
		return nil, utils.NewValidationError("email already registered")
	}

	// Create user model
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
		Role:     "patient",
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		return nil, utils.NewInternalServer("failed to hash password")
	}

	// Save to database
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, utils.NewDatabaseError("failed to create user")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, utils.NewAuthError("failed to generate token")
	}

	// Convert to GraphQL model
	graphqlUser := s.toGraphQLUser(user)

	return &model.AuthPayload{
		Token: token,
		User:  graphqlUser,
	}, nil
}

// Login authenticates user and returns token
func (s *AuthService) Login(ctx context.Context, input model.LoginInput) (*model.AuthPayload, error) {
	// Validate inputs
	if err := s.validateLoginInput(input); err != nil {
		return nil, err
	}

	// Sanitize email
	input.Email = utils.SanitizeString(input.Email)

	// Find user by email
	user, err := s.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if err == database.ErrUserNotFound {
			return nil, utils.NewAuthError("invalid credentials")
		}
		return nil, utils.NewDatabaseError("failed to find user")
	}

	// Verify password
	if !user.CheckPassword(input.Password) {
		return nil, utils.NewAuthError("invalid credentials")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, utils.NewAuthError("failed to generate token")
	}

	// Convert to GraphQL model
	graphqlUser := s.toGraphQLUser(user)

	return &model.AuthPayload{
		Token: token,
		User:  graphqlUser,
	}, nil
}

// GetCurrentUser retrieves current authenticated user
func (s *AuthService) GetCurrentUser(ctx context.Context, userID uint) (*model.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if err == database.ErrUserNotFound {
			return nil, utils.NewNotFound("user not found")
		}
		return nil, utils.NewDatabaseError("failed to get user")
	}

	return s.toGraphQLUser(user), nil
}

// Private helper methods

func (s *AuthService) validateRegisterInput(input model.RegisterInput) error {
	if err := utils.ValidateName(input.Name); err != nil {
		return utils.NewValidationError("name is required")
	}

	if err := utils.ValidateEmail(input.Email); err != nil {
		return utils.NewValidationError("invalid email format")
	}

	if err := utils.ValidatePassword(input.Password); err != nil {
		return utils.NewValidationError("password must be at least 6 characters")
	}

	return nil
}

func (s *AuthService) validateLoginInput(input model.LoginInput) error {
	if err := utils.ValidateEmail(input.Email); err != nil {
		return utils.NewValidationError("invalid email format")
	}

	if input.Password == "" {
		return utils.NewValidationError("password is required")
	}

	return nil
}

func (s *AuthService) toGraphQLUser(user *models.User) *model.User {
	return &model.User{
		ID:        strconv.FormatUint(uint64(user.ID), 10),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     &user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
