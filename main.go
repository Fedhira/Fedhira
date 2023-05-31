package main

import (
	"log"

	"github.com/Fedhira/Tugas_1214028/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/Fedhira/Tugas_1214028/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/Fedhira/Tugas_1214028/docs"
)

// @title TES SWAG
// @version 1.0
// @description This is a sample swagger server.

// @contact.name API Support
// @contact.url https://github.com/Fedhira
// @contact.email sfedhira@gmail.com

// @host fedhirasyaila.heroku.com
// @BasePath /
//@schemes https http

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}