package main

import (
	"fmt"

	"github.com/dheerajkumardk/pos-consensus/blockchain"
	// "github.com/dheerajkumardk/pos-consensus/consensus"
	// "github.com/dheerajkumardk/pos-consensus/network"
)

func main() {
	print("hello!");

	bc := blockchain.NewBlockchain()

	bc.AddBlock("Data 1", "Validator A")
	bc.AddBlock("Data 2", "Validator B")

	fmt.Printf("Blockchain valid? %t \n", bc.ValidateChain())

	println(bc.Blocks[1].CurrentHash)
	println(bc.Blocks[2].PrevHash)
}