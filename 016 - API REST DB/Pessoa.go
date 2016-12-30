package main

type Pessoa struct {
	ID    int    `json:"id,omitempty"`
	Nome  string `json:"nome,omitempty"`
	CPF   string `json:"cpf,omitempty"`
	Senha string `json:"senha,omitempty"`
	Email string `json:"email,omitempty"`
}

type Pessoas []Pessoa
