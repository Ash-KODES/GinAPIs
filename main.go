package main

import (
	"socialapp/utils"
	"socialapp/routes"
)

func main() {
	utils.ConnectDB()
	router := routes.SetupRouter()
	router.Run(":8080")
}
