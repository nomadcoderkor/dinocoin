package main

import (
	"fmt"

	"github.com/dinocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	for _, block := range chain.AllBlocks() {
		// fmt.Println(block.Data)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %s\n", block.Hash)
		fmt.Printf("PrevHash : %s\n", block.PrevHash)
	}
}
