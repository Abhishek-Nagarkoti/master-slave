package slave

import (
	"fmt"
	"net"
	// "os"
	"bufio"
	"strconv"
	"strings"
	"time"
)

func CreateSlave(address, port string) {
	// connect to this socket
	conn, err := net.DialTimeout("tcp", address+":"+port, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	for {
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err.Error() == "EOF" {
			break
		}
		arr := strings.Split(message, ",")
		sort(arr, conn)
	}
}

func sort(arr []string, conn net.Conn) {
	var temp string = ""
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			var1, _ := strconv.Atoi(arr[i])
			var2, _ := strconv.Atoi(arr[j])
			if var1 > var2 {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Fprintf(conn, strings.ReplaceAll(strings.Join(arr, ","), "\n", "")+"\n")
}
