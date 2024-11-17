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
	articles, err := loadArticles()
	if err != nil {
		fmt.Println(err)
	}

	_, err = app.DB.DB.Exec("DROP TABLE article ; ")

	if err != nil {
		fmt.Println(err)
	}
	err = app.DB.InitTables(articles)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}

func loadArticles() ([]database.Article, error) {
	var courses []string
	var articles []database.Article
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(cwd, "courses.json"))
	if err != nil {
		return nil, err
	}
	byteCourses, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteCourses, &courses)
	if err != nil {
		return nil, err
	}
	for _, c := range courses {
		var article database.Article
		article.Name = c
		articles = append(articles, article)
	}
	return articles, nil
}
