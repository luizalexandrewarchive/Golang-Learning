package main

import "net"
import "fmt"
import "bufio"
import "os"

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
		fmt.Print(message)
	}
}
