package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nomadcoderkor/dinocoin/utils"

	"github.com/nomadcoderkor/dinocoin/blockchain"
)

const port string = ":4000"

// URL api url type
type URL string

// MarshalText Test
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// URLDescription API Description
type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// AddBlockBody Add block Post blocks API
type AddBlockBody struct {
	Message string
}

func (u URLDescription) String() string {
	return "return string"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "GET",
			Description: "See All Block",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Create Block",
			Payload:     "data:string",
		},
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := blockchain.GetBlockchain().AllBlocks()
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(data)
	case "POST":
		var addBlockBody AddBlockBody
		// 아래 코드는 post로 받은 메세지를 addBlockBody에 바인딩을 해준다는 의미
		// 포인터를 사용한 이유는 Decode 함수는 포인터를 받게 되어있다.
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))

		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)

		rw.WriteHeader(http.StatusCreated)
	}

}

func main() {
	fmt.Println(URLDescription{
		URL:         "/",
		Method:      "GET",
		Description: "See Documentation",
	})
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Start Server http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
