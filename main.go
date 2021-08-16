package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/linkpreview/routes"
)
 



func main() {
  app := fiber.New()
  
  routes.Setup(app)

  app.Listen(":3000")
}
