package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	callback := r.FormValue("callback")

	text, _ := json.Marshal(*r)
	txt := bytes.NewBuffer(text)
	resp, err := http.Post(callback, "text/plain", txt)
	if err != nil {
		fmt.Println("POST ERROR:", err)
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	fmt.Fprintf(w, "OK: %v", string(text))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
