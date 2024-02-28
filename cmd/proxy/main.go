package main

import (
	"fmt"

	"github.com/pythonwithsean/Go-Proxy-Server/proxy"
)

func main() {

	s := proxy.NewServer("google.com")
	err := s.Listen()
	fmt.Printf("There was a problem and it was %s", err)

}
