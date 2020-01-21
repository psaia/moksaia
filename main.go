package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title string
	When  string
	Slug  string
	Video string
	Body  string
}

func getContent() []Page {
	jsonFile, err := ioutil.ReadFile("./content.json")
	if err != nil {
		fmt.Println(err)
	}

	var all []Page

	err = json.Unmarshal(jsonFile, &all)
	if err != nil {
		fmt.Println(err)
	}

	return all
}

func main() {
	tplFile, err := ioutil.ReadFile("./tpl.html")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	pages := getContent()

	for idx, page := range pages {
		t, err := template.New("p").Parse(string(tplFile))
		if err != nil {
			fmt.Println(err)
		}

		var filename string

		if idx == 0 {
			filename = "index.html"
		} else {
			filename = page.Slug + ".html"
		}

		f, err := os.Create("docs/" + filename)
		if err != nil {
			log.Println("create file: ", err)
			return
		}

		if err = t.Execute(f, page); err != nil {
			log.Fatal("error with template: ", err)
			return
		}

		_ = f.Close()
	}
}
