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
	Title     string
	When      string
	Where     string
	Slug      string
	Video     string
	Body      string
	Filename  string
	Next      string
	Prev      string
	NextTitle string
	PrevTitle string
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

func createPage(t *template.Template, page *Page, filename string) {
	f, err := os.Create(WebDir + filename)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	if err := t.Execute(f, page); err != nil {
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

	if len(pages) > 0 {
		pages[0].Prev = pages[1].Slug
		pages[0].PrevTitle = pages[1].Title
	}

	createPage(t, &pages[0], "index.html")

	for idx := range pages {
		p := &pages[idx]

		if idx > 0 {
			// if val, ok := pages[idx - 1k
			p.Next = pages[idx-1].Slug
			p.NextTitle = pages[idx-1].Title
		}

		if idx != len(pages)-1 {
			p.Prev = pages[idx+1].Slug
			p.PrevTitle = pages[idx+1].Title
		}

		createPage(t, p, p.Slug)
	}
}
