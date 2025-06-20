package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model will provide the ID, CreatedAt, UpdatedAt fields
	FirstName  string ` json:"first_name"`
	LastName   string ` json:"last_name"`
	Email      string `gorm:"unique; notNull" json:"email"`
	Password   string ` json:"password"`
	Role       string `gorm:"default:user" json:"role"`
	Age        int    ` json:"age"`
}
