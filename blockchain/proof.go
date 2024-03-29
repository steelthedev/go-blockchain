package blockchain

import "math/big"

//consensus algrorithms are proof of works, proof os state ...
// we use it to secure our blockchain by forcing the network to do work by adding a block to the chain

// Take the data from the block

// create a counter which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a sets of requirements

// Requirements:
// The first few bytes must contain 0s

const Difficulty = 12 // doesn't stay static in a real blockchain

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}
