package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

func IntToBytes(n int64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

func BytesToInt64(t []byte) int64 {
	return int64(binary.BigEndian.Uint64(t))
}

func Json2StrArray(jsonStr string) []string {
	var sArr []string
	if err := json.Unmarshal([]byte(jsonStr), &sArr); err != nil {
		log.Panic(err)
	}
	return sArr
}
