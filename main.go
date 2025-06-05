package main

import ("github.com/dheerajkumardk/pos-consensus/blockchain"
//  "github.com/dheerajkumardk/pos-consensus/consensus"
//  "github.com/dheerajkumardk/pos-consensus/network"
)

func main() {
	print("hello!");

	// call package blockchain
	blockchain.Block();

	blockchain.Chain();
}