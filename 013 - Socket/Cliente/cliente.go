package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Email: ")
		email, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, email+"\n")
		fmt.Print("Senha: ")
		senha, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, senha+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "Logado com sucesso\n" {
			fmt.Println(message)
			break
		} else {
			fmt.Println(message)
		}
	}
	fmt.Println("----------------Buscando Shows-----------------")
	qtnShow, _ := bufio.NewReader(conn).ReadString('\n')

	qtdShowInt, _ := strconv.Atoi(qtnShow[:len(qtnShow)-1])

	for count := 0; qtdShowInt > count; count++ {
		show, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(show[:len(show)-1])

	}
	fmt.Println("-----Identifique qual show quer comprar--------")
	compra, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, compra+"\n")

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)
}
