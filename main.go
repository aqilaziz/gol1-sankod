package main

import (
	"go_santri_koding/backend-api/config"
	"go_santri_koding/backend-api/database"
	"go_santri_koding/backend-api/routes"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	//inisialiasai Gin
	router := routes.SetupRouter()

	

	//mulai server dengan port 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
