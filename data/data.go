package data

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	// "strconv"
)

type MasterNode struct {
	Connections []net.Conn
}

type Slices struct {
	Slices []string
}

var MasterNodeVar MasterNode

func GetJSON() Slices {
	jsonFile, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data Slices
	json.Unmarshal(byteValue, &data)
	return data
}
