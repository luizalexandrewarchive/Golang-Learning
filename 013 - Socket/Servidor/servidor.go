package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Iniciando servidor de vendas...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	for {
		email, _ := bufio.NewReader(conn).ReadString('\n')
		senha, _ := bufio.NewReader(conn).ReadString('\n')

		isValido := validar(string(email), string(senha))

		if isValido == true {
			fmt.Println("Usuário Logado")
			conn.Write([]byte(string("Logado com sucesso\n")))
		} else {
			conn.Write([]byte(string("Usuário ou senha errada\n")))
		}

	}
}

func listarIngressos() {

}

func validar(email string, senha string) bool {

	if email == "luiz@live.com\n" && senha == "123\n" {
		return true
	}
	return false
}
