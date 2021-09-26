package main

import "github.com/dinocoin/blockchain"

// func (chain *blockchain) getLastBlock() block {
// 	if len(chain.blocks) > 0 {
// 		lastBlock := chain.blocks[len(chain.blocks)-1]
// 		return lastBlock
// 	}
// 	return undefined
// }

// func (chain *blockchain) getLastHash() string {
// 	if len(chain.blocks) > 0 {
// 		return chain.blocks[len(chain.blocks)-1].hash
// 	}
// 	return ""
// }

// func (chain *blockchain) addBlock(data string) {
// 	newBlock := block{data, "", chain.getLastHash()}
// 	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
// 	newBlock.hash = fmt.Sprintf("%x", hash)
// 	chain.blocks = append(chain.blocks, newBlock)
// }

// func (chain *blockchain) listBlocks() {
// 	for _, block := range chain.blocks {
// 		fmt.Printf("Data : %s\n", block.data)
// 		fmt.Printf("Hash : %s\n", block.hash)
// 		fmt.Printf("prevHash : %s\n", block.prevHash)
// 	}
// }
func main() {
	chain := blockchain.GetBlockchain()
}
