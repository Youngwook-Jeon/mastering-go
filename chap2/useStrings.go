package main

import (
	"fmt"
	s "strings"
	"unicode"
)

// f is a alias for fmt.Printf
var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title("tHis wiLL be A title!"))

	f("EqualFold: %v\n", s.EqualFold("Lucas", "luc"))
	f("EqualFold: %v\n", s.EqualFold("Lucas", "lUCAs"))

	f("Prefix: %v\n", s.HasPrefix("Lucas", "Lu"))
	f("Prefix: %v\n", s.HasPrefix("Lucas", "lu"))
	f("Suffix: %v\n", s.HasSuffix("Lucas", "as"))
	f("Suffix: %v\n", s.HasSuffix("Lucas", "AS"))

	f("Index: %v\n", s.Index("Lucas", "ca"))
	f("Index: %v\n", s.Index("Lucas", "s"))
	f("Count: %v\n", s.Count("Lucas", "b"))
	f("Count: %v\n", s.Count("Lucas", "u"))
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
