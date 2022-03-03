package BLG

import (
	"fmt"
	"testing"
)

func TestNewDbHelper(t *testing.T) {
	db := NewDbHelper("../db/test.db")
	//err := db.Put([]byte("test"), []byte("test2"))
	//if err != nil {
	//	return
	//}
	data, _ := db.Get([]byte("test"))
	fmt.Println(string(data))

	has, err := db.Has([]byte("test"))
	if err != nil {
		return
	}
	fmt.Println(has)
}
