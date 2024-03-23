package web

import (
	"net/http"
)

func GetFormInt(r *http.Request, name string) int {
	return CleanInt(r.FormValue(name))
}
func GetFormString(r *http.Request, name string) string {
	return r.FormValue(name)
}
func GetQueryInt(r *http.Request, name string) int {
	return CleanInt(r.URL.Query().Get(name))
}
func GetQueryString(r *http.Request, name string) string {
	return r.URL.Query().Get(name)

}
