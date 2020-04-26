package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	strip "github.com/grokify/html-strip-tags-go"
)

const (
	ContentFile = "./content.json"
	TplFile     = "./tpl.poem.html"
	ArchiveFile = "./tpl.archive.html"
	WebDir      = "docs/"
)

type Byline []string

type Page struct {
	Title     string
	When      string
	Audio     string
	Where     string
	Slug      string
	Video     string
	Body      string
	BodyText  string
	Filename  string
	Next      string
	Prev      string
	NextTitle string
	PrevTitle string
	Who       []Byline
}

type AllPages struct {
	Pages []Page
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

func createPage(t *template.Template, page interface{}, filename string) {
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

func createTpl(file string) *template.Template {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("p").Parse(string(contents))
	if err != nil {
		fmt.Println(err)
	}

	return t
}

func main() {
	pages := getContent()
	poemTpl := createTpl(TplFile)
	archiveTpl := createTpl(ArchiveFile)

	if len(pages) > 0 {
		pages[0].Prev = pages[1].Slug
		pages[0].PrevTitle = pages[1].Title
	}

	pages[0].BodyText = strip.StripTags(pages[0].Body)

	createPage(archiveTpl, &pages, "archive.html")
	createPage(poemTpl, &pages[0], "index.html")

	for idx := range pages {
		p := &pages[idx]

		if idx > 0 {
			p.Next = pages[idx-1].Slug
			p.NextTitle = pages[idx-1].Title
		}

		if idx != len(pages)-1 {
			p.Prev = pages[idx+1].Slug
			p.PrevTitle = pages[idx+1].Title
		}

		p.BodyText = strip.StripTags(p.Body)

		createPage(poemTpl, p, p.Slug)
	}
}
