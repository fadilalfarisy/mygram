package main

import (
	"challenge-4/databases"
	"challenge-4/router"
)

func main() {
	databases.StartDB()
	r := router.StartApp()
	r.Run(":3000")
}
