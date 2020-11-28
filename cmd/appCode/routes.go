package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *appInjection) myRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", app.Welcome).Methods("GET")
	router.HandleFunc("/user/{username}", app.SingleUser).Methods("GET")
	router.HandleFunc("/signIn", app.SignIn).Methods("POST")
	router.HandleFunc("/signUp", app.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/address", app.SaveAddress).Methods("POST", "OPTIONS")

	// log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), 
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), 
	// 	handlers.AllowedOrigins([]string{"*"}))(router)))
	return router
}
