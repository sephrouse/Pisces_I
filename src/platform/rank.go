package platform

import "github.com/kataras/iris"

//Rank ...
func Rank(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}
