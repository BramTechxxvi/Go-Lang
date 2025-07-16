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

func (diary *Diary) UnlockDiary(password string) {
	if diary.password == password {
		diary.isLocked = false
	}
}

func (diary *Diary) LockDiary() {
	diary.isLocked = true
}

func (diary *Diary) IsLocked() bool {
	return diary.isLocked()
}
