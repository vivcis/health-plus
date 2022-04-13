package main

import (
	"github.com/decadev/squad10/healthplus/db"
	"github.com/decadev/squad10/healthplus/router"
)

func main() {
	db.SetupDB()

	router.SetupRouter()
}
