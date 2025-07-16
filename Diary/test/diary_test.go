package test

import (
	"diary/src"
	"testing"
)

func TestCanUnlockDiary__UnlockDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	diary.UnlockDiary("pass")
	if !diary.IsLocked() {
		t.Error("Incorrect password")
	}
	diary.UnlockDiary("password")
	if diary.IsLocked() {
		t.Error("Incorrect password")
	}
}

func TestCanLockDiary__LockDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	diary.UnlockDiary("password")
	if diary.IsLocked() {
		t.Error("Incorrect password")
	}
	diary.LockDiary()
	if !diary.IsLocked() {
		t.Error("Error")
	}
}

func TestCanCreateEntry__CreateEntry(t *testing.T) {
	diary := src.NewDiary("Books", "password")
	diary.UnlockDiary("password")
	diary.CreateEntry("Maths", "Addition, subtraction, division")
	diary.CreateEntry("English", "Figures of speech, essay")

	entries := diary.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries got %d", len(entries))
	}
}

func TestCanDeleteEntry__DeleteEntry2(t *testing.T) {
	diary := src.NewDiary("Books", "password")
	diary.UnlockDiary("password")
	diary.CreateEntry("Maths", "Addition, subtraction, division")
	diary.CreateEntry("English", "Figures of speech, essay")

	entries := diary.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries got %d", len(entries))
	}

}
