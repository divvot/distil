package main

import (
	ds "github.com/divvot/distil/server"
)

const PORT = ":5500"

func main() {
	ds.Serve(PORT)
}
