package System

import "time"

func Time() string {
	t := time.Now()
	return t.Format("15:04:05")
}
