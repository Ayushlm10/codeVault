package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", fileServer))

	scsMiddleware := alice.New(a.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", scsMiddleware.ThenFunc(a.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", scsMiddleware.ThenFunc(a.snippetView))
	router.Handler(http.MethodGet, "/snippet/create", scsMiddleware.ThenFunc(a.snippetCreate))
	router.Handler(http.MethodPost, "/snippet/create", scsMiddleware.ThenFunc(a.snippetCreatePost))

	router.Handler(http.MethodGet, "/user/signup", scsMiddleware.ThenFunc(a.userSignUp))
	router.Handler(http.MethodPost, "/user/signup", scsMiddleware.ThenFunc(a.userSignUpPost))
	router.Handler(http.MethodGet, "/user/login", scsMiddleware.ThenFunc(a.userLogin))
	router.Handler(http.MethodPost, "/user/login", scsMiddleware.ThenFunc(a.userLoginPost))
	router.Handler(http.MethodPost, "/user/logout", scsMiddleware.ThenFunc(a.userLogout))

	standardMiddleware := alice.New(a.recovcerPanic, a.logRequest, secureHeaders)
	return standardMiddleware.Then(router)
}
