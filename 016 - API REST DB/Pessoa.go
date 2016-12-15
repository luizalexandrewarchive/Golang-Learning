package main

type Pessoa struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Cidade string `json:"cidade"`
}

type Pessoas []Pessoa
