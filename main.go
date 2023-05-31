package main

import (
	"github.com/gofiber/swagger"
	"log"

	"github.com/Fedhira/Tugas_1214028/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/Fedhira/Tugas_1214028/url"

	"github.com/gofiber/fiber/v2"

	"github.com/Fedhira/Tugas_1214028/docs"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}