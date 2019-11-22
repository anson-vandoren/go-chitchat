package main

import (
	"net/http"

	"github.com/anson-vandoren/gwp/chitchat/data"
)

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

// GET /signup
// Show the signup page
func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
// Create the user account
func signupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := data.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if err := data.CheckPassword(user.Password, r.PostFormValue("password")); err == nil {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// GET /logout
// Logs the user out
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
