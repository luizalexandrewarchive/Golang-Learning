package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Para utilizar a API utilize /pessoas ou /pessoas/{id} :D")
}

func GetAllPessoas(w http.ResponseWriter, r *http.Request) {

	pessoas := Pessoas{
		Pessoa{ID: 01, Nome: "Luiz Alexandre de Sousa Freitas", Email: "luiz.alexandre@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 02, Nome: "Luana Almeida Goncalves", Email: "luana.arte@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 03, Nome: "Eduardo Oliveira Cavalcanti", Email: "EduardoOliveiraCavalcanti@dayrep.com", Cidade: "Morrinhos"},
		Pessoa{ID: 04, Nome: "Daniel Goncalves Sousa", Email: "daniel.golsalves@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 05, Nome: "Gabrielly Rodrigues Santos", Email: "GabriellyRodriguesSantos@rhyta.com", Cidade: "Morrinhos"},
		Pessoa{ID: 06, Nome: "Sofia Lima Alves", Email: "SofiaLimaAlves@rhyta.com", Cidade: "Morrinhos"},
		Pessoa{ID: 07, Nome: "Caio Ribeiro Ferreira", Email: "CaioRibeiroFerreira@armyspy.com", Cidade: "Morrinhos"},
		Pessoa{ID: 10, Nome: "Carolina Goncalves Oliveira", Email: "CarolinaGoncalvesOliveira@dayrep.com", Cidade: "Morrinhos"},
		Pessoa{ID: 11, Nome: "Vinícius Pinto Santos", Email: "ViniciusPintoSantos@dayrep.com", Cidade: "Morrinhos"},
	}

	if err := json.NewEncoder(w).Encode(pessoas); err != nil {
		panic(err)
	}
}

func GetPessoaID(w http.ResponseWriter, r *http.Request) {

	pessoas := Pessoas{
		Pessoa{ID: 01, Nome: "Luiz Alexandre de Sousa Freitas", Email: "luiz.alexandre@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 02, Nome: "Luana Almeida Goncalves", Email: "luana.arte@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 03, Nome: "Eduardo Oliveira Cavalcanti", Email: "EduardoOliveiraCavalcanti@dayrep.com", Cidade: "Morrinhos"},
		Pessoa{ID: 04, Nome: "Daniel Goncalves Sousa", Email: "daniel.golsalves@live.com", Cidade: "Morrinhos"},
		Pessoa{ID: 05, Nome: "Gabrielly Rodrigues Santos", Email: "GabriellyRodriguesSantos@rhyta.com", Cidade: "Morrinhos"},
		Pessoa{ID: 06, Nome: "Sofia Lima Alves", Email: "SofiaLimaAlves@rhyta.com", Cidade: "Morrinhos"},
		Pessoa{ID: 07, Nome: "Caio Ribeiro Ferreira", Email: "CaioRibeiroFerreira@armyspy.com", Cidade: "Morrinhos"},
		Pessoa{ID: 10, Nome: "Carolina Goncalves Oliveira", Email: "CarolinaGoncalvesOliveira@dayrep.com", Cidade: "Morrinhos"},
		Pessoa{ID: 11, Nome: "Vinícius Pinto Santos", Email: "ViniciusPintoSantos@dayrep.com", Cidade: "Morrinhos"},
	}

	vars := mux.Vars(r)
	todoId := vars["pessoaId"]

	i, err := strconv.Atoi(todoId)

	if err == nil {
		json.NewEncoder(w).Encode(pessoas[i+1])
	} else {
		fmt.Fprintln(w, "ID inválido")
	}

}
