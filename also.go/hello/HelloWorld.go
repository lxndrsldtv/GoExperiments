package main

import (
	"fmt"; "net/http"; "io/ioutil"
)


func main() {
	fmt.Println("Hello Go World")
	http.HandleFunc("/service/rest", requestHandler)
	http.ListenAndServe("localhost:9090", nil)
}

func requestHandler (w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile("/Users/admin/IdeaProjects/HelloGo/www/index.html")
	fmt.Fprintf(w, "%s", data)
}

