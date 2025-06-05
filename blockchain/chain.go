package blockchain

type Blockchain struct {
	Blocks []*Block
}

// function
// create new blockchain -> initializes with genesis block
func NewBlockchain() *Blockchain {
	// Create genesis block
	genesisBlock := NewBlock(0, "", "", "")
	return &Blockchain{
		Blocks: []*Block{genesisBlock},
	}
}

// add block
func (bc *Blockchain) AddBlock(data string, validator string) {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(prevBlock.Index + 1, data, prevBlock.CurrentHash, validator)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// validate chain
func (bc *Blockchain) ValidateChain() bool {
	// check prevHash in current block
	// check currentHash
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if prevBlock.CurrentHash != currentBlock.PrevHash {
			return false
		}

		if currentBlock.CurrentHash != currentBlock.CalculateHash() {
			return false
		}
	}
	return true
}
