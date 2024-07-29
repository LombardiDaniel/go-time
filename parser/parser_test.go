package parser_test

import (
	"testing"
	"time"

	"github.com/LombardiDaniel/go-time/parser"
)

func TestParseTimeStr(t *testing.T) {
	testStr := "15:00"
	v, err := parser.ParseTimeStr(testStr)
	if err != nil {
		t.Error(err)
	}
	if *v != 15 * time.Minute {
		t.Errorf("wrong parse %s", *v)
	}

	testStr = "01:15:00"
	v, err = parser.ParseTimeStr(testStr)
	if err != nil {
		t.Error(err)
	}
	if *v != 1 * time.Hour + 15 * time.Minute {
		t.Errorf("wrong parse %s", *v)
	}

	testStr = "10"
	v, err = parser.ParseTimeStr(testStr)
	if err != nil {
		t.Error(err)
	}
	if *v != 10 * time.Second {
		t.Errorf("wrong parse %s", *v)
	}

	testStr = "aa"
	_, err = parser.ParseTimeStr(testStr)
	if err != parser.ErrNotParse {
		t.Error(err)
	}
}