package blockchain

import (
	"fmt"
	"sync"

	"github.com/nomadcoderkor/dinocoin/db"
	"github.com/nomadcoderkor/dinocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
	// utils.HandleErr(gob.NewDecoder(bytes.NewReader(data)).Decode(b))
}

func (b *blockchain) persist() {

	db.SaveBlockChain(utils.ToBytes(b))

}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

// Blockchain Define
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			// step1 : DB에 저장된 블록이 있는지 확인하고, 있으면 가져온다
			checkpoint := db.CheckPoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				// step2 : byte로 저장되어있는 블록정보를 복원한다.
				b.restore(checkpoint)
			}
		})
	}
	fmt.Printf("22 NewestHash : %s\nHeight : %d\n", b.NewestHash, b.Height)
	return b
}
