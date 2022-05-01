package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	newBlock.hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data+newBlock.prevHash)))
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("PrevHash : %s\n", block.prevHash)
		fmt.Println("-------------")
	}
}

func main() {

	chain := blockchain{}
	chain.addBlock("Nebula genesis block")
	chain.addBlock("Second block")
	chain.addBlock("Third block")
	chain.listBlocks()

	//genesisBlock := block{"Nebula's genesis block", "", ""}
	//hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	//hexHash := fmt.Sprintf("%x", hash)
	//genesisBlock.hash = hexHash
	//
	//fmt.Println(genesisBlock)
}
