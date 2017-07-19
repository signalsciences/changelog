// Package changelog provides functionality for managing markdown changelogs
package changelog

import (
	"bytes"
	"fmt"
	"strings"
)

// Entry represents a single changelog entry
type Entry struct {
	Version string // TBD SEMVER
	Date    string // TBD a real datetime
	Notes   string
}

// MarshalText satisfies the TextMarshaler interface
func (e Entry) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("## " + e.Version + " " + e.Date + "\n\n")
	buf.WriteString(e.Notes)
	buf.WriteString("\n")
	return buf.Bytes(), nil
}

func (e Entry) String() string {
	raw, _ := e.MarshalText()
	return string(raw)
}

// ChangeLog represents the parts of a changelog
type ChangeLog struct {
	Intro      string
	Unreleased string
	Released   []Entry
}

// Top returns the first entry listed in the file
// (and not first release, oldest)
func (cl ChangeLog) Top() Entry {
	if len(cl.Released) == 0 {
		return Entry{}
	}
	return cl.Released[0]
}

// MarshalText satisfies the TextMarshaler interface
func (cl ChangeLog) MarshalText() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(cl.Intro)
	buf.WriteString("\n")
	if cl.Unreleased != "" {
		buf.WriteString("## Unreleased\n\n")
		buf.WriteString(cl.Unreleased)
		buf.WriteString("\n")
	}
	for _, e := range cl.Released {
		buf.WriteString("\n")
		ebytes, err := e.MarshalText()
		if err != nil {
			return buf.Bytes(), err
		}
		buf.Write(ebytes)
	}
	return buf.Bytes(), nil
}

// String generates a normalized ChangeLog
func (cl ChangeLog) String() string {
	raw, _ := cl.MarshalText()
	return string(raw)
}

// parses a "VERSION [space] YYYY-MM-DD" line
func parseHeader(s string) (Entry, error) {
	e := Entry{}

	// maybe use strings.Fields
	parts := strings.SplitN(s, " ", 2)
	if len(parts) != 2 {
		return e, fmt.Errorf("Unable to find date in %q", s)
	}
	e.Version = strings.TrimSpace(parts[0])
	e.Date = strings.TrimSpace(parts[1])
	return e, nil
}

// Parse gets most recent entry (i.e. the top entry)
//
// TODO: change to []byte since this is likely to come from
// a File or io.Reader
func Parse(raw string) (ChangeLog, error) {
	cl := ChangeLog{}
	paragraphs := strings.Split(raw, "\n## ")
	if len(paragraphs) == 0 {
		return cl, fmt.Errorf("did not find any releases")
	}

	cl.Intro = strings.TrimSpace(paragraphs[0])

	paragraphs = paragraphs[1:]

	released := make([]Entry, 0, len(paragraphs))
	for n, p := range paragraphs {
		parts := strings.SplitN(p, "\n", 2)
		if len(parts) != 2 {
			return cl, fmt.Errorf("Unable to get release header from %q", p)
		}

		// Unreleased can only be the first entry
		if strings.HasPrefix(strings.ToLower(strings.TrimSpace(parts[0])), "unreleased") {
			if n != 0 {
				return cl, fmt.Errorf("Got an unreleased section multiple times")
			}
			cl.Unreleased = strings.TrimSpace(parts[1])
			continue
		}
		e, err := parseHeader(parts[0])
		if err != nil {
			return cl, err
		}
		e.Notes = strings.TrimSpace(parts[1])
		released = append(released, e)
	}
	cl.Released = released
	return cl, nil
}
