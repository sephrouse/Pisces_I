package platform

import "github.com/kataras/iris"

//VocabularyTest ...
func VocabularyTest(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}
