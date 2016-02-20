package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		t.ExecuteTemplate(w, "T", template.HTML("<script>alert('blah!')</script>"))
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("LsitenAndServe: ", err)
	}
}

// import (
// 	"fmt"
// 	"net/http"
// )

// type MyMux struct{}

// func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		sayHelloName(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// 	return
// }

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello myroute!")
// }

// func main() {
// 	mux := &MyMux{}
// 	http.ListenAndServe(":9090", mux)
// }
// --------------------------------------------------------------
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// func sayHelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key: ", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!")
// }

// func main() {
// 	http.HandleFunc("/", sayHelloName)
// 	err := http.ListenAndServe(":9090", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }
