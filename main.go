package main

import (
	"fmt"

	"github.com/steelthedev/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockchain()

	chain.AddBlock("First block of my blockchain after genesis")
	chain.AddBlock("Second block of mu blockchain")
	chain.AddBlock("Third block mate")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in the block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
