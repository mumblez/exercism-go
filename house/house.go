// Package house builds the song programmatically
package house

import "strings"

type animal struct {
	name   string
	action string
}

var animals = []animal{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func Embed(middle, end string) string {
	return middle + " " + end
}

func Verse(start string, middle []string, end string) string {
	return start + " " + strings.Join(append(middle, end), " ")
}

func buildVerse(v int) string {
	v = v - 1
	var song string
	song += "This is the " + animals[v].name

	for n := v; n != 0; n-- {
		song += "\nthat " + animals[n].action + " the " + animals[n-1].name
	}
	return song
}

// Song returns the entire song
func Song() (song string) {
	for i := 1; i <= len(animals); i++ {
		if i > 1 {
			song += "\n\n"
		}
		song += buildVerse(i)
	}
	return
}
