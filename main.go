package main

import (
	"github.com/nomadcoderkor/dinocoin/cli"
	"github.com/nomadcoderkor/dinocoin/db"
)

func main() {
	// Defer는 함수가 종료될때 실행 된다.
	defer db.Close()
	cli.Start()
}
