package main

import (
	"FTS/Bill_Operation"
	"log"
	"net/http"

)

func main(){

	http.Handle("/",http.FileServer(http.Dir("./web")))
	http.HandleFunc("/addbill", Bill_Operation.Add_Bill)
	http.HandleFunc("/delbill", Bill_Operation.Delete_Bill())
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("ListenAndServe" ,err)
	}

}


