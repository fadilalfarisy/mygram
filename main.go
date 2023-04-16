package main

import (
	"challenge-4/databases"
	_ "challenge-4/docs"
	"challenge-4/router"
)

// @title MyGram application
// @version 1.0
// @description This is a simple servie for managing social media
// @termOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

func main() {
	databases.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
