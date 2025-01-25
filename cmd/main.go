package main

import (
	"cloud-martini-backend/db"
	"cloud-martini-backend/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Started HAckthon")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	db.ConnectMongo(MONGO_URI)
	defer db.DisconnectMongo()

	r := router.SetupRouter()
	APP_PORT := os.Getenv("APP_PORT")
	r.Run(":" + APP_PORT)
}
