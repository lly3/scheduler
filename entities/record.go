package entities

import (
	"time"

)

type Record struct {
	Id string
	ScheduleId string
	Items []RecordItem	
}

type RecordItem struct {
	Title string
	Start time.Time
	End time.Time
}
