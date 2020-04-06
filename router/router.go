package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	makeStaticRouter(r)
	makeWebViewRouter(r)
	makeApiRouter(r)
	return r
}

func makeApiRouter(mainRouter *mux.Router) {
	apiRouter := mainRouter.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(jsonResponseMiddleware)
	//apiRouter.HandleFunc("/health", JSONHandler(statusHandler)).Methods("GET")
	apiRouter.HandleFunc("/login", JSONHandler(loginHandler)).Methods("POST")

	authRouter := apiRouter.PathPrefix("").Subrouter()
	authRouter.Use(authorizeByCookie)

	apiRouter.HandleFunc("/logout", JSONHandler(logoutHandler)).Methods("GET")
	apiRouter.HandleFunc("/status", JSONHandler(userStatusHandler)).Methods("GET")
	authRouter.HandleFunc("/temperature", JSONHandler(userTemperatureHandler)).Methods("GET")
	authRouter.HandleFunc("/pressure", JSONHandler(userPressureHandler)).Methods("GET")
	authRouter.HandleFunc("/humidity", JSONHandler(userHumidityHandler)).Methods("GET")
	authRouter.HandleFunc("/weather", JSONHandler(userWeatherHandler)).Methods("GET")

	adminRouter := authRouter.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/battery", JSONHandler(authHandler)).Methods("GET")
	adminRouter.HandleFunc("/info", JSONHandler(authHandler)).Methods("GET")
	adminRouter.HandleFunc("/work", JSONHandler(authHandler)).Methods("GET")

}

func makeStaticRouter(mainRouter *mux.Router) {
	mainRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}

func makeWebViewRouter(mainRouter *mux.Router) {
	webRouter := mainRouter.PathPrefix("/").Subrouter()

	webRouter.HandleFunc("/login", loginWebHandler).Methods("GET")

	webAuthRouter := webRouter.PathPrefix("").Subrouter()
	//webAuthRouter.Use(authorizeByCookie)
	//User webView
	webAuthRouter.HandleFunc("/", viewWebHandler).Methods("GET")
	webAuthRouter.HandleFunc("/logout", JSONHandler(logoutHandler)).Methods("GET")
}
