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

func TestCanDeleteEntry__DeleteEntry(t *testing.T) {
	diary := src.NewDiary("Books", "password")
	diary.UnlockDiary("password")
	diary.CreateEntry("Maths", "Addition, subtraction, division")
	diary.CreateEntry("English", "Figures of speech, essay")

	if len(diary.Entries()) != 2 {
		t.Errorf("Expected 2 entries got %d", len(diary.Entries()))
	}
	var entry int
	found := false
	for _, value := range diary.Entries() {
		if value.GetTitle() == "English" {
			entry = value.GetID()
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Entry not found")
	}

	diary.Delete(entry)
	updatedEntries := diary.Entries()
	if len(updatedEntries) != 1 {
		t.Errorf("Expected 1 entries got %d", len(updatedEntries))
	}
}

func TestCanUpdateEntry__UpdateEntry(t *testing.T) {
	diary := src.NewDiary("Books", "password")
	diary.UnlockDiary("password")
	diary.CreateEntry("Maths", "Addition, subtraction, division")
	diary.CreateEntry("English", "Figures of speech, essay")

	entries := diary.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries got %d", len(entries))
	}

	var entry src.Entry
	found := false
	for _, value := range entries {
		if value.GetTitle() == "English" {

			found := true
			break
		}
	}
}
