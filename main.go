package main

import (
	"github.com/kataras/iris"
	"fmt"
)

// define structs for types 
type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
}

type Contact struct {
	Id int64 `json:"id"`	
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

	app.Get("/contacts", func(ctx iris.Context) {
		users := make([]string, 0)
		users = append(users, "user1")
		ctx.JSON(users)
	})

	app.Get("/contacts/{id:long min(1)}", func(ctx iris.Context) {

		details := make(map[string]string)
		details["id"] = ctx.Params().Get("id")
		details["username"] = "admin"

		ctx.JSON(details)
	})

	app.Run(iris.Addr(":8080"))
}
