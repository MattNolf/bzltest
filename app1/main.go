package main

import (
	"fmt"
	_ "cloud.google.com/go/pubsub"
	_ "cloud.google.com/go/storage"
)

func main() {
	fmt.Println("Hello world")
}
