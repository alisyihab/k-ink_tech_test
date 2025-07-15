package main

import (
	"registration-smart-reward/internal/config"
	"registration-smart-reward/internal/router"
)

func main() {
	config.InitMongo()

	r := router.SetupRouter()

	r.Run(":8000")
}