package domain

import "time"

const (
	BUYER  = "buyer"
	SELLER = "seller"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index; unique; not null"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	Code      int       `json:"code"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
