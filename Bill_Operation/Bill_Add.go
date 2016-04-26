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
	err := errors.New("Belirtilen döneme ait fatura zaten eklenmiştir!")

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

	//Error kontrolü yapılmalı bu döneme ait fatura zaten oluşturuldu gibi!!!
	switch Screen_Question.Billtype {
	case 1:
		//Eklenirken aynı döneme ait 2. fatura ekleme kontrolü yapılacak
		Customer.Account.Bills.Electricity = append(Customer.Account.Bills.Electricity, bill)
	case 2:
		Customer.Account.Bills.Gas = append(Customer.Account.Bills.Gas, bill)
	case 3:
		Customer.Account.Bills.Water = append(Customer.Account.Bills.Water, bill)
	case 4:
		Customer.Account.Bills.Phone = append(Customer.Account.Bills.Phone, bill)
	case 5:
		Customer.Account.Bills.Other = append(Customer.Account.Bills.Other, bill)
	}

	billaddjs, err := json.MarshalIndent(Customer, "", " ")
	Error.WriteJsonFile(err)
	fmt.Println(string(billaddjs))

	ioutil.WriteFile("data/Account.json", billaddjs, 0644)
}
