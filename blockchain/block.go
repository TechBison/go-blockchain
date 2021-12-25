package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Slice of blocks to form block chain
type BlockChain struct {
	Blocks []*Block
}

// Basic Block Structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Function to add Block in blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
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
