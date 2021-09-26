package main

import (
	"crypto/sha256"
	"fmt"
)

// Block 선언
type block struct {
	data     string
	hash     string
	prevHash string
}

// Blockchain 선언
// 이 체인이 성공을 하게 되었을때 chain이 얼마나 커질까? 궁금하다.
type blockchain struct {
	blocks []block
}

// func (chain *blockchain) getLastBlock() block {
// 	if len(chain.blocks) > 0 {
// 		lastBlock := chain.blocks[len(chain.blocks)-1]
// 		return lastBlock
// 	}
// 	return undefined
// }

func (chain *blockchain) getLastHash() string {
	if len(chain.blocks) > 0 {
		return chain.blocks[len(chain.blocks)-1].hash
	}
	return ""
}

func (chain *blockchain) addBlock(data string) {
	newBlock := block{data, "", chain.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	chain.blocks = append(chain.blocks, newBlock)
}

func (chain *blockchain) listBlocks() {
	for _, block := range chain.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("prevHash : %s\n", block.prevHash)
	}
}
func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
	// for _, byteItem := range "test byte array" {
	// 	fmt.Printf("%b\n", byteItem)
	// }
	// genesisBlock := block{"Genesis Block", "", ""}
	// hash := sha256.Sum256([]byte(genesisBlock.data))
	// fmt.Println(hash)
	// fmt.Printf(sha256.Sum256([]byte(genesisBlock.data)))
	// genesisBlock.hash = sha256.Sum256(genesisBlock.data)
}
