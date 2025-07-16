package src

type Diary struct {
	username string
	password string
	isLocked bool
	entries  []Entry
}

func newDiary(username string, password string) Diary {
	return Diary{
		username: username,
		password: password,
		isLocked: true,
		entries:  make([]Entry, 0),
	}
}
