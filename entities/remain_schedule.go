package entities

import (
	"fmt"
	"time"
)

type RemainSchedule struct {
	RemainItems []RemainItem
}

type RemainItem struct {
	Title string
	Remain time.Duration
}

func (r RemainSchedule) String() string {
	s := ""

	for _, v := range r.RemainItems {
		s += fmt.Sprintln(v.Title, v.Remain)
	}

	return s
}
