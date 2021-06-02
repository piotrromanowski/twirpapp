package main

import (
	"fmt"

	"github.com/piotrromanowski/twirpapp/rpc"
)

func main() {
	url := &rpc.LongUrl{Url: "Hello World"}

	fmt.Println(url)
}
