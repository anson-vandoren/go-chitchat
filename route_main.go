package main

import (
	"net/http"

	"github.com/anson-vandoren/gwp/chitchat/data"
)

// GET /err?msg=
// shows the error message page
func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "navbar", "public", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "navbar", "private", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, "Cannot get threads")
	} else {
		_, err = session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "public", "index")
		} else {
			generateHTML(w, threads, "layout", "navbar", "private", "index")
		}
	}
}
