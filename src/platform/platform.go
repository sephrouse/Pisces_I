package platform

import "github.com/kataras/iris"

func Index(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}
