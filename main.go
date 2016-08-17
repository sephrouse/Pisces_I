package main

import (
	"pisces_i/src/login"

	"github.com/iris-contrib/template/html"
	"github.com/kataras/iris"
)

func main() {

	// init web server
	iris.Static("/public", "./public/assets", 1)
	iris.Config.Gzip = true
	iris.UseTemplate(html.New()).Directory("./src/templates", ".html")

	// define web routes
	iris.Get("/login", login.Login)
	iris.Post("/login", login.PostLogin)

	iris.Listen(":8080")
}
