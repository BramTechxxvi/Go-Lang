package src

type Diary struct {
	username string
	password string
	isLocked bool
	entries  []Entry
}

func NewDiary(username string, password string) Diary {
	return Diary{
		username: username,
		password: password,
		isLocked: true,
		entries:  make([]Entry, 0),
	}
}

func unlockDiary(password string) {

}
