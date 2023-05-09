package entities

import (
	"fmt"
	"time"
)

type Schedule struct {
	Id string
	Todos []ScheduleItem
}

type ScheduleItem struct {
	Title string
	Duration time.Duration
}

func (s Schedule) String() string {
	str := fmt.Sprintln(s.Id)

	for _, v := range s.Todos {
		str += fmt.Sprintln(v.Title, v.Duration)
	}

	return str
}
