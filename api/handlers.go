package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
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
	var courses []string
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}
	file, err := os.Open(filepath.Join(cwd, "courses.json"))
	if err != nil {
		fmt.Println(err)
	}
	byteCourses, err := io.ReadAll(file)
	err = json.Unmarshal(byteCourses, &courses)
	if err != nil {
		fmt.Println(err)
	}
	td := TemplateData{}
	td.Data = make(map[string]any)
	fmt.Println(courses)
	td.Data["articles"] = courses
	_ = render(w, r, "/main.gohtml", &td)
}

// post

func (app *Application) Post(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
	fmt.Println("here")
}
