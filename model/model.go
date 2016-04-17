package Model

import "FTS/System_Info"

type User struct { // Her kullanıcı genel bir fatura hesabna sahiptir.
	Account UserAccount `json:"Account"`
}
type UserAccount struct { //Hesabta tanımlı kullanıcı adı,şifre ve faturalar vardır.
	Username string   `json:"Username"`
	Password int      `json:"Password"`
	Bills    BillType `json:"Bills"`
}
type BillType struct { //Çeşitli fatura mevcuttur.
	Electricity, Water, Gas, Phone, Other []Bill
}

type Bill struct { //Fatura Bilgisi
	Month       string      `json:"Month"`
	Amount      int         `json:"Amount"`
	DeadLine    string      `json:"DeadLine"`
	Description string      `json:"Description"`
	SystemInfo  System.Info `json:"System_Info"` //Sisteme ait olan otomatik bilgidir.ew
}
