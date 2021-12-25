package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Slice of blocks to form block chain
type BlockChain struct {
	blocks []*Block
}

// Basic Block Structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Function to add Block in blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Function to calculate hash of each added block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create a block with data and hash of prev block as input
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// First Block in Block Chain seed Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Init a new BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("data in Block : %s \n", block.Data)
		fmt.Printf("Prev Hash : %x\n", block.PrevHash)
		fmt.Printf("Current Hash of block : %x\n", block.Hash)
	}

}
