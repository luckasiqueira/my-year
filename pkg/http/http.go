package http

import (
	"log"
	"my-year/pkg/database"
	"my-year/ui"

	"github.com/kataras/iris/v12"
)

func Start() {
	srv := iris.Default()
	routes(srv)
	srv.Listen(":8008")
}

func routes(ctx *iris.Application) {
	ctx.Get("/", func(ctx iris.Context) {
		list, err := database.List()
		if err != nil {
			log.Println(err)
		}
		ctx.RenderComponent(ui.Home(list))
	})
	ctx.Patch("/", func(ctx iris.Context) {
		item := ctx.URLParam("action")
		database.Sum(item)
		list, err := database.List()
		if err != nil {
			log.Println(err)
		}
		ctx.RenderComponent(ui.Activities(list))
	})
}
