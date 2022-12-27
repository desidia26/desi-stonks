package main

import (
	"fmt"
	"desi-stonks/pkg/app"
)

func main() {
	fmt.Println("hi mom")
	app.NewServer().Run()
}