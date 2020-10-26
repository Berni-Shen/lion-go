package router

import (
	"net/http"
	"reflect"
	"strings"
)

type tree struct {
	node *map[string]treenode
}

type treenode struct {
	handleFuncs *[]handle
}

type handle struct {
	methods *[]method
}

type method struct {
	paramNum   int
	paramsType *[]reflect.Type
	f          func(args ...interface{})
}

func (t *tree) add(item *RouterItem) {
	if t.node == nil {
		t.node = &map[string]treenode{}
	}
	path := strings.ToLower(item.Route)
	n, ok := (*t.node)[path]
	for !ok {
		(*t.node)[path] = treenode{
			handleFuncs: &[]handle{
				GET:    {methods: &[]method{}},
				POST:   {methods: &[]method{}},
				PUT:    {methods: &[]method{}},
				DELETE: {methods: &[]method{}},
			},
		}
		n, ok = (*t.node)[path]
	}

	var h handle
	switch item.Method {
	case http.MethodGet:
		h = (*n.handleFuncs)[GET]
	case http.MethodPost:
		h = (*n.handleFuncs)[POST]
	case http.MethodPut:
		h = (*n.handleFuncs)[PUT]
	case http.MethodDelete:
		h = (*n.handleFuncs)[DELETE]
	default:
	}
	h.methods = addHandleFunc(h.methods, &item.F)
}
