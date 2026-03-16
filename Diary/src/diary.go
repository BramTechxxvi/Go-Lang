package src

var counter int

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
	return diary.isLocked
}

func (diary *Diary) CreateEntry(title string, description string) {
	if diary.isLocked == false {
		counter++
		entryId := counter
		entry := NewEntry(entryId, title, description)
		diary.entries = append(diary.entries, entry)
	}

}

func (diary *Diary) Entries() []Entry {
	return diary.entries
}

func (diary *Diary) Delete(entryId int) {
	for i, entry := range diary.entries {
		if entry.GetID() == entryId {
			diary.entries = append(diary.entries[:i], diary.entries[1+1:]...)
			break
		}
	}
}

func (diary *Diary) UpdateEntry(entryToUpdate Entry) {
	for i, existing := range diary.entries {
		if existing.GetID() == entryToUpdate.GetID() {
			if entryToUpdate.title != "" {
				diary.entries[i].SetTitle(entryToUpdate.GetTitle())
			}
			if entryToUpdate.body != "" {
				diary.entries[i].SetBody(entryToUpdate.GetBody())
			}
			break
		}
	}
}
