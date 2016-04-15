package Bill_Operation
import
(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"FTS/Error"
	"FTS/model"
	"FTS/Screen_Question"
)

func Add_Bill()  {

	file, e := ioutil.ReadFile("../data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var customer Model.User
	json.Unmarshal(file, &customer)

	bill := Model.Bill{Month:Screen_Question.BillMonth,Amount:Screen_Question.BillAmount,DeadLine:Screen_Question.BillDate,Description:Screen_Question.BillDesc,SystemInfo:Model.Info{CreateTime:"10.11.12"}}

	switch Screen_Question.Billtype {
	case 1:
		customer.Account.Bills.Electricity = append(customer.Account.Bills.Electricity,bill)
	case 2:
		customer.Account.Bills.Gas = append(customer.Account.Bills.Gas,bill)
	case 3:
		customer.Account.Bills.Water = append(customer.Account.Bills.Water,bill)
	case 4:
		customer.Account.Bills.Phone = append(customer.Account.Bills.Phone,bill)
	case 5:
		customer.Account.Bills.Other = append(customer.Account.Bills.Other,bill)
	}

	billjs,err:= json.Marshal(customer)
	Error.WriteJsonFile(err)
	fmt.Println(string(billjs))

	ioutil.WriteFile("../data/Account.json",billjs,0644)
}
