package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nomadcoderkor/dinocoin/explorer"
	"github.com/nomadcoderkor/dinocoin/rest"
)

func usage() {
	fmt.Printf("Dino Coin에 오신걸 환영합니다.\n")
	fmt.Printf("아래의 명령어를 사용해 주세요.\n")
	fmt.Printf("-port: Start HTML Explorer.\n")
	fmt.Printf("-mode: Choose between 'html' and 'rest'.\n")
	os.Exit(0)
}

// Start CLI start function
func Start() {
	// if len(os.Args) < 2 {
	// 	usage()
	// }

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	fmt.Println(*port, *mode)

	switch *mode {
	case "html":
		fmt.Println("Start html")
		explorer.Start(*port)
	case "rest":
		fmt.Println("Start rest")
		rest.Start(*port)
	default:
		usage()
	}
}
