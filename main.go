package main

import (
	"Blockchain/database"
	"Blockchain/entrypoint"
	"Blockchain/gateway"
	"Blockchain/settings"
	"fmt"
	"github.com/kardianos/service"
	"net/http"
	"os"
	"time"
)

type program struct{}

var db, _ = database.Connection()

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	fmt.Println("http://localhost:6006/")
	settings.WritingSync.Lock()
	settings.ServiceIsRunning = true
	settings.WritingSync.Unlock()
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	settings.WritingSync.Lock()
	settings.ServiceIsRunning = false
	settings.WritingSync.Unlock()
	for settings.ProgramIsRunning {
		fmt.Println(s.String() + " stopping...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {

	fs := http.FileServer(http.Dir("static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Frontend
	mux.HandleFunc("/", entrypoint.HandlerLoginPage)
	mux.HandleFunc("/homepage", entrypoint.HandlerHomePage)

	// API for frontend
	mux.HandleFunc("/new_user", gateway.HandleNewUser)
	mux.HandleFunc("/get_user_info", gateway.HandleGetUserInfo)
	mux.HandleFunc("/login_user", gateway.HandleLoginUser)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), mux)
	if err != nil {
		fmt.Println("Problem starting web server: " + err.Error())
		os.Exit(-1)
	}
}

func main() {

	serviceConfig := &service.Config{
		Name:        "Blockchain",
		DisplayName: "Blockchain",
		Description: "Blockchain",
	}

	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}
