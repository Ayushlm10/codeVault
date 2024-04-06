package main

import "net/http"

func (a *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet/create", a.snippetCreate)
	mux.HandleFunc("/snippet/view", a.snippetView)

	return mux
}
