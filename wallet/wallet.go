package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/nomadcoderkor/dinocoin/utils"
)

const (
	fileName string = "dinocoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string `json:"address"`
}

var w *wallet

// 이미 생성된 월렛이 있는지 확인
func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

// 새로운 지갑 생성
func createPrivKey() *ecdsa.PrivateKey {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

// 지갑파일 생성
func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

// 키파일 복구
// 아래처럼 리턴할 변수명을 지정해줄 경우에는 return 변수명 해줄필요없이 return 만 해주면 됨
// 공식문서에서는 매우 짧은 함수에서만 사용하길 권장함
// 리턴값을 다시 확인해야 하는 상황이 발생할때 불편함
func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsByte, err := os.ReadFile(fileName)
	utils.HandleErr(err)
	key, err = x509.ParseECPrivateKey(keyAsByte)
	utils.HandleErr(err)
	return
}

func encodeBigInts(a, b []byte) string {
	z := append(a, b...)
	return fmt.Sprintf("%x", z)
}

func addressFromKey(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
	// z := append(key.X.Bytes(), key.Y.Bytes()...)
	// return fmt.Sprintf("%x", z)
}

// Sign .
func Sign(payload string, w *wallet) string {
	payloadAsByte, err := hex.DecodeString(payload)
	utils.HandleErr(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsByte)
	// r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, []byte(payload))
	utils.HandleErr(err)
	return encodeBigInts(r.Bytes(), s.Bytes())
	// signature := append(r.Bytes(), s.Bytes()...)
	// return fmt.Sprintf("%x", signature)
}

func restoreBigInts(signature string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(signature)
	if err != nil {
		return nil, nil, err
	}
	utils.HandleErr(err)
	firstHalfBytes := bytes[:len(bytes)/2]
	secondHalfBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(secondHalfBytes)
	return &bigA, &bigB, nil
}

// Verify .
func Verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)
	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	payloadBytes, err := hex.DecodeString(payload)
	ok := ecdsa.Verify(&publicKey, payloadBytes, r, s)

	return ok
}

// Wallet .
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		// 월렛이 존재하는지 확인
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else {
			key := createPrivKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = addressFromKey(w.privateKey)
	}
	return w
}

// Start .
func Start() {

}
