package main

import (
	"github.com/goadesign/goa"
	"github.com/konchanSS/Todoz/app"
)

// TodosController implements the todos resource.
type TodosController struct {
	*goa.Controller
}

// NewTodosController creates a todos controller.
func NewTodosController(service *goa.Service) *TodosController {
	return &TodosController{Controller: service.NewController("TodosController")}
}

// Create runs the create action.
func (c *TodosController) Create(ctx *app.CreateTodosContext) error {
	// TodosController_Create: start_implement

	// Put your logic here

	// TodosController_Create: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *TodosController) Delete(ctx *app.DeleteTodosContext) error {
	// TodosController_Delete: start_implement

	// Put your logic here

	// TodosController_Delete: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
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

	// TodosController_List: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *TodosController) Show(ctx *app.ShowTodosContext) error {
	// TodosController_Show: start_implement

	// Put your logic here

	// TodosController_Show: end_implement
	return nil
}

// Update runs the update action.
func (c *TodosController) Update(ctx *app.UpdateTodosContext) error {
	// TodosController_Update: start_implement

	// Put your logic here

	// TodosController_Update: end_implement
	res := &app.Todo{}
	return ctx.OK(res)
}
