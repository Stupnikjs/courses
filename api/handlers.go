package api

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"

	"github.com/Stupnikjs/courses/database"
)

var pathToTemplates = "./static/templates/"

type TemplateData struct {
	Data map[string]any
}

func render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) error {
	_ = r.Method

	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, t), path.Join(pathToTemplates, "/base.layout.gohtml"))
	if err != nil {
		return err
	}
	err = parsedTemplate.Execute(w, td)
	if err != nil {
		return err
	}
	return nil

}

// template rendering

func (app *Application) RenderAccueil(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}

	td.Data = make(map[string]any)
	// get articles from db
	articles, _ := app.DB.GetAllArticles()

	td.Data["articles"] = articles
	_ = render(w, r, "/main.gohtml", &td)
}
func (app *Application) RenderAddArticle(w http.ResponseWriter, r *http.Request) {

	td := TemplateData{}

	td.Data = make(map[string]any)
	// get articles from db
	articles, _ := app.DB.GetAllArticles()

	td.Data["articles"] = articles
	_ = render(w, r, "/addarticle.gohtml", &td)
}

// post

func (app *Application) SelectArticlePost(w http.ResponseWriter, r *http.Request) {
	var SelectedArticles []database.Article
	body := r.Body
	bytesBody, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		fmt.Println(err)
	}
	splited := bytes.Split(bytesBody, []byte("&"))

	for _, b := range splited {
		var article database.Article
		value := bytes.Split(b, []byte("="))[0]
		if bytes.Contains(value, []byte("%20")) {
			spaceB := bytes.ReplaceAll(value, []byte("%20"), []byte(" "))
			article.Name = string(spaceB)
			SelectedArticles = append(SelectedArticles, article)
		} else {
			article.Name = string(value)
			SelectedArticles = append(SelectedArticles, article)
		}

	}
	fmt.Println(SelectedArticles)
	err = app.DB.PushSelectedArticles(SelectedArticles)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

func (app *Application) PostAddArticle(w http.ResponseWriter, r *http.Request) {

	body := r.Body
	bytesBody, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		fmt.Println(err)
	}
	articleName, found := bytes.CutPrefix(bytesBody, []byte("article="))
	if !found {
		w.Write([]byte(err.Error()))
	}
	articleName = bytes.Replace(articleName, []byte("%20"), []byte(" "), 10)
	err = app.DB.InsertOneArticle(string(articleName))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(string(bytesBody)))

}
