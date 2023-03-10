package api

import (
	"fmt"
	"log"
	"os"

	"blog-application/api/controllers"
	"blog-application/api/seed"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load() //load environment variables from a file, will return nil if error not exist
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	//connect to database
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")

}
