package main

import (
	"golang-crud/config"
	"golang-crud/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Start(":8080")
}
