package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/nomadcoderkor/dinocoin/blockchain"
	"github.com/nomadcoderkor/dinocoin/utils"
)

// const port string = ":4000"
var port string

// URL api url type
type url string

// MarshalText Test
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// urlDescription API Description
type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// addBlockBody Add block Post blocks API
type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func (u urlDescription) String() string {
	return "return string"
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Block",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Create Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
		},
	}
	// rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := blockchain.GetBlockchain().AllBlocks()
		// rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(data)
	case "POST":
		var addBlockBody addBlockBody
		// 아래 코드는 post로 받은 메세지를 addBlockBody에 바인딩을 해준다는 의미
		// 포인터를 사용한 이유는 Decode 함수는 포인터를 받게 되어있다.
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))

		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)

		rw.WriteHeader(http.StatusCreated)
	}

}

func getBlock(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id := vars["height"]
	// rw.WriteHeader(http.StatusOK)
	// fmt.Fprintf(rw, "ID: %v\n", vars["id"])
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}

}

// Start function
func Start(aPort int) {
	// http.ListenAndServe(fmt.Sprintf(":%d", aPort), nil)
	// 이부분이 두개의 포트로 서버를 뛰울때 걑은 URL이 중복될시 문제가 된다.
	// 그래서 handler 를 만들어 수정을 한다.Í
	// rest.go, explorer.go 두파일 모두 수정을 한다.
	// handler := http.NewServeMux()
	// 고릴라 mux 로 교체
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", getBlock).Methods("GET")
	fmt.Printf("Start Server http://localhost%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", aPort), router))

}
