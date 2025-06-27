package main

import (
	"github.com/noahzyl/mindmate/routers"
)

// Test gin
func main() {
	r := routers.Router()
	r.Run(":9999") // Listen on localhost:9999
}
