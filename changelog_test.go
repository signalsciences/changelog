package changelog

import (
	"bytes"
	"log"
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

var raw3 = `# Big Changelog

introduction stuff

## Unreleased

whatever

## 1.2.3 2016-01-04

Last Release

## 1.2.2 2016-01-03

Release 3

## 1.2.1 2016-01-02

Release 2

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

func TestFindByVersion(t *testing.T) {
	cl, err := Parse(raw1)
	if err != nil {
		t.Error(err)
	}

	want := Entry{"1.2.0", "2016-01-01", "First Release"}
	got, err := cl.FindByVersion("1.2.0")

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("cl.FindByVersion = %v, want %v", got, want)
	}

	_, err = cl.FindByVersion("not found")
	if err == nil {
		t.Error("err should not be nil")
	}
}

func TestGetRange(t *testing.T) {
	cl, err := Parse(raw3)
	if err != nil {
		t.Error(err)
	}

	want := []Entry{
		{"1.2.3", "2016-01-04", "Last Release"},
		{"1.2.2", "2016-01-03", "Release 3"},
		{"1.2.1", "2016-01-02", "Release 2"},
	}

	got, err := cl.GetRange("1.2.0", "1.2.3")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("t.GetRange: wanted %v, got %v", want, got)
	}

	// if from or to are invalid, return error

	_, err = cl.GetRange("1.1.1", "1.2.3")
	if err == nil {
		t.Error("should have err")
	}

	_, err = cl.GetRange("1.2.0", "1.2.4")
	if err == nil {
		t.Error("should have err")
	}

	// if to is higher than from, err
	_, err = cl.GetRange("1.2.3", "1.2.0")
	if err == nil {
		t.Error("should err")
	}
}

func TestMarshalText(t *testing.T) {
	cl := ChangeLog{
		Unreleased: "unreleased test",
		Released: []Entry{
			{"1.2.3", "2017-01-01", "notes"},
		},
	}
	b, err := cl.MarshalText()
	if err != nil {
		log.Fatal(err)
	}

	want := []byte(`
## Unreleased

unreleased test

## 1.2.3 2017-01-01

notes
`)
	if !bytes.Equal(b, want) {
		t.Errorf("cl.MarshalText() = %q, want %q", string(b), string(want))
	}

}
