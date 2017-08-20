// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "todoz": Models
//
// Command:
// $ goagen
// --design=github.com/konchanSS/Todoz/design
// --out=$(GOPATH)/src/github.com/konchanSS/Todoz
// --version=v1.2.0-dirty

package models

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/konchanSS/Todoz/app"
	"time"
)

// Todo Relational Model
type Todo struct {
	ID         int `gorm:"primary_key"` // primary key
	Body       string
	IsFinished bool
	CreatedAt  time.Time  // timestamp
	DeletedAt  *time.Time // nullable timestamp (soft delete)
	UpdatedAt  time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Todo) TableName() string {
	return "todos"

}

// TodoDB is the implementation of the storage interface for
// Todo.
type TodoDB struct {
	Db *gorm.DB
}

// NewTodoDB creates a new storage type.
func NewTodoDB(db *gorm.DB) *TodoDB {
	return &TodoDB{Db: db}
}

// DB returns the underlying database.
func (m *TodoDB) DB() interface{} {
	return m.Db
}

// TodoStorage represents the storage interface.
type TodoStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Todo, error)
	Get(ctx context.Context, id int) (*Todo, error)
	Add(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, todo *Todo) error
	Delete(ctx context.Context, id int) error

	ListTodo(ctx context.Context) []*app.Todo
	OneTodo(ctx context.Context, id int) (*app.Todo, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *TodoDB) TableName() string {
	return "todos"

}

// CRUD Functions

// Get returns a single Todo as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *TodoDB) Get(ctx context.Context, id int) (*Todo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "todo", "get"}, time.Now())

	var native Todo
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Todo
func (m *TodoDB) List(ctx context.Context) ([]*Todo, error) {
	defer goa.MeasureSince([]string{"goa", "db", "todo", "list"}, time.Now())

	var objs []*Todo
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *TodoDB) Add(ctx context.Context, model *Todo) error {
	defer goa.MeasureSince([]string{"goa", "db", "todo", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Todo", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *TodoDB) Update(ctx context.Context, model *Todo) error {
	defer goa.MeasureSince([]string{"goa", "db", "todo", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Todo", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *TodoDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "todo", "delete"}, time.Now())

	var obj Todo

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Todo", "error", err.Error())
		return err
	}

	return nil
}