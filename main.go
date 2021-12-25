package main

import (
	"fmt"
)

func main() {

	chain := InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("data in Block : %s \n", block.Data)
		fmt.Printf("Prev Hash : %x\n", block.PrevHash)
		fmt.Printf("Current Hash of block : %x\n", block.Hash)
	}

}

func InitBlockChain() {
	panic("unimplemented")
}
