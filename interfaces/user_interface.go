package interfaces

import (
	"go_mongo/models"
)

type UserInterface interface {
	Insert(user models.User) (models.User, error)
	FindMany(conditions map[string]interface{}) (users []models.User,err error)
	FindOne(conditions map[string]interface{}) (user models.User, err error)
	Update(user models.User, conditions map[string]interface{}) (models.User, error)
	Delete(conditions map[string]interface{}) error
}