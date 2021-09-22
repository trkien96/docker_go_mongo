package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_mongo/constant"
	"go_mongo/models"
	"log"
)

type User struct {
	Db *mongo.Database
}

var UserCollection = "users"

func (mongo User) Insert(user models.User) (models.User, error) {
	userData, _ := bson.Marshal(user)
	_, err := mongo.Db.Collection(UserCollection).InsertOne(context.Background(), userData)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo User) FindMany(conditions map[string]interface{}) (users []models.User, err error) {
	//find records
	//pass these options to the Find method
	//findOptions := options.Find()
	//Set the limit of the number of record to find
	//findOptions.SetLimit(5)
	//result, _ := mongo.Db.Collection(UserCollection).Find(context.Background(), conditions, findOptions)
	result, _ := mongo.Db.Collection(UserCollection).Find(context.Background(), conditions)
	for result.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var item models.User

		err = result.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		//item.ID, _ = primitive.ObjectIDFromHex(item.ID.Hex())
		users = append(users, item)
	}

	if err = result.Err(); err != nil {
		log.Fatal(err)
	}

	result.Close(context.TODO())
	if len(users) == 0 {
		return users, err
	}

	return users, nil
}

func (mongo User) FindOne(conditions map[string]interface{}) (user models.User, err error) {
	result := mongo.Db.Collection(UserCollection).FindOne(context.Background(), conditions)
	err = result.Decode(&user)
	if user == (models.User{}) {
		return user, constant.ERR_USER_NOT_FOUND
	}

	if err != nil {
		return user, err
	}
	return user, nil
}

func (mongo User) Update(user models.User, conditions map[string]interface{}) (models.User, error) {
	userData, _ := bson.Marshal(user)
	_, err := mongo.Db.Collection(UserCollection).UpdateOne(context.Background(), conditions, userData)
	if err != nil {
		log.Println(err)
		return user, err
	}
	return user, nil
}

func (mongo User) Delete(conditions map[string]interface{}) error {
	_, err := mongo.Db.Collection(UserCollection).DeleteOne(context.Background(), conditions)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
