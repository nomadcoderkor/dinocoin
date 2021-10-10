package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

// HandleErr Error Handler
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// ToBytes Interface
func ToBytes(i interface{}) []byte {
	var utilBuffer bytes.Buffer
	encoder := gob.NewEncoder(&utilBuffer)
	err := encoder.Encode(i)
	HandleErr(err)
	return utilBuffer.Bytes()
}

// FromBytes Interface
func FromBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}
