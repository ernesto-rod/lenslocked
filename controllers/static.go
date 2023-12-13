package controllers

import (
	"github.com/ernesto-rod/lenslocked/views"
	"net/http"
)

func StaticHanlder(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
