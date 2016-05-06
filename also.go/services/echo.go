package main 

import (
	"fmt"; "log"; "net/http"; "os"; "io/ioutil"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Println(os.Args[0])
}

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, readFile(r.URL.Path))
	data, err := ioutil.ReadFile("/Users/admin/GoWorkspace/www" + r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	} else {
		fmt.Fprintf(w, "%s", data)	
	}
}

//func readFile(path string) {
//	f, err := os.OpenFile(path)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "echo: %v\n", err) 
//	}
//	f.Read
//}