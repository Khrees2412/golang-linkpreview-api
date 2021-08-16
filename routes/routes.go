package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	previewer "github.com/khrees2412/linkpreview/preview"
)

type Data struct {
	Link string `json:"link"`
}


type Res struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	Image string `json:"image"`
}

var data Res

func Setup (app *fiber.App){	
	app.Post("/generate-preview", GeneratePreview )
}

func GeneratePreview(c *fiber.Ctx) error {
		d := new(Data)

		if err := c.BodyParser(d); err != nil {
			return err
		}
		previewer.Preview(d.Link)

		SendPreview()

		return c.JSON(fiber.Map{
			"message": "success",
			"data" : data,
		})

}

func SendPreview(){
		file, err := ioutil.ReadFile("preview.json")
		if err != nil {
		log.Fatal("Error while reading from file: ", err)
		}

		jsonerr := json.Unmarshal(file, &data)
		if err != nil {
		log.Fatal("Error: ", jsonerr)
		}
}

