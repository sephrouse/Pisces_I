package main

import (
	"pisces_i/src/landing"
	"pisces_i/src/login"
	"pisces_i/src/platform"

	"github.com/iris-contrib/template/html"
	"github.com/kataras/iris"
)

func main() {

	// init web server
	iris.Config.IsDevelopment = true

	//iris.StaticServe("./public/assets", "/public")
	iris.Static("/public", "./public/assets", 1)
	iris.Static("/p/public", "./public/assets", 2)
	iris.Config.Gzip = true
	iris.Config.Charset = "UTF-8"
	iris.UseTemplate(html.New(html.Config{Layout: "dashboard_layout.html"})).Directory("./src/templates", ".html")

	// show 404 page
	// iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
	// 	ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	// })

	// define web routes
	iris.Get("/", landing.Landing)

	iris.Get("/login", login.Login)
	iris.Post("/login", login.PostLogin)

	pf := iris.Party("/p")
	{
		// should use a middleware func to check session. whether the user is valid or not.
		//platform.UseFunc(func(c *iris.Context) { // check user session ...})
		//read the book to implement the validation.
		pf.UseFunc(func(c *iris.Context) {
			c.Next()
		})

		//pf.Static("/p", "./public/assets", 1)

		pf.Get("/", platform.Index)
		pf.Get("/prepareForEx", platform.PrepareForEx)
		pf.Get("/record", platform.Record)
		pf.Get("/rank", platform.Rank)
		pf.Get("/vocabularyTest", platform.VocabularyTest)
		pf.Get("/profile", platform.Profile)
	}

	// iris.Get("/p", platform.Index)
	// iris.Get("/p/prepareForEx", platform.PrepareForEx)
	// iris.Get("/p/record", platform.Record)
	// iris.Get("/p/rank", platform.Rank)
	// iris.Get("/p/vocabularyTest", platform.VocabularyTest)

	iris.Listen(":8080")
}
