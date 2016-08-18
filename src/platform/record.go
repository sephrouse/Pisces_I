package platform

import (
	"time"

	"github.com/kataras/iris"
)

//Record ...
func Record(ctx *iris.Context) {
	time.Sleep(3 * time.Second)
	ctx.Render("index.html", nil)
}
