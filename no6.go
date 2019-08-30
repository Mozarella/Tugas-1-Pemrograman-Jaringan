package main 

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprint(w,"Hello,you've requested: %s \n",r.URL.Path)
	})

	http.ListenAndServe(":8000",nil)
}