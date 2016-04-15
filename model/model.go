package Model

type User struct {
	Account UserAccount `json:"Account"`
}
type UserAccount struct {
	Username string   `json:"Username"`
	Password int      `json:"Password"`
	Bills    BillType `json:"Bills"`
}
type BillType struct {
	Electricity, Water, Gas, Phone, Other []Bill
}

type Bill struct {
	Month       string `json:"Month"`
	Amount      int    `json:"Amount"`
	DeadLine    string `json:"DeadLine"`
	Description string `json:"Description"`
	SystemInfo  Info   `json:"System_Info"`
}

type Info struct {
	CreateTime string
}


