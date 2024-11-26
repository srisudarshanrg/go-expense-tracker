package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/srisudarshanrg/go-expense-tracker/server/database"
	"github.com/srisudarshanrg/go-expense-tracker/server/setup"
)

const portNumber = ":8500"

func main() {
	// session
	session := scs.New()
	session.Cookie.Persist = true
	session.Lifetime = 1 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	// database setup
	db, err := database.CreateDatabaseConn()
	if err != nil {
		log.Fatal(err)
	}

	setup.DBAccess(db)

	// routes
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	log.Println("server running on port number", portNumber)
	server.ListenAndServe()
}

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/login", setup.Login)
	mux.Get("/register", setup.Register)
	mux.Get("/expenses", setup.Expenses)
	mux.Get("/tracker", setup.Tracker)
	mux.Get("/budget", setup.Budget)
	mux.Get("/profile", setup.Budget)
	mux.Get("/logout", setup.Logout)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
