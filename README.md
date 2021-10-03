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

## Package Sync

> Once는 정확히 단 한번만 수행하는 객체이다

```go
// channel을 통한 방법 (sync.Once 사용 전)
// 목적은 초기화는 한 번만 수행하고, 모든 분산 처리 고루틴이 이 초기화를 하고 나서 수행되어야 하는 경우!
func main() {
    done := make(chan struct{})
    go func() { // 초기화 진행용 고루틴! 한번만 수행
        defer close(done)   // 고루틴 종료 전, 채널로 데이터(=시그널) 전송
        fmt.Println("Init") // 초기화 진행
    }()

    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)        // WaitGroup에 고루틴 등록
        go func(i int) { // 분산 처리 고루틴, 초기화 진행용 고루틴 이후에 진행
            defer wg.Done()               // 고루틴 종료 전, 고루틴 종료 시그널 발생
            <-done                        // 채널로부터 데이터(=시그널) 수신
            fmt.Println("Goroutine: ", i) // 개별 고루틴 진행
        }(i)
    }
    wg.Wait() // 모든 고루틴 종료될 때까지 대기(blocking)
}

// 초기화 진행용 고루틴에서는 초기화를 단한번만 수행하도록 한다
// 그러기 위해서 먼저 초기화 진행 후 미리 만들어놓은 channel(done)을 통해 데이터(=시그널, close)을 전송한다
// channel(done)을 통해 수신 받은(=시그널, close) 고루틴에서는 그(초기화) 이후에 작업을 수행한다

==================================================================================
// sync.Once 통한 방법
// 목적은 초기화는 한 번만 수행하고, 모든 분산 처리 고루틴이 이 초기화를 하고 나서 수행되어야 하는 경우!
// sync.Once을 활용한 방법으로 전체소스를 보면,
func main() {
    var once sync.Once    // Once 객체 생성
    var wg sync.WaitGroup // WaitGroup 객체 생성
    for i := 0; i < 3; i++ {
        wg.Add(1)        // 고루틴 등록
        go func(i int) { // 개별 고루틴 생성 및 수행
            defer wg.Done()  // 고루틴 종료 시그널
            once.Do(func() { // 한번만 실행하게끔 Do() 메서드 수행
                fmt.Println("Init") // 한번만 실행될 내용
            })
            fmt.Println("Goroutine: ", i) // 개별 고루틴 수행
        }(i)
    }
    wg.Wait() // 모든 고루틴 종료까지 Wait
}
// 목적은 위 channel 방식과 같다
// 중요한것은 once.Do() 내용이다
// 보면 개별 고루틴이 3개 수행되는데, 개별 고루틴 안에 once.Do(...)내용이 들어가 있다
// 하지만 몇개의 고루틴이 수행되던간에 once.Do(...) 한번만 수행된다는 것이다
```

---

## Css를 쉽게 적용하기 위해서 MVP.css 사용 (외국애들은 참 무료로 잘 풀어 놓는것 같다 ㅎ)

[MVP.css Link](https://andybrewer.github.io/mvp/)

---

## Go에서 UI(html) 연결하기

  1. html 파일을 직접연결 하여 사용
      * [잘 정리된 블로그 LINK](https://dksshddl.tistory.com/entry/Go-web-programming-%ED%85%9C%ED%94%8C%EB%A6%BF%EA%B3%BC-%ED%85%9C%ED%94%8C%EB%A6%BF-%EC%97%94%EC%A7%84)

  2. Header, Content, Footer 재사용을위한 설정
      * Partials 이라는 것이 있다.
            - [Golang Web - Render Partial HTML Template Link Click](https://dev.to/egaprsty/golang-web-render-partial-html-template-3h1m)

---

## struct field tag

* Go lang 에서는 대문자 시작은 자동 Export 소문자 시작은 Export가 되지 않는다.
  그런데, struct 필드를 소문자로 구성하고 Export를 하고 싶을때가 있을것이다.
  이런 상황에 쓰이는 것이 struct field tag 이다.

  ```go
  type User struct {
    Name string `json:"name"`
    Password string `json:"password"`
  }
  // 이렇게 작성하면 Export가 되면서 외부에서는 대문자가 아닌 소문자로 보여지게 된다.
  ```

* omitempty 비어있는값은 숨김처리 해주는 기능
  
  ```go
    User struct {
    Name string `json:"name"`
    Password string `json:"password"`
    Test string `json:"test, omitempty"`
  }
  ```

  자세한 내용은 아래 링크에서 확인   

  [GO Lang 공식 문서 링크](https://pkg.go.dev/encoding/json#Marshal)

  [정리된 누군가의 블로그](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)

---

### VSCode Extention 소개  잠시

> rest client Extention
> vscode 에서 작업을 하면서 API 테스트를 도와준다.
> 프로젝트 내에 ***.http 파일을 생성
> 해당파일로 API 테스트 코드를 작성
> 사용방법은 검색을 통해 해보면 아주 좋은 Extention 이라는 것을 알수 있다.

---

### Multiplexer

[링크로가서 잘 보세요](https://dejavuqa.tistory.com/314)

### Gorilla mux

> Go 기본 패키지에 없는 다양한 기능들을 가지고 있는 Lib 이다.
> 차근차근 Study를 해봐야 한다.. 할게 많다..
> 사용을 위해서 go get -u github.com/gorilla/mux 명령어로 설치가 가능하다.

* [Gorilla mux Git hub](https://github.com/gorilla/mux)
* [Gorilla mux Web](https://www.gorillatoolkit.org)
* adapter 패턴도 함께 더 봐야할 부분이다.

### Go Lang Type Converter

strconv.Atoi(**String**)

