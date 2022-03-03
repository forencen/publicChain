package BLG

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt64(t []byte) int64 {
	return int64(binary.BigEndian.Uint64(t))
}
