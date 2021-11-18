package gateway

import (
	"log"
	"net/http"
)

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Query())
	login := r.URL.Query()["login"]
	pass := r.URL.Query()["pass"]

	//if !ok || len(keys[0]) = 1 {
	//	log.Println("Not all of the arguments")
	//	return
	//}

	log.Println("Url Param 'login' is: " + string(login[0]))
	log.Println("Url Param 'pass' is: " + string(pass[0]))
}
