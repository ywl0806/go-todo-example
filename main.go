package main

import (
	"net/http"
)

func main() {
	println("hoge")
	handler := new(testHandler)
	http.Handle("/", handler)

	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

// testHandler 객체의 method
func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}
