package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"FTS/Error"
	"FTS/model"
)

func main() {
	file, e := ioutil.ReadFile("../data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var customer Model.User
	json.Unmarshal(file, &customer)

	fmt.Println("Fatura İşlemleri")
	fmt.Println("Elektrik Faturası	: (1)")
	fmt.Println("Doğalgaz Faturası	: (2)")
	fmt.Println("Su Faturası	: (3)")
	fmt.Println("Telefon Faturası	: (4)")
	fmt.Println("Diğer Faturası	: (5)")

	var billtype int
	fmt.Scanf("%d\n",&billtype)


	fmt.Println("Fatura Ayı:")
	var billMonth string
	fmt.Scanf("%s\n",&billMonth)

	fmt.Println("Tutar:")
	var billAmount int
	fmt.Scanf("%d\n",&billAmount)

	fmt.Println("Son Ödeme Tarihi:")
	var billDate string
	fmt.Scanf("%s\n",&billDate)

	fmt.Println("Açıklama:")
	var billDesc string
	fmt.Scanf("%s\n",&billDesc)

	bill := Model.Bill{Month:billMonth,Amount:billAmount,DeadLine:billDate,Description:billDesc,SystemInfo:Model.Info{CreateTime:"10.11.12"}}

	switch billtype {
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


