package Bill_Operation
import (
	"net/http"
	"fmt"
	"html/template"
)


func Deneme (w http.ResponseWriter, r *http.Request)  {
	fmt.Println("girdim deneme")

	w.Header().Set("Content-Type","text/html;text/css")

	t,_ := template.ParseFiles("./web/index2.html")
	t.Execute(w,nil)
}
