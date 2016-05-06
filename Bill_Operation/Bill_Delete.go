package Bill_Operation
import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"errors"
	"FTS/Error"
	"net/http"
)


func Delete_Bill(w http.ResponseWriter, r *http.Request)  {

	f := fmt.Println
	frm_billtype := r.FormValue("BillTypeSelect")
	frm_ay := r.FormValue("dropdown_ay")

	f("Fatura Tipi:", frm_billtype)
	f("Fatura Kesim Ayı:", frm_ay)

	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)

	errValueDelete := errors.New("Belirtilen döneme ait fatura bulunamadı!")
	switch frm_billtype {
	case "Elektrik":
		for i,value := range Customer.Account.Bills.Electricity{
			if  frm_ay == value.Month {
				Customer.Account.Bills.Electricity = append(Customer.Account.Bills.Electricity[:i], Customer.Account.Bills.Electricity[i+1:]...)
			}else {
				fmt.Fprintln(w,errValueDelete)
			}
		}
	case "Doğalgaz":
		for i,value := range Customer.Account.Bills.Gas{
			if  frm_ay == value.Month {
				Customer.Account.Bills.Gas = append(Customer.Account.Bills.Gas[:i], Customer.Account.Bills.Gas[i+1:]...)
			}else {
				fmt.Fprintln(w,errValueDelete)
			}
		}
	case "Su":
		for i,value := range Customer.Account.Bills.Water {
			if frm_ay == value.Month {
				Customer.Account.Bills.Water = append(Customer.Account.Bills.Water[:i], Customer.Account.Bills.Water[i + 1:]...)
			}else {
				fmt.Fprintln(w, errValueDelete)
			}
		}

	case "Telefon":
		for i,value := range Customer.Account.Bills.Phone{
			if  frm_ay == value.Month {
				Customer.Account.Bills.Phone = append(Customer.Account.Bills.Phone[:i], Customer.Account.Bills.Phone[i+1:]...)
			}else {
				fmt.Fprintln(w,errValueDelete)
			}
		}
	case "Diğer":
		for i,value := range Customer.Account.Bills.Other{
			if  frm_ay == value.Month {
				Customer.Account.Bills.Other = append(Customer.Account.Bills.Other[:i], Customer.Account.Bills.Other[i+1:]...)
			}else {
				fmt.Fprintln(w,errValueDelete)
			}

		}
	}

		billdeljs, err := json.MarshalIndent(Customer, "", " ")
		Error.WriteJsonFile(err)
		fmt.Println(string(billdeljs))

		ioutil.WriteFile("data/Account.json", billdeljs, 0644)

		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type","application/json")
		w.Write(billdeljs)
}
