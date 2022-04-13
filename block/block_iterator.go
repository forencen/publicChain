package block

import "publicChain/db"

type ChainIterator struct {
	Cursor []byte
	db     *db.DbHelper
}

func (it *ChainIterator) Next() *Block {
	if it.Cursor == nil {
		return nil
	}
	cursorBlockBytes, _ := it.db.Get(it.Cursor)
	if cursorBlockBytes == nil {
		return nil
	}
	cursorBlock := Deserialize(cursorBlockBytes)
	it.Cursor = cursorBlock.PrevBlockHash
	return cursorBlock
}
