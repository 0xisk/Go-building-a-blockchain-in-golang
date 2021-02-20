package main

import (
	"fmt"
	"rsc.io/quote"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func main() {
	fmt.Println(quote.Hello())
}
