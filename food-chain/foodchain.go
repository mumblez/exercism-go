// Package foodchain builds the song programmatically
package foodchain

type animal struct {
	name     string
	sentence string
}

const (
	testVersion = 2
	iknow       = "I know an old lady who swallowed a "
)

var animals = []animal{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func buildVerse(v int) string {
	v = v - 1
	var song string
	song += iknow + animals[v].name + ".\n" + animals[v].sentence

	if animals[v].name == "horse" {
		return song
	}

	for n := v; n != 0; n-- {
		song += "\nShe swallowed the " + animals[n].name + " to catch the " + animals[n-1].name
		if animals[n-1].name == "spider" {
			song += " that wriggled and jiggled and tickled inside her."
		} else {
			song += "."
		}
		if n == 1 {
			song += "\n" + animals[0].sentence
		}

	}
	return song
}

// Verse returns a specific verse out of 8
func Verse(n int) (song string) {
	song = buildVerse(n)
	return
}

// Verses returns 2 verses joined together
func Verses(v1, v2 int) (song string) {
	song += buildVerse(v1) + "\n\n" + buildVerse(v2)
	return
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
