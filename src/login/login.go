package login

import "github.com/kataras/iris"

//UserFormData ...
type UserFormData struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

/*Login ...
Show Login page.
*/
func Login(ctx *iris.Context) {
	ctx.Render("login.html", nil)
}

//PostLogin ...
func PostLogin(ctx *iris.Context) {
	user := UserFormData{}
	ctx.ReadForm(&user)
	ctx.Write("in function of PostLogin. %s, %s", user.Email, user.Password)
}
