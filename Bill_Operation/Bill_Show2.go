package Bill_Operation
import (
	"net/http"
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"html/template"
)

var tpl *template.Template
func init(){
	tpl=template.Must(template.ParseFiles("web/showbill2.html"))
}


func Bill_Show2(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("girdim-Show2")
	tpl.Execute(w,nil)
	file, e := ioutil.ReadFile("data/Account.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	json.Unmarshal(file, &Customer)

	switch show_billtype {
	case "Elektrik":
		for _,valueShow := range Customer.Account.Bills.Electricity{
			if show_billMonth == valueShow.Month{
				r.Form["ShowBillType2"][0] = show_billtype
				fmt.Println("Esas bak!!!:",r.Form["ShowBillType2"][0])
			}
		}
	}
}
