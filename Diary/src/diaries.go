package src

type Diaries struct {
	diaries []Diary
}

func NewDiaries() *Diaries {
	return &Diaries{diaries: []Diary{}}
}

func (d *Diaries) AddNew(diary Diary) {
	d.diaries = append(d.diaries, diary)
}

func (d *Diaries) Delete(diary Diary) {
	for i, values := range d.diaries {
		if values.username == diary.username {
			d.diaries = append(d.diaries[:i], d.diaries[i+1:]...)
		}
	}
}

func (d *Diaries) Count() int {
	return len(d.diaries)
}
