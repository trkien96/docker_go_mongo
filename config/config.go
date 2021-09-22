package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST       string `json:"db_host"`
	DB_PORT       string `json:"db_port"`
	DB_DATABASE   string `json:"db_database"`
	DB_USERNAME   string `json:"db_username"`
	DB_PASSWORD   string `json:"db_password"`
	MONGO_PORT    string `json:"mongo_port"`
	SERVER_HOST   string `json:"server_host"`
	SERVER_PORT   string `json:"server_port"`
}

var Global = Config{}

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Global.DB_HOST 		 = os.Getenv("DB_HOST")
	Global.DB_PORT 		 = os.Getenv("DB_PORT")
	Global.DB_DATABASE 	 = os.Getenv("DB_DATABASE")
	Global.DB_USERNAME 	 = os.Getenv("DB_USERNAME")
	Global.DB_PASSWORD 	 = os.Getenv("DB_PASSWORD")
	Global.MONGO_PORT 	 = os.Getenv("MONGO_PORT")
	Global.SERVER_HOST 	 = os.Getenv("SERVER_HOST")
	Global.SERVER_PORT 	 = os.Getenv("SERVER_PORT")
}
