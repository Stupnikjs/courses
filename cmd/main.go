package main

import (
	"fmt"
	"net/http"

	"github.com/Stupnikjs/courses/api"
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
	/*
		in local no db
		conn, err := api.ConnectToDB()
		if err != nil {
			fmt.Println(err)
		}

		app.DB = &database.PostgresRepo{
			DB: conn,
		}

		err = app.DB.InitTables()
		if err != nil {
			fmt.Println(err)
		}
	*/
	http.ListenAndServe(fmt.Sprintf(":%d", app.Port), app.Routes())

}
