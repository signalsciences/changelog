package changelog

import (
	"reflect"
	"testing"
)

var raw1 = `# Big Changelog

introduction stuff

##         unRelEASed - don't release me!

not ready

## 1.2.3 2016-10-10

Last Release

## 1.2.0 2016-01-01

First Release
`

var raw2 = `# Big Changelog

introduction stuff

## Unreleased

not ready

## 1.2.3

Last Release, invalid, missing date

## 1.2.0 2016-01-01

First Release
`

func TestCLParse(t *testing.T) {
	cl, err := Parse(raw1)
	if err != nil {
		t.Fatalf("could not parse: %s", err)
	}
	if len(cl.Released) != 2 {
		t.Errorf("Expected two releases, got %d", len(cl.Released))
	}
	if cl.Unreleased != "not ready" {
		t.Errorf("Parse error on cl.Unreleased: %q", cl.Unreleased)
	}
	if cl.Intro != "# Big Changelog\n\nintroduction stuff" {
		t.Errorf("Parse error of cl.Intro: %q", cl.Intro)
	}
}

func TestCLParseError(t *testing.T) {
	_, err := Parse(raw2)
	if err == nil {
		t.Errorf("invalid changelog was parsed!")
	}
}

func TestMostRecent(t *testing.T) {
	// MostRecent on empty changelog should return empty Entry
	cl := ChangeLog{}
	want := Entry{}
	got := cl.Top()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("cl.MostRecent = %v, want %v", got, want)
	}
	cl, err := Parse(raw1)
	if err != nil {
		t.Fatalf("could not parse: %s", err)
	}
	want = Entry{"1.2.3", "2016-10-10", "Last Release"}
	got = cl.Top()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("cl.MostRecent = %v, want %v", got, want)
	}
}
