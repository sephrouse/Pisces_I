package landing

import "github.com/kataras/iris"

//Landing ...
func Landing(ctx *iris.Context) {
	ctx.Render("landing.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
}
