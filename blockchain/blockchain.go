package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block struct
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	Blocks []*Block
}

// Derive hash from data and prevHash data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Create a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return &block
}

// add block to chain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}

// Generate first block (Genesis Block)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initialize blockchain
func InitBlockchain() *BlockChain {
	return &BlockChain{
		Blocks: []*Block{Genesis()},
	}
}
