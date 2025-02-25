package main

import (
	"fmt"

	"go-blockchain/blockchain"
)

func main() {
	fmt.Println("Hello Crypto Zombies! Starting Blockchain... 🔗")

	// Create Genesis Block
	genesisBlock := blockchain.Block{
		Index:        0,
		Timestamp:    "2025-02-25 12:00:00",
		Data:         "Genesis Block @ HumbleBee",
		PreviousHash: "",
	}

	genesisBlock.Hash = blockchain.CalculateHash(genesisBlock)
	fmt.Println("Genesis Block: ", genesisBlock)

	// Second Block
	secondBlock := blockchain.GenerateBlock(genesisBlock, "Second Block to my Future")
	fmt.Println("Second Block: ", secondBlock)

	// Validation
	fmt.Println("Is Second Block valid?", blockchain.IsBlockValid(secondBlock, genesisBlock))
}
