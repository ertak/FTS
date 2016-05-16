package Bill_Operation
import (
	"net/http"
	"fmt"
	"html/template"
)

type ShowResult struct  {
	rType string
	rMonth string
	rAmount string
	rDeadline string
	rDesc string

}

func Add_ShowResult(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("girdim Add_ShowResult")

	w.Header().Set("Content-Type","text/html;text/css")

	BillResult:= ShowResult{rType:r.FormValue("BillTypeSelect"),rMonth:r.FormValue("dropdown_ay"),rAmount:r.FormValue("txt_tutar"),rDeadline:r.FormValue("drp_sonodeme"),rDesc:r.FormValue("description")}

	t, _ := template.ParseFiles("./web/showResult.html")

	t.Execute(w,BillResult)
}