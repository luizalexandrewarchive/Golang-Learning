package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllPessoas",
		"GET",
		"/pessoas",
		GetAllPessoas,
	},
	Route{
		"GetPessoaID",
		"GET",
		"/pessoas/{pessoaId}",
		GetPessoaID,
	},
	Route{
		"DeletePessoaID",
		"DELETE",
		"/pessoas/{pessoaId}",
		DeletePessoaID,
	},
	Route{
		"CriarPessoa",
		"POST",
		"/pessoas",
		CriarPessoa,
	},
	Route{
		"AlterarPessoa",
		"PATCH",
		"/pessoas/{pessoaId}",
		AlterarSenhaPessoa,
	},
}
