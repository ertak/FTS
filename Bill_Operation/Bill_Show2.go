package Bill_Operation

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func Bill_Show2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("girdim-Show2")
	w.Header().Set("Content-Type", "text/html;text/css")
	//template.Must(template.ParseFiles("./web/showbill2.html")).Execute(w, nil)
	t, _ := template.ParseFiles("./web/showbill2.html")
	f := fmt.Println
	f("Fatura Tipi:", r.FormValue("ShowBillType"))
	f("Fatura Kesim Ayı:", r.FormValue("ShowMonth"))
	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)

	switch r.FormValue("ShowBillType") {
	case "Elektrik":
		for _, valueShow := range Customer.Account.Bills.Electricity {
			if r.FormValue("ShowMonth") == valueShow.Month {
				r.Form["ShowBillType2"][0] = r.FormValue("ShowBillType")
				fmt.Println("secenklerde")
				fmt.Println("Esas bak!!!:", r.Form["ShowBillType2"][0])
			}
		}
	}
	t.Execute(w, nil)
}
