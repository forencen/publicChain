package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const walletFile = "wallet.dat"
const addressChecksumLen = 4

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	return &Wallet{private, public}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

func TransformPublicKey(pubKey []byte) []byte {
	pubKeyHash := HashPubKey(pubKey)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checkSum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checkSum...)
	return Base58Encode(fullPayload)
}

// GetAddressPubKeyHash 更具string的地址获取地址的 哈希过的的公钥
func GetAddressPubKeyHash(address string) []byte {
	fullPayload := Base58Decode([]byte(address))
	return fullPayload[1 : len(fullPayload)-4]
}

func TransformPublicKeyHash(pubKeyHash []byte) []byte {
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checkSum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checkSum...)
	return Base58Encode(fullPayload)
}

// GetAddress 公钥获取流程：
// base58(version + ripemd160(sha256(PublicKey)) + sha256(sha256(ripemd160(sha256(PublicKey)))))
func (w *Wallet) GetAddress() []byte {
	return TransformPublicKey(w.PublicKey)
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressChecksumLen]
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	RIPEMD160Hasher.Write(publicSHA256[:])
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	return publicRIPEMD160
}

// ValidateAddress 校验地址是否合法
// 用base58解码地址，拿出原始公钥，配合version 做校验和，和输入地址的后四位做比较
func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	inputCheckSum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	targetChecksum := checksum(pubKeyHash[:len(pubKeyHash)-addressChecksumLen])
	return bytes.Compare(inputCheckSum, targetChecksum) == 0
}
