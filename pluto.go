package pluto

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"time"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	authConfig := basicauth.Config{
		Users:   map[string]string{"myusername": "mypassword", "name": "password"},
		Realm:   "Auth Required",
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	app.Use(recover.New())
	app.Use(logger.New())

	needAuth := app.Party("/admin", authentication)

	needAuth.Get("/", h)
	needAuth.Get("/profile", h)
	needAuth.Get("/settings", h)

	router(app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func newApp() *iris.Application {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Status": "OK"})
	})

	return app
}

func router(app *iris.Application) {
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"Message": "Hello iris!"})
	})
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}
