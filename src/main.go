package main

import(
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "These coding changes are only visible in the staging area. Hello World!")
	})
}

func main(){
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}
