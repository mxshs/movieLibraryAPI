package main

import (
	server "mxshs/movieLibrary/cmd"
)

func main() {
	server.BootstrapAPI("v1")
}
