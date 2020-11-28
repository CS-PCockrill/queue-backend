package main

import (

	"encoding/json"
	"fmt"
	// "time"
	"net/http"

	// "github.com/golang/gddo/httputil/header"
	// "github.com/CS-PCockrill/queue/pkg/forms"
	"github.com/CS-PCockrill/queue/pkg/models"
	"github.com/gorilla/mux"
	// "github.com/gomodule/redigo/redis"
)

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Context-Type", "application/x-www-form-urlencoded")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
 }

func (app *appInjection) Welcome(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome to Queue Delivery"))
	allUser := app.user.GetUsers()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(allUser)
}

func (app *appInjection) SingleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	user := app.user.GetUser(username)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (app *appInjection) SignUp(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		fmt.Println(w)
		fmt.Fprintln(w)
		return
	}

	// Guide addressing headers, syntax error's, and preventing extra data fields
	// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

	var newUser models.User
	//Parse the form data
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	//TODO: Validate the form

	//If there is no error and the form is validated, create a new user from http request
	//Insert the new user into the database
	err = app.user.Insert(
		newUser.UserName,
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		string(newUser.Password))

	fmt.Fprintln(w, "Record Inserted")

	//And redirect the user to the login page
	// _, _ = w.Write([]byte("Sign Up"))

}


func (app *appInjection) SignIn(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		fmt.Println(w)
		fmt.Fprintln(w)
		return
	}
	// err := r.ParseForm()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	//Check that user credentials are valid, if not send user a message,
	//then re-display the login page

	// form := forms.New(r.PostForm)
	id, err := app.user.Authenticate(user.Email, string(user.Password))
	//TODO: Placeholder to add ID of the current user to the session.
	
	//For now, just write the ID
	fmt.Fprintln(w, id)
	fmt.Fprintln(w, user.Email)

	//If SignIn is successful, redirect the user to the page that will be displayed.
}

func (app *appInjection) SaveAddress(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		fmt.Println(w)
		fmt.Fprintln(w)
		return
	}
	var address models.Address
	
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id := app.user.Update(address.Street, address.City, address.State, address.Zip)

	fmt.Println(id)
}
