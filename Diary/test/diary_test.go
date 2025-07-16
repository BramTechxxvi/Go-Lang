package test

import (
	"diary/src"
	"testing"
)

func TestCanUnlockDiary__UnlockDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	diary.UnlockDiary("password")
	if diary.IsLocked() {
		t.Error("Incorrect password")
	}
	if diary.UnlockDiary("pass") {
		t.Error("Incorrect password")
	}
}

func TestCanLockDiary__LockDiary(t *testing.T) {
	diary := src.NewDiary("Grades", "password")
	unlocked := diary.UnlockDiary("password")
	if !unlocked {
		t.Error("Incorrect password")
	}
	lockedDiary := diary.lockDiary()
	if !lockedDiary {
		t.Error("Error")
	}

}
