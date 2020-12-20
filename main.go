package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func sayHello(responseWriter http.ResponseWriter, request *http.Request) {
	//request.ParseForm()
	fmt.Println(request.Method)
	request.FormValue("12")
	fmt.Println(request.Form)
	request.FormValue("12")

	_, _ = responseWriter.Write([]byte("index"))

}

func index(responseWriter http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadFile("/root/go/http_server/index.html")
	if err != nil {

	}
	num := rand.Intn(10)
	str := string(data)
	if num > 5 {
		str = strings.Replace(str, "{ooxx}", "<li>my world</li>", 1)
	} else {
		str = strings.Replace(str, "{ooxx}", "<li>my dog</li>", 1)
	}
	_, _ = fmt.Fprint(responseWriter, str)
	//_, _ = responseWriter.Write([]byte("1234567890"))
}

func templateFunc(responseWriter http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("/root/go/http_server/index.html")
	data := "<li>my world</li>"
	t.Execute(responseWriter, data)

}
func main() {
	log.SetPrefix("")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds | log.Llongfile | log.Lmsgprefix | log.Flags())
	log.Print(os.Getwd())

	http.HandleFunc("/web/", sayHello)
	http.HandleFunc("/web", sayHello)
	http.HandleFunc("/index", index)
	http.HandleFunc("/template/", templateFunc)
	a := http.DefaultServeMux
	fmt.Println(a)
	panic(http.ListenAndServe(":80", nil))

}
