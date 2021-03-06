// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "todoz": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/konchanSS/Todoz/design
// --out=$(GOPATH)/src/github.com/konchanSS/Todoz
// --version=v1.2.0-dirty

package app

import (
	"fmt"
	"strings"
)

// TodosHref returns the resource href.
func TodosHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/todos/%v", paramid)
}
