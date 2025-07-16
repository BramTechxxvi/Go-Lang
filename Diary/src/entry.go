package src

import "time"

type Entry struct {
	iD          int
	title       string
	body        string
	dateCreated time.Time
}

func newEntry(id int, title string, body string) Entry {
	return Entry{
		iD:          id,
		title:       title,
		body:        body,
		dateCreated: time.Now(),
	}
}
