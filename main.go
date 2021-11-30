package main

import (
	"Blockchain/database"
	"Blockchain/gateway"
	"Blockchain/settings"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kardianos/service"
	"html/template"
	"net/http"
	"os"
	"time"
)

type program struct{}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	fmt.Printf("http://164.90.238.31:%s\n", os.Getenv("SERVER_PORT"))
	fmt.Printf("http://localhost:%s\n", os.Getenv("SERVER_PORT"))
	fmt.Printf("http://127.0.0.1:%s", os.Getenv("SERVER_PORT"))
	settings.WritingSync.Lock()
	settings.ServiceIsRunning = true
	settings.Db, _ = database.Connection()
	settings.WritingSync.Unlock()
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	settings.WritingSync.Lock()
	settings.ServiceIsRunning = false
	settings.Db.Close()
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
	mux.HandleFunc("/", HandlerLoginPage)
	mux.HandleFunc("/homepage", HandlerHomePage)
	mux.HandleFunc("/transactions", HandlerTransactions)
	mux.HandleFunc("/api_docs", HandlerAPIDocs)

	// API for frontend
	mux.HandleFunc("/new_user", gateway.NewUserGateway)
	mux.HandleFunc("/get_user_info", gateway.GetUserInfoGateway)
	mux.HandleFunc("/login_user", gateway.LoginUserGateway)
	mux.HandleFunc("/new_transaction", gateway.NewTransactionGateway)
	mux.HandleFunc("/random_wallet", gateway.RandomWalletGateway)
	mux.HandleFunc("/latest_transactions", gateway.GetLatestTransactionsGateway)
	mux.HandleFunc("/user_transactions", gateway.GetUserTransactionsGateway)

	// API
	mux.HandleFunc("/api/all_transactions", gateway.GetAllTransactionsGateway)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), mux)
	if err != nil {
		fmt.Println("Problem starting web server: " + err.Error())
		os.Exit(-1)
	}
}

var tpl_api_docs = template.Must(template.ParseFiles("static/api_docs_page/index.html"))

func HandlerAPIDocs(w http.ResponseWriter, r *http.Request) {
	tpl_api_docs.Execute(w, nil)
}

var tpl_transactions = template.Must(template.ParseFiles("static/transactions_page/index.html"))

func HandlerTransactions(w http.ResponseWriter, r *http.Request) {
	tpl_transactions.Execute(w, nil)
}

var tpl_home = template.Must(template.ParseFiles("static/homepage/index.html"))

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	tpl_home.Execute(w, nil)
}

var tpl_login = template.Must(template.ParseFiles("static/login_page/index.html"))

func HandlerLoginPage(w http.ResponseWriter, r *http.Request) {
	tpl_login.Execute(w, nil)
}

func main() {

	err := godotenv.Load(".env")

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
