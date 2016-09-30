package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter();
	router := r.PathPrefix("/api").Subrouter()
	for _, route := range routes {
		router.Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
	}

	return router
}

