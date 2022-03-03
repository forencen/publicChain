package block

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChanWithGenesisBlock() *BlockChain {
	block := CreateGenesisBlock("block chain with genesis ")
	return &BlockChain{[]*Block{block}}
}

func (bc *BlockChain) AddBlockToBlockChan(data string) {
	nowBlock := bc.Blocks[len(bc.Blocks)-1]
	block := NewBlock(nowBlock.Height+1, data, nowBlock.Hash)
	bc.Blocks = append(bc.Blocks, block)
}
