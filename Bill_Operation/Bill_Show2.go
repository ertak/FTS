package Bill_Operation

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Show struct{
	Type string
	Month string
	Amount string
	Deadline string
	Desc string
}



func Bill_Show2(w http.ResponseWriter, r *http.Request) {
	//template.Must(template.ParseFiles("./web/showbill2.html")).Execute(w, nil)

	fmt.Println("girdim-Show2")

	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)
	//Parse Etme
	w.Header().Set("Content-Type", "text/html;text/css")
	var Show_Amount string
	var Show_Deadline string
	var Show_Description string


	switch r.FormValue("ShowBillType") {
	case "Elektrik":
		for _,valueShow := range Customer.Account.Bills.Electricity{
			if r.FormValue("ShowMonth") == valueShow.Month {
				Show_Amount = valueShow.Amount
				Show_Deadline = valueShow.DeadLine
				Show_Description = valueShow.Description
			}
		}
	case "Su":
		for _,valueShow := range Customer.Account.Bills.Water{
			if r.FormValue("ShowMonth") == valueShow.Month {
				Show_Amount = valueShow.Amount
				Show_Deadline = valueShow.DeadLine
				Show_Description = valueShow.Description
			}
		}
	case "Doğalgaz":
		for _,valueShow := range Customer.Account.Bills.Gas{
			if r.FormValue("ShowMonth") == valueShow.Month {
				Show_Amount = valueShow.Amount
				Show_Deadline = valueShow.DeadLine
				Show_Description = valueShow.Description
			}
		}
	case "Telefon":
		for _,valueShow := range Customer.Account.Bills.Phone{
			if r.FormValue("ShowMonth") == valueShow.Month {
				Show_Amount = valueShow.Amount
				Show_Deadline = valueShow.DeadLine
				Show_Description = valueShow.Description
			}
		}
	case "Diğer":
		for _,valueShow := range Customer.Account.Bills.Other{
			if r.FormValue("ShowMonth") == valueShow.Month {
				Show_Amount = valueShow.Amount
				Show_Deadline = valueShow.DeadLine
				Show_Description = valueShow.Description
			}
		}
	}

	BillInfo:=Show{Type:r.FormValue("ShowBillType"),Month:r.FormValue("ShowMonth"),Amount:Show_Amount,Deadline:Show_Deadline,Desc:Show_Description}

	t, _ := template.ParseFiles("./web/showbill2.html")

	t.Execute(w, BillInfo)
}

