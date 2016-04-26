package Bill_Operation
import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"FTS/Screen_Question"
	"errors"
	"FTS/Error"
)


func Delete_Bill()  {


	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)

	//Error kontrolü yapıldı o dönemin faturası olmadığında hata veriyor fakat kullanıcıya tekrardan hak tanınması da yapılabilinir daha sonra!
	err := errors.New("Belirtilen döneme ait fatura bulunamadı!")
	switch Screen_Question.Billtype {
	case 1:
		for i,value := range Customer.Account.Bills.Electricity{
			if  Screen_Question.BillMonth == value.Month {
				Customer.Account.Bills.Electricity = append(Customer.Account.Bills.Electricity[:i], Customer.Account.Bills.Electricity[i+1:]...)
			}else {
				fmt.Println(err)
			}
		}
	case 2:
		for i,value := range Customer.Account.Bills.Gas{
			if  Screen_Question.BillMonth== value.Month {
				Customer.Account.Bills.Gas = append(Customer.Account.Bills.Gas[:i], Customer.Account.Bills.Gas[i+1:]...)
			}else {
				fmt.Println(err)
			}
		}
	case 3:
		for i,value := range Customer.Account.Bills.Water{
			if  Screen_Question.BillMonth== value.Month {
				Customer.Account.Bills.Water = append(Customer.Account.Bills.Water[:i], Customer.Account.Bills.Water[i+1:]...)
			}else {
				fmt.Println(err)
			}
		}

	case 4:
		for i,value := range Customer.Account.Bills.Phone{
			if  Screen_Question.BillMonth== value.Month {
				Customer.Account.Bills.Phone = append(Customer.Account.Bills.Phone[:i], Customer.Account.Bills.Phone[i+1:]...)
			}else {
				fmt.Println(err)
			}
		}
	case 5:
		for i,value := range Customer.Account.Bills.Other{
			if  Screen_Question.BillMonth == value.Month {
				Customer.Account.Bills.Other = append(Customer.Account.Bills.Other[:i], Customer.Account.Bills.Other[i+1:]...)
			}else {
				fmt.Println(err)
			}

		}
	}


	billdeljs, err := json.MarshalIndent(Customer, "", " ")
	Error.WriteJsonFile(err)
	fmt.Println(string(billdeljs))

	ioutil.WriteFile("data/Account.json", billdeljs, 0644)
}
