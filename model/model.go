package model
import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model 
	Name string
	LastName string
	Age int32
}