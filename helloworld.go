package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client:=&http.Client{}
	req, _ := http.NewRequest("GET", "http://ip.qw0.cc", nil)
	req.Header.Set("User-Agent","curl/7.64.1")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("请求失败: " + err.Error())
	}
	respBody,_:=ioutil.ReadAll(resp.Body)
	target := string(respBody)
	w.Header().Set("Content-Type","application/json")
	fmt.Fprintf(w, "%s!\n", target)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8012"
	}
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

