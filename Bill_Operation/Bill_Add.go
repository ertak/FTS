package Bill_Operation

import (
	"FTS/System_Info"
	"FTS/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
	"net/http"
"FTS/Error"
)

var Customer Model.User

func Add_Bill(w http.ResponseWriter, r *http.Request) {

	f := fmt.Println
	frm_billtype := r.FormValue("BillTypeSelect")
	frm_sonodeme := r.FormValue("drp_sonodeme")
	frm_tutar := r.FormValue("txt_tutar")
	frm_ay := r.FormValue("dropdown_ay")
	frm_desc := r.FormValue("description")

	f("Fatura Tipi:", frm_billtype)
	f("Son Ödeme Tarihi:", frm_sonodeme)
	f("Tutar:", frm_tutar)
	f("Fatura Kesim Ayı:", frm_ay)
	f("Açıklama:",frm_desc)

	//Fatura oluşturuldu bildirimi yapılacak!

	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

		json.Unmarshal(file, &Customer)

		bill := Model.Bill{
			Month:       frm_ay,
			Amount:      frm_tutar,
			DeadLine:    frm_sonodeme,
			Description: frm_desc,
			SystemInfo:  System.Info{CreateDate: System.Date(), CreateTime: System.Time(), CreateIP: System.IPControl()},
		}

		//Fakat Error verdikten sonra kullanıcıya tekrardan giriş için hak tanınmadı direk uyarı verip program sonlandırıldı!!!!
		errValueAdd := errors.New("Aynı döneme ait fatura sistemde eklidir!")
		switch frm_billtype {
		case "Elektrik":
			counter1 := 0
			for _, valueAdd := range Customer.Account.Bills.Electricity {
				if frm_ay == valueAdd.Month {
					counter1++
					if counter1 > 0 {
						fmt.Fprintln(w,errValueAdd)
					}
				}
			}
			if counter1 == 0 {
				Customer.Account.Bills.Electricity = append(Customer.Account.Bills.Electricity, bill)

			}

		case "Doğalgaz":
			counter2 := 0
			for _, valueAdd := range Customer.Account.Bills.Gas {
				if frm_ay == valueAdd.Month {
					counter2++
					if counter2 > 0 {
						fmt.Println(errValueAdd)
						os.Exit(3)
					}
				}
			}
			if counter2 == 0 {
				Customer.Account.Bills.Gas = append(Customer.Account.Bills.Gas, bill)
			}
		case "Su":
			counter3 := 0
			for _, valueAdd := range Customer.Account.Bills.Water {
				if frm_ay == valueAdd.Month {
					counter3++
					if counter3 > 0 {
						fmt.Println(errValueAdd)
						os.Exit(3)
					}
				}
			}
			if counter3 == 0 {
				Customer.Account.Bills.Water = append(Customer.Account.Bills.Water, bill)
			}
		case "Telefon":
			counter4 := 0
			for _, valueAdd := range Customer.Account.Bills.Phone {
				if frm_ay == valueAdd.Month {
					counter4++
					if counter4 > 0 {
						fmt.Println(errValueAdd)
						os.Exit(3)
					}
				}
			}
			if counter4 == 0 {
				Customer.Account.Bills.Phone = append(Customer.Account.Bills.Phone, bill)
			}
		case "Diğer":
			counter5 := 0
			for _, valueAdd := range Customer.Account.Bills.Other {
				if frm_ay == valueAdd.Month {
					counter5++
					if counter5 > 0 {
						fmt.Fprintln(w,errValueAdd)
					}
				}
			}
			if counter5 == 0 {
				Customer.Account.Bills.Other = append(Customer.Account.Bills.Other, bill)
				billaddjs, err := json.MarshalIndent(Customer, "", " ")
				Error.WriteJsonFile(err)
				//fmt.Println(string(billaddjs))

				ioutil.WriteFile("data/Account.json", billaddjs, 0644)


				if err != nil {
					http.Error(w,err.Error(),http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type","application/json")
				w.Write(billaddjs)
			}
		}





	}

