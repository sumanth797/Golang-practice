package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://www.google.com")
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}
