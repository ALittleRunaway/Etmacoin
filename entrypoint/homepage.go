package entrypoint

import (
	"html/template"
	"net/http"
)

var tpl_home = template.Must(template.ParseFiles("static/homepage/index.html"))

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	tpl_home.Execute(w, nil)
}
