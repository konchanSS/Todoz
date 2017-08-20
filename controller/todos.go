package controller

import (
	"github.com/goadesign/goa"
	"github.com/konchanSS/Todoz/app"
	"github.com/konchanSS/Todoz/models"
	"github.com/jinzhu/gorm"
)

// TodosController implements the todos resource.
type TodosController struct {
	*goa.Controller
	db *gorm.DB
}

// NewTodosController creates a todos controller.
func NewTodosController(service *goa.Service, db *gorm.DB) *TodosController {
	return &TodosController{
		Controller: service.NewController("TodosController"),
		db: db,
	}
}

// Create runs the create action.
func (c *TodosController) Create(ctx *app.CreateTodosContext) error {
	// TodosController_Create: start_implement

	// Put your logic here
	t := &models.Todo{}
	t.Body = ctx.Body
	t.IsFinished = false
	tdb := models.NewTodoDB(c.db)
	err := tdb.Add(ctx.Context, t)
	if err != nil {
		return ctx.BadRequest(err)
	}
	// TodosController_Create: end_implement
	return ctx.Created()
}

// Delete runs the delete action.
func (c *TodosController) Delete(ctx *app.DeleteTodosContext) error {
	// TodosController_Delete: start_implement

	// Put your logic here
	tdb := models.NewTodoDB(c.db)
	err := tdb.Delete(ctx.Context, ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// TodosController_Delete: end_implement
	return nil
}

// DeleteAll runs the delete all action.
func (c *TodosController) DeleteAll(ctx *app.DeleteAllTodosContext) error {
	// TodosController_DeleteAll: start_implement

	// Put your logic here

	// TodosController_DeleteAll: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *TodosController) List(ctx *app.ListTodosContext) error {
	// TodosController_List: start_implement

	// Put your logic here
	tdb := models.NewTodoDB(c.db)
	t := tdb.ListTodo(ctx.Context)

	// TodosController_List: end_implement
	res := app.TodoCollection{}
	res = t
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TodosController) Show(ctx *app.ShowTodosContext) error {
	// TodosController_Show: start_implement

	// Put your logic here
	tdb := models.NewTodoDB(c.db)
	t, err := tdb.Get(ctx.Context, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	// TodosController_Show: end_implement
	res := &app.Todo{}
	res = t.TodoToTodo()
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TodosController) Update(ctx *app.UpdateTodosContext) error {
	// TodosController_Update: start_implement

	// Put your logic here
	t := &models.Todo{}
	t.ID = ctx.ID
	t.Body = ctx.Body
	t.IsFinished = ctx.IsFinished
	tdb := models.NewTodoDB(c.db)
	err := tdb.Update(ctx.Context, t)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}else {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// TodosController_Update: end_implement
	return nil
}
