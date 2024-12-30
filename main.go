package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	// "io"
	"log"
	"os"

	// "github.com/russross/blackfriday"
	"github.com/yuin/goldmark"

	"github.com/joho/godotenv"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generateHTML(content string) string {
	shell := os.Getenv("SHELL")
	var buf bytes.Buffer
	t, _ := template.ParseFiles(shell)
	data := map[string]interface{}{
		"Content": template.HTML(content),
	}
	err := t.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func generateBlogs() string {
	jasonList := os.Getenv("JSON_LIST")
	blogsShell := os.Getenv("BOLOG_SHELL")
	var buf bytes.Buffer
	t, _ := template.ParseFiles(blogsShell)

	jsonData, _ := os.ReadFile(jasonList)
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	templateData := map[string]interface{}{
		"Content": data["content"],
	}
	t.Execute(&buf, templateData)

	return buf.String()
}



func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.Table, extension.Strikethrough, extension.Linkify),
		goldmark.WithRendererOptions(html.WithUnsafe()), // Allow raw HTML
	)

	jasonList := os.Getenv("JSON_LIST")
	jsonData, _ := os.ReadFile(jasonList)
	var data map[string][]map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}
	mdFolder := os.Getenv("MD_FOLDER")
	htmlFolder := os.Getenv("HTML_FOLDER")

	for _, value := range data["content"] {

		var buf bytes.Buffer
		sourcPath := fmt.Sprint(mdFolder, value["sourc"])
		source, _ := os.ReadFile(sourcPath)
		err = md.Convert(source, &buf)
		if err != nil {
			log.Fatal(err)
		}
		page := generateHTML(buf.String())
		pathToSave := fmt.Sprint(htmlFolder, value["link"])
		err = os.WriteFile(pathToSave, []byte(page), 0644)
		check(err)

	}
	// ganarting the blogs
	homePgae := generateBlogs()
	blogsPage := fmt.Sprint(htmlFolder, "blogs.html")
	err = os.WriteFile(blogsPage, []byte(homePgae), 0644)
	check(err)
	//  making the home page in here
	HOMESHELL := os.Getenv("HOME_SHELL")
	homeShell, _ := os.ReadFile(HOMESHELL)
	indexPage := fmt.Sprint(htmlFolder, "index.html")
	err = os.WriteFile(indexPage, homeShell, 0644)
	check(err)
	//  making the about shell
	ABOUTSHELL := os.Getenv("ABOUT_SHELL")
	var buff bytes.Buffer
		source, _ := os.ReadFile(ABOUTSHELL)
		err = md.Convert(source, &buff)
		if err != nil {
			log.Fatal(err)
		}
	page := generateHTML(buff.String())

	err = os.WriteFile(strings.Replace(ABOUTSHELL,".md",".html",1), []byte(page), 0644)
	check(err)

	fmt.Printf("all done the source files are in the %v", htmlFolder)
}
