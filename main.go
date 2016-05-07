package main

import (
	"FTS/Bill_Operation"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.Handle("css/index.css", http.StripPrefix("css/", http.FileServer(http.Dir("./web/css/index.css"))))
	http.HandleFunc("/addbill", Bill_Operation.Add_Bill)
	http.HandleFunc("/delbill", Bill_Operation.Delete_Bill)
	http.HandleFunc("/showbill", Bill_Operation.Bill_Show)
	http.HandleFunc("/showbill2", Bill_Operation.Bill_Show2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}
