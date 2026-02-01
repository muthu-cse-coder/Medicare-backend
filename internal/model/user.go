// package models

// import (
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID        uint      `gorm:"primarykey;autoIncrement" json:"id"`
// 	Name      string    `gorm:"size:255;not null" json:"name"`
// 	Email     string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
// 	Password  string    `gorm:"size:255;not null" json:"-"`
// 	Phone     string    `gorm:"size:20" json:"phone"`
// 	Role      string    `gorm:"size:20;default:patient" json:"role"`
// 	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
// }

// // BeforeCreate hook - hash password before saving
// func (u *User) BeforeCreate(tx *gorm.DB) error {
// 	if u.Password != "" {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			return err
// 		}
// 		u.Password = string(hashedPassword)
// 	}
// 	return nil
// }

// // CheckPassword verifies if provided password matches stored hash
// func (u *User) CheckPassword(password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// 	return err == nil
// }

// // SetPassword hashes and sets new password
//
//	func (u *User) SetPassword(password string) error {
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//		if err != nil {
//			return err
//		}
//		u.Password = string(hashedPassword)
//		return nil
//	}
package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never expose password in JSON
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HashPassword hashes the user password
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies if provided password matches stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// Sanitize removes sensitive data
func (u *User) Sanitize() {
	u.Password = ""
}
