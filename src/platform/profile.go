package platform

import "github.com/kataras/iris"

//Profile ...
func Profile(ctx *iris.Context) {
	ctx.Render("profile.html", nil)
}
