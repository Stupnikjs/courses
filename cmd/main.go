package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Stupnikjs/courses/api"
	"github.com/Stupnikjs/courses/database"
	_ "google.golang.org/api/option"
)

func main() {
	/*
		if err := godotenv.Load("./.env"); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	*/
	app := api.Application{
		Port: 8080,
	}

	conn, err := api.ConnectToDB()
	if err != nil {
		fmt.Println(err)
	}

	app.DB = &database.PostgresRepo{
		DB: conn,
	}
	articles := loadArticles()
	fmt.Println(articles)
	err = app.DB.InitTables(articles)
	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}

func loadArticles() []database.Article {
	var courses []string
	var articles []database.Article
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Open(filepath.Join(cwd, "courses.json"))
	if err != nil {
		fmt.Println(err)
	}
	byteCourses, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(byteCourses, &courses)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range courses {
		var article database.Article
		article.Name = c
		articles = append(articles, article)
	}
	return articles
}
