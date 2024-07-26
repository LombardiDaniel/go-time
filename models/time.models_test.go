package models_test

import (
	"testing"

	"github.com/LombardiDaniel/go-time/models"
)

func TestFormatNum(t *testing.T) {
	var got string
	var expected string

	got = models.FormatNum(1, 60)
	expected = "01"
	if got != expected {
		t.Errorf("err on %s: got %s", expected, got)
	}

	got = models.FormatNum(50, 60)
	expected = "50"
	if got != expected {
		t.Errorf("err on %s: got %s", expected, got)
	}
}