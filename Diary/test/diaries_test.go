package test

import (
	"diary/src"
	"testing"
)

func TestCanAddDiary__AddDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	diaries := src.NewDiaries()

	diaries.AddNew(diary)
	if diaries.Count() != 1 {
		t.Errorf("Count should be 1: got  %d", diaries.Count())
	}
}

func TestCanRemoveDiary__deleteDiary(t *testing.T) {
	firstDiary := src.NewDiary("Grades", "password")
	secondDiary := src.NewDiary("Groceries", "password")
	diaries := src.NewDiaries()
	diaries.AddNew(firstDiary)
	diaries.AddNew(secondDiary)

	if diaries.Count() != 2 {
		t.Errorf("Count should be 2: got  %d", diaries.Count())
	}

	diaries.Delete(secondDiary)
	if diaries.Count() > 1 {
		t.Errorf("Count should be 1: got  %d", diaries.Count())
	}

	if diaries.Count() < 1 {
		t.Errorf("Count should be 1: got  %d", diaries.Count())
	}

}
