package main

import (
	"github.com/nomadcoderkor/dinocoin/cli"
	"github.com/nomadcoderkor/dinocoin/db"
)

func main() {
	// difficulty := 5
	// target := strings.Repeat("0", difficulty)
	// nonce := 1
	// for {
	// 	hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
	// 	if strings.HasPrefix(hash, target) {
	// 		fmt.Println(hash)
	// 		break
	// 	} else {
	// 		nonce++
	// 		fmt.Println(hash)
	// 	}
	// }

	// fmt.Printf("%x\n", hash)
	// Defer는 함수가 종료될때 실행 된다.
	// defer db.Close()
	// cli.Start()
	// wallet.Wallet()
	defer db.Close()
	cli.Start()
}
