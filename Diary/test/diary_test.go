package test

import (
	"diary/src"
	"testing"
)

func testCanUnlockDiary__UnlockDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	unlocked := diary.UnlockDiary("password")
	if unlocked {
		t.Error("Incorrect password")
	}
	unlocked = diary.UnlockDiary("pass")
	if !unlocked {
		t.Error("Incorrect password")
	}
}
