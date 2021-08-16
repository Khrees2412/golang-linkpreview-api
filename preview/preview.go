package previewer

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)


type Res struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	Image string `json:"image"`
}

func Preview(previewUrl string){

	var  title, url, desc, image string;

	collector := colly.NewCollector()

 	collector.OnError(func (_ *colly.Response, err error){
    	log.Fatal("Error: ", err.Error())
  	})

	collector.OnHTML("meta",func (element *colly.HTMLElement)  {

	   name := element.Attr("property")
       content := element.Attr("content")
       contentString := string(content)

       if name == "og:title" {
       title = contentString
	 	}

       if name == "og:url" {
       url = contentString
        }

       if name == "og:image" {
       image = contentString
        }

       if name == "og:description" {
       desc = contentString
        }
	  
	})

	collector.Visit(previewUrl)

	go writeToFile(title,desc,url,image)

}

func writeToFile (title,desc,url,image string) {

	result := Res{
		Title: title,
		Description: desc,
		Url: url,
		Image: image,
	}

	file, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		log.Fatal("Unable to create json file: ", err)
	}
	_ = ioutil.WriteFile("preview.json", file, 0644)
	
 }