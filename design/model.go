package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("todos", func() {
	Description("todo Model")
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQL")
		Model("Todo", func() {
			RendersTo(TodoMediaType)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})

			Field("body", gorma.String)
			Field("is_finished", gorma.Boolean)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
	})
})
