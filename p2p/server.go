package p2p

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
)

const protocol = "tcp"
const commandLength = 12
const nodeVersion = 1

var nodePort = 1331
var nodeAddress string
var knownNodes = []string{"192.168.31.222"}

//var blocksInTransit = [][]byte{}
//var memPool = make(map[string]Transaction)

type PBlockChain interface {
	GetBestHeight() int64
}

func commandToBytes(command string) []byte {
	var commandBytes [commandLength]byte
	for no, i := range command {
		if no >= commandLength {
			return nil
		}
		commandBytes[no] = byte(i)
	}
	return commandBytes[:]
}

func bytesToCommand(command []byte) string {
	var c []byte
	for _, i := range command {
		if i != 0x0 {
			c = append(c, i)
		}
	}
	return fmt.Sprintf("%s", c)
}

func extractCommand(request []byte) []byte {
	return request[:commandLength]
}

func sendData(address string, data []byte) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Printf("%s is not available\n", address)
		// 如果此ip不能链接，清除掉knownNodes中的此节点信息
		var availableNodes []string
		for _, node := range knownNodes {
			if node != address {
				availableNodes = append(availableNodes, node)
			}
		}
		knownNodes = availableNodes
		return
	}
	defer conn.Close()
	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

func gobEncode(data interface{}) []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

func gobDecode(data []byte, typ interface{}) interface{} {
	var buffer bytes.Buffer
	buffer.Write(data)
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&typ)
	if err != nil {
		log.Panic(err)
	}
	return typ
}

func nodeIsKnown(addr string) bool {
	for _, node := range knownNodes {
		if node == addr {
			return true
		}
	}

	return false
}
