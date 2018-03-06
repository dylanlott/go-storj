package main

import (
	"github.com/kataras/iris"
	"fmt"
)

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
}

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		user := make(map[string]string)

		user["username"] = "admin" 

		fmt.Println(user)
		// respond with json of map m
		ctx.JSON(user)
	})

	app.Get("/users/{id:long min(1)}", func(ctx iris.Context) {
		var user User

		id, _ := ctx.Params().GetInt64("id")

		user.Id = id
		user.Username ="admin"

		ctx.JSON(user)
	})

	app.Run(iris.Addr(":8080"))
}
