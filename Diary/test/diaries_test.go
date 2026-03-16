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

func TestCanFindDiaryByUsername__FindDiaryByUsername(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	diary2 := src.NewDiary("Books", "password")
	diaries := src.NewDiaries()
	diaries.AddNew(diary)
	diaries.AddNew(diary2)

	var foundDiary = diaries.FindByUsername("Books")

	if foundDiary == nil {
		t.Errorf("Couldn't find diary by username")
	}
}
