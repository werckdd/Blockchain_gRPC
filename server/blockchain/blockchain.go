package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Hash string
	PrevBlockHash string
	Data string
}

type Blockchain struct{
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, preBlockHash string) *Block {
	block := &Block{
		Data: data,
		PrevBlockHash:preBlockHash,
	}
	block.setHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks) - 1]
	NewBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, NewBlock)

	return NewBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesisi Block", "")
}
