package entrypoint

import (
	"html/template"
	"net/http"
)

var tpl_login = template.Must(template.ParseFiles("static/login_page/index.html"))

func HandlerLoginPage(w http.ResponseWriter, r *http.Request) {
	tpl_login.Execute(w, nil)
}
