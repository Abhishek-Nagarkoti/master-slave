package master

import (
	"bufio"
	"fmt"
	"github.com/Abhishek-Nagarkoti/master-slave/data"
	"net"
	"strconv"
	"strings"
)

func handleConnection(conn net.Conn, str chan string) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	data.MasterNodeVar.Connections = append(data.MasterNodeVar.Connections, conn)
	if len(data.MasterNodeVar.Connections) == 3 {
		go SendToSlaves(str)
	}
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		// output message received
		if err != nil && err.Error() == "EOF" {
			for i := 0; i < len(data.MasterNodeVar.Connections); i++ {
				if data.MasterNodeVar.Connections[i] == conn {
					data.MasterNodeVar.Connections = append(data.MasterNodeVar.Connections[:i], data.MasterNodeVar.Connections[i+1:]...)
				}
			}
			break
		}
		str <- string(message)
		if len(str) == 3 {
			close(str)
		}
	}
	conn.Close()
}

func CreateMaster(address, port string) {
	/* Listen for incoming messages */
	ln, _ := net.Listen("tcp", ":"+port)
	data.MasterNodeVar = data.MasterNode{[]net.Conn{}}
	str := make(chan string, 3)
	/* accept connection on port */
	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", err)
			}
		} else {
			go handleConnection(connIn, str)
		}
	}
}

func SendToSlaves(str chan string) {
	slices := data.GetJSON()
	indice := len(slices.Slices) / 3
	var message []string
	message = append(message, strings.Join(slices.Slices[:indice], ","))
	message = append(message, strings.Join(slices.Slices[indice:(indice*2)], ","))
	message = append(message, strings.Join(slices.Slices[(indice*2):], ","))
	// send new string back to client
	for i := 0; i < len(data.MasterNodeVar.Connections); i++ {
		data.MasterNodeVar.Connections[i].Write([]byte(message[i] + "\n"))
	}
	Sort(str)
}

func Sort(str chan string) {
	var sortedList []string
	var slices [][]string
	var i, j, k int = 0, 0, 0
	for {
		if len(str) == 3 {
			for val := range str {
				slices = append(slices, strings.Split(strings.ReplaceAll(val, "\n", ""), ","))
			}
			// not implemented this part
			for i < len(slices[0]) && j < len(slices[1]) && k < len(slices[2]) {
				var1, _ := strconv.Atoi(slices[0][i])
				var2, _ := strconv.Atoi(slices[1][j])
				var3, _ := strconv.Atoi(slices[2][k])
				if var1 < var2 && var1 < var3 {
					sortedList = append(sortedList, slices[0][i])
					i++
				}
				if var2 < var1 && var2 < var3 {
					sortedList = append(sortedList, slices[1][j])
					j++
				}
				if var3 < var2 && var3 < var1 {
					sortedList = append(sortedList, slices[2][k])
					k++
				}
			}
			for i < len(slices[0]) && j < len(slices[1]) {
				var1, _ := strconv.Atoi(slices[0][i])
				var2, _ := strconv.Atoi(slices[1][j])
				if var1 < var2 {
					sortedList = append(sortedList, slices[0][i])
					i++
				} else {
					sortedList = append(sortedList, slices[1][j])
					j++
				}
			}
			for j < len(slices[1]) && k < len(slices[2]) {
				var1, _ := strconv.Atoi(slices[1][j])
				var2, _ := strconv.Atoi(slices[2][k])
				if var1 < var2 {
					sortedList = append(sortedList, slices[1][j])
					j++
				} else {
					sortedList = append(sortedList, slices[2][k])
					k++
				}
			}
			for i < len(slices[0]) && k < len(slices[2]) {
				var1, _ := strconv.Atoi(slices[0][i])
				var2, _ := strconv.Atoi(slices[2][k])
				if var1 < var2 {
					sortedList = append(sortedList, slices[0][i])
					i++
				} else {
					sortedList = append(sortedList, slices[2][k])
					k++
				}
			}
			for i < len(slices[0]) {
				sortedList = append(sortedList, slices[0][i])
				i++
			}
			for j < len(slices[1]) {
				sortedList = append(sortedList, slices[1][j])
				j++
			}
			for k < len(slices[2]) {
				sortedList = append(sortedList, slices[2][k])
				k++
			}
			break
		}
	}
	fmt.Println("sorted list", sortedList)
}
