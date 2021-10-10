package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/nomadcoderkor/dinocoin/db"
	"github.com/nomadcoderkor/dinocoin/utils"
)

// Block Define
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

// func (b *Block) toBytes() []byte {
// 	var blockBuffer bytes.Buffer
// 	encoder := gob.NewEncoder(&blockBuffer)
// 	err := encoder.Encode(b)
// 	utils.HandleErr(err)
// 	return blockBuffer.Bytes()
// }

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

// ErrorNotFound Find block not found Error Message
var ErrorNotFound = errors.New("Block not found")

// FindBlock 해쉬값으로 블록 정보를 찾는다.
func FindBlock(hash string) (*Block, error) {
	blockByte := db.Block(hash)
	if blockByte == nil {
		return nil, ErrorNotFound
	}
	block := &Block{}
	block.restore(blockByte)
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
