package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const (
	ContentFile = "./content.json"
	TplFile     = "./tpl.html"
	WebDir      = "docs/"
)

type Page struct {
	Title    string
	When     string
	Where    string
	Slug     string
	Video    string
	Body     string
	Filename string
	Next     string
	Prev     string
}

func getContent() []Page {
	jsonFile, err := ioutil.ReadFile(ContentFile)
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

func createPage(t *template.Template, idx int, pages []Page, filename string) {
	f, err := os.Create(WebDir + filename)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	if err := t.Execute(f, pages[idx]); err != nil {
		log.Fatal("error with template: ", err)
		return
	}

	_ = f.Close()
}

func main() {
	pages := getContent()

	tplFile, err := ioutil.ReadFile(TplFile)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	t, err := template.New("p").Parse(string(tplFile))
	if err != nil {
		fmt.Println(err)
	}

	createPage(t, 0, pages, "index.html")

	for idx := range pages {
		createPage(t, idx, pages, pages[idx].Slug)
	}
}
