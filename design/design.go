package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("todoz", func() {
	Title("todoz")
	Description("Treasure mid")
	Host("localhost:8080")
	Scheme("http")

	Origin("http://localhost:8080/swagger", func() {
		Expose("X-Time")
		Methods("GET", "POST", "PUT", "DELETE")
		MaxAge(600)
		Credentials()
	})
})

var _ = Resource("todos", func() {
	BasePath("/todos")
	DefaultMedia(TodoMediaType)

	Action("list", func() {
		Description("TodoList一覧")
		Routing(GET("/"))
		Response(OK,CollectionOf(TodoMediaType))
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("Todoの詳細")
		Routing(GET("/:id"))
		Params(func() {
			Param("id")
			Required("id")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Description("Todoの作成")
		Routing(POST("/"))
		Params(func() {
			Param("body")
			Required("body")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Description("Todoの更新")
		Routing(PUT("/:id"))
		Params(func() {
			Param("id")
			Param("body")
			Param("is_finished")
			Required("id", "body", "is_finished")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Description("Todoの削除")
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id")
			Required("id")
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)

	})

	Action("delete all", func() {
		Description("Todoの全削除")
		Routing(DELETE("/"))
		Response(OK)
		Response(BadRequest, ErrorMedia)

	})
})

var TodoMediaType = MediaType("application/todo.+json", func() {
	Description("Todos")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})

		Attribute("body", String, "body", func() {
			Example("hoge")
		})

		Attribute("is_finished", Boolean, "is_finished", func() {
			Example(false)
		})
		Required("id", "body", "is_finished")
	})
	View("default", func() {
		Attribute("id")
		Attribute("body")
		Attribute("is_finished")
		Required("id", "body", "is_finished")
	})
})
