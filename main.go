package main

import (
	"FTS/Bill_Operation"
	"FTS/Screen_Question"
)

func main() {

	Screen_Question.Question()
	if Screen_Question.Bill_opr_selected == 1 {
		Bill_Operation.Add_Bill()
	}else {
		Bill_Operation.Delete_Bill()
	}


}


