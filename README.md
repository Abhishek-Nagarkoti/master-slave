Run master 

`go run main.go --makeMasterOnError --myport 8001`

Run slave

`go run main.go --clusterip 127.0.0.1:8001`