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
)

func Add_Bill() {

	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var customer Model.User
	json.Unmarshal(file, &customer)

	bill := Model.Bill{
		Month:       Screen_Question.BillMonth,
		Amount:      Screen_Question.BillAmount,
		DeadLine:    Screen_Question.BillDate,
		Description: Screen_Question.BillDesc,
		SystemInfo:  System.Info{CreateDate: System.Date(), CreateTime: System.Time(), CreateIP: System.IPControl()},
	}

	switch Screen_Question.Billtype {
	case 1:
		customer.Account.Bills.Electricity = append(customer.Account.Bills.Electricity, bill)
	case 2:
		customer.Account.Bills.Gas = append(customer.Account.Bills.Gas, bill)
	case 3:
		customer.Account.Bills.Water = append(customer.Account.Bills.Water, bill)
	case 4:
		customer.Account.Bills.Phone = append(customer.Account.Bills.Phone, bill)
	case 5:
		customer.Account.Bills.Other = append(customer.Account.Bills.Other, bill)
	}

	billjs, err := json.MarshalIndent(customer, "", " ")
	Error.WriteJsonFile(err)
	fmt.Println(string(billjs))

	ioutil.WriteFile("data/Account.json", billjs, 0644)
}
