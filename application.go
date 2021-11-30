package main

import (
	_server "github.com/nomada-sh/levita-stp/server"
)

const PORT = ":9000"

func main() {
	_server.NewRouter(PORT)
}
