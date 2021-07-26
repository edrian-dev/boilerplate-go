package main

import (
	_server "github.com/siends/siends-api/server"
)

const PORT = ":9000"

func main() {
	_server.NewRouter(PORT)
}
