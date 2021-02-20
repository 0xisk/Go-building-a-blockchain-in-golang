package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"rsc.io/quote"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash();
	return block
}

func main() {
	fmt.Println(quote.Hello())
}
