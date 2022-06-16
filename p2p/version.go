package p2p

type Version struct {
	Version     int
	AddressFrom string
	BestHeight  int64
}

func SendVersion(bestHeight int64, address string) {
	payload := gobEncode(Version{nodeVersion, nodeAddress, bestHeight})
	sendData(address, append(commandToBytes("version"), payload...))
}

func HandleVersion(request []byte, bc PBlockChain) {
	var payload Version
	gobDecode(request[:commandLength], payload)
	// 查看自己的高度
	myBestHeight := bc.GetBestHeight()
	otherBestHeight := payload.BestHeight
	if myBestHeight < otherBestHeight {
		sendGetBlocks(payload.AddressFrom)
	} else if myBestHeight > otherBestHeight {
		SendVersion(myBestHeight, payload.AddressFrom)
	}
	if !nodeIsKnown(payload.AddressFrom) {
		knownNodes = append(knownNodes, nodeAddress)
	}
}
