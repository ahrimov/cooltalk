package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ahrimov/cooltalk-backend/internal/api"
	"github.com/ahrimov/cooltalk-backend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	mainDB := database.OpenDatabase()
	defer mainDB.CloseDatabase()
	router := api.SetUpRouter(mainDB)

	serverUrl := fmt.Sprint(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	router.Run(serverUrl)

}
