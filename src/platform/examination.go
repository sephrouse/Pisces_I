package platform

import "github.com/kataras/iris"

//PrepareForEx ...
func PrepareForEx(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}
