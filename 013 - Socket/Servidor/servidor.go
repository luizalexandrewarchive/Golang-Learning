package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
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
			quantidadeShow := getQuantidadeShow()

			time.Sleep(100 * time.Millisecond)
			conn.Write([]byte(quantidadeShow + "\n"))
			time.Sleep(100 * time.Millisecond)
			conn.Write([]byte("1 - Lambda Lambda Lambda (R$120,00)\n"))
			time.Sleep(100 * time.Millisecond)
			conn.Write([]byte("2 - Hora do show pohaa (R$320,00)\n"))
			time.Sleep(100 * time.Millisecond)
			conn.Write([]byte("3 - Processinho a galope (R$560,00)\n"))
			time.Sleep(100 * time.Millisecond)
			conn.Write([]byte("4 - BIRRRRRRR (R$890,00)\n"))
			time.Sleep(100 * time.Millisecond)

			compra, _ := bufio.NewReader(conn).ReadString('\n')
			compraString, _ := strconv.Atoi(compra)
			fmt.Println("\nCompra realizada pelo usuario: ", email)
			fmt.Println("ID Show comprado:", compra)

			conn.Write([]byte("Item " + string(compraString) + "Comprado com sucesso\n"))

		} else {
			conn.Write([]byte(string("Usuário ou senha errada\n")))
		}

	}
}

func getQuantidadeShow() string {
	return "4"
}

func validar(email string, senha string) bool {

	if email == "luiz@live.com\n" && senha == "123\n" {
		return true
	}
	return false
}
