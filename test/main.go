package graph

// import (
// 	"context"
// 	"errors"
// 	"medicare-backend/auth"
// 	"medicare-backend/database"
// 	"medicare-backend/models"
// )

// type Resolver struct{}

// // Register User
// func (r *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthPayload, error) {
//     // Check if user already exists
//     var existingUser models.User
//     result := database.DB.Where("email = ?", input.Email).First(&existingUser)
//     if result.RowsAffected > 0 {
//         return nil, errors.New("user with this email already exists")
//     }

//     // Create new user
//     user := models.User{
//         Name:     input.Name,
//         Email:    input.Email,
//         Password: input.Password, // Will be hashed in BeforeCreate
//         Phone:    *input.Phone,
//         Role:     "patient",
//     }

//     if err := database.DB.Create(&user).Error; err != nil {
//         return nil, err
//     }

//     // Generate JWT token
//     token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
//     if err != nil {
//         return nil, err
//     }

//     return &AuthPayload{
//         Token: token,
//         User:  &user,
//     }, nil
// }

// // Login User
// func (r *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthPayload, error) {
//     var user models.User

//     // Find user by email
//     if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
//         return nil, errors.New("invalid credentials")
//     }

//     // Check password
//     if !user.CheckPassword(input.Password) {
//         return nil, errors.New("invalid credentials")
//     }

//     // Generate JWT token
//     token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
//     if err != nil {
//         return nil, err
//     }

//     return &AuthPayload{
//         Token: token,
//         User:  &user,
//     }, nil
// }

// // Get Current User
// func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
//     // Get user from context (set by auth middleware)
//     userID, ok := ctx.Value("userID").(uint)
//     if !ok {
//         return nil, errors.New("unauthorized")
//     }

//     var user models.User
//     if err := database.DB.First(&user, userID).Error; err != nil {
//         return nil, err
//     }

//     return &user, nil
// }

// func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
// func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }
