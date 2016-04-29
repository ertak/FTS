package Bill_Operation

import (
	"FTS/Error"
	"FTS/Screen_Question"
	"FTS/System_Info"
	"FTS/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
)

var Customer Model.User

func Add_Bill() {

	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)

	bill := Model.Bill{
		Month:       Screen_Question.BillMonth,
		Amount:      Screen_Question.BillAmount,
		DeadLine:    Screen_Question.BillDate,
		Description: Screen_Question.BillDesc,
		SystemInfo:  System.Info{CreateDate: System.Date(), CreateTime: System.Time(), CreateIP: System.IPControl()},
	}

	//Eklenirken aynı döneme ait 2. fatura ekleme kontrolü yapıldı!
	//Fakat Error verdikten sonra kullanıcıya tekrardan giriş için hak tanınmadı direk uyarı verip program sonlandırıldı!!!!
	errValueAdd := errors.New("Aynı döneme ait fatura sistemde eklidir!")
	switch Screen_Question.Billtype {
	case 1:
		for _,valueAdd := range Customer.Account.Bills.Electricity {
			if Screen_Question.BillMonth == valueAdd.Month {
				fmt.Println(errValueAdd)
				os.Exit(3)
			}else {
				Customer.Account.Bills.Electricity = append(Customer.Account.Bills.Electricity, bill)
			}
		}

	case 2:
		for _,valueAdd := range Customer.Account.Bills.Gas {
			if Screen_Question.BillMonth == valueAdd.Month {
				fmt.Println(errValueAdd)
				os.Exit(3)
			}else {
				Customer.Account.Bills.Gas = append(Customer.Account.Bills.Gas, bill)
			}
		}
	case 3:
		for _,valueAdd := range Customer.Account.Bills.Water {
			if Screen_Question.BillMonth == valueAdd.Month {
				fmt.Println(errValueAdd)
				os.Exit(3)
			}else{
				Customer.Account.Bills.Water = append(Customer.Account.Bills.Water, bill)
			}
		}
	case 4:
		for _,valueAdd := range Customer.Account.Bills.Phone {
			if Screen_Question.BillMonth == valueAdd.Month {
				fmt.Println(errValueAdd)
				os.Exit(3)
			}else{
				Customer.Account.Bills.Phone = append(Customer.Account.Bills.Phone, bill)
			}
		}
	case 5:
		for _,valueAdd := range Customer.Account.Bills.Other {
			if Screen_Question.BillMonth == valueAdd.Month {
				fmt.Println(errValueAdd)
				os.Exit(3)
			}else{
				Customer.Account.Bills.Other = append(Customer.Account.Bills.Other, bill)
			}
		}
	}
	billaddjs, err := json.MarshalIndent(Customer, "", " ")
	Error.WriteJsonFile(err)
	fmt.Println(string(billaddjs))

	ioutil.WriteFile("data/Account.json", billaddjs, 0644)
}
