package main

import (
	"flag"
	"fmt"
	// "github.com/Abhishek-Nagarkoti/master-slave/data"
	"github.com/Abhishek-Nagarkoti/master-slave/master"
	"github.com/Abhishek-Nagarkoti/master-slave/slave"
	"net"
	"strings"
)

/* The entry point for our System */
func main() {
	/* Parse the provided parameters on command line */
	makeMasterOnError := flag.Bool("makeMasterOnError", false, "make this node master if unable to connect to the cluster ip provided.")
	clusterip := flag.String("clusterip", "127.0.0.1:8001", "ip address of any node to connnect")
	myport := flag.String("myport", "8001", "ip address to run this node on. default is 8001.")
	flag.Parse()

	myIp, _ := net.InterfaceAddrs()

	if *makeMasterOnError {
		fmt.Println("Will start this node as master.")
		master.CreateMaster(myIp[0].String(), *myport)
	} else {
		slave.CreateSlave(strings.Split(*clusterip, ":")[0], strings.Split(*clusterip, ":")[1])
	}
}
