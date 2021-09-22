package services

import (
	"go_mongo/driver"
	"go_mongo/interfaces"
	"go_mongo/repository"
)

func Factory() interfaces.UserInterface {
	return repository.User {
		Db: driver.Mongo.Db,
	}
}