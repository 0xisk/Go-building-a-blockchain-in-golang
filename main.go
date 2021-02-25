package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/iskanderandrews/practicing-golang-blockchain/blockchain"
)

type CommandLine struct {
	blockchain *blockchain.Blockchain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to the blockchain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockchain.Iterator()

	for  {
		block := iter.Next()

		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printBlockCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("Block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handler(err)

	case "print":
		err := printBlockCmd.Parse(os.Args[2:])
		blockchain.Handler(err)

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printBlockCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockchain()
	defer chain.Database.Close()

	cli := CommandLine{chain}
	cli.run()
}
