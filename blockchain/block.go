package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// block
type Block struct {
	Index 		int			// height of the block
	Timestamp	time.Time	// time of block creation
	Data        string		// data
	PrevHash    string		// hash of prev block
	Validator   string		// ID of the validator who proposed this block
	CurrentHash string		// hash of this block (calculated)
}

// calculate Hash
func (b *Block) CalculateHash() string {
	input := string(b.Index) + b.Timestamp.String() + b.Data + b.PrevHash + b.Validator
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// new block
func NewBlock(index int, data string, prevHash string, validator string) *Block {
	block := &Block{
		Index: index,
		Timestamp: time.Now(),
		Data: data,
		PrevHash: prevHash,
		Validator: validator,
	}
	block.CurrentHash = block.CalculateHash()
	return block
}
