package controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

type AppContext struct {
	PageData map[string]any
	Context  context.Context
	Scope    string
	TraceID  string
}

func Get(r *http.Request, Scope string) (*AppContext, error) {

	AP := AppContext{}
	AP.Context = r.Context()
	AP.Scope = Scope

	AP.TraceID = middleware.GetReqID(AP.Context)
	AP.PageData = make(map[string]any)

	return &AP, nil
}
