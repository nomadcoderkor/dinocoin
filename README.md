# GO Languege로 시작하는 Blockchain
1. GO 프로젝트 생성
   * go mod init git url / 레파지토리 명    
      ex) go mod init github.com/nomadcoderkor/dinocoin
2. 위 과정을 실행 하면 go.mod 파일이 생성된다. go.mod파일은 node.js의 package.json과 비슷한 파일이라 생각하면 된다.
3. root에 main.go 파일 작성 -> 터미널에서 go run main.go 실행

실행이 잘 된다면, go 프로젝트 셋팅이 완료 된것 입니다.   

---

## 개발 중간중간 정리를 하면서 넘어간다.
> Go는 class를 갖고 있지 않습니다. 그러나 타입과 위에 method를 정의할 수 있습니다.
> method는 특별한 receiver arguments를 갖고 있는 함수입니다.
> receiver는 자신의 arguments를 func와 method keyword 사이에 작성된다.
```go
type blockchain struct {
  blocks []block
}
func (b blockchain) addBlock(data string){
  fmt.Println(b)
}
```
* 위에 보이는것과 같이 함수명 앞에 타입과 변수명이 붙은것을 Receiver라고 부른다.   
  리시버가 붙으면 더이상 함수가 아니라 메소드가 된다. 
* 위 예제는 addBlock 함수가 아니라 blockchain이 갖게 되는 메소드가 된것이다.
  다른곳에서 __blockchain.addBlock__와 같은 형태로 쓰이게 된다.



