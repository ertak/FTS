package System

import "time"

func Date() string {
	t := time.Now()
	return t.Format("02.01.2006")
}
