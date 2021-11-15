package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khrees2412/linkpreview/routes"
	
	"os"
)
 



func main() {
  app := fiber.New()
  
  routes.Setup(app)
	
	port := os.Getenv("PORT")
	if port == ""{
		port = ":8080"
	}
  app.Listen(":"+port)
}
