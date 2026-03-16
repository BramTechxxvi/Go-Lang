package src

import "time"

type Entry struct {
	iD          int
	title       string
	body        string
	dateCreated time.Time
}

func NewEntry(id int, title string, body string) Entry {
	return Entry{
		iD:          id,
		title:       title,
		body:        body,
		dateCreated: time.Now(),
	}
}

func (e *Entry) GetID() int {
	return e.iD
}

func (e *Entry) setID(id int) {
	e.iD = id
}

func (e *Entry) GetTitle() string {
	return e.title
}

func (e *Entry) SetTitle(title string) {
	e.title = title
}

func (e *Entry) GetBody() string {
	return e.body
}

func (e *Entry) SetBody(body string) {
	e.body = body
}

func (e *Entry) GetDateCreated() time.Time {
	return e.dateCreated
}

func (e *Entry) SetDateCreated(date time.Time) {
	e.dateCreated = date
}
