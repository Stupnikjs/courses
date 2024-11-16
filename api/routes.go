package api

import (
	"net/http"

	"github.com/Stupnikjs/courses/database"
	"github.com/go-chi/chi/v5"
)

type Application struct {
	Port int
	DB   *database.PostgresRepo
}

func (app *Application) Routes() http.Handler {

	mux := chi.NewRouter()

	// register routes
	mux.Get("/", app.RenderAccueil)
	mux.Post("/post", app.SelectArticlePost)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
