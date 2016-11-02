package anagram

//package main

import (
	"bytes"
	"strings"
)

func stringBytes(s string) []byte {
	var sb []byte
	for i := 0; i < len(s); i++ {
		sb = append(sb, s[i])
	}
	return sb
}

func notContain(res []string, s string) bool {
	for _, v := range res {
		if v == s {
			return false
		}
	}
	return true
}

// Detect anagrams
func Detect(subject string, candidates []string) []string {
	var results []string
	subject = strings.ToLower(subject)
	for _, v := range candidates {
		v = strings.ToLower(v)
		if len(v) != len(subject) || subject == v {
			continue
		}
		s := stringBytes(subject)
		for i := 0; i < len(v); i++ {
			if bytes.ContainsRune(s, rune(v[i])) {
				index := bytes.IndexByte(s, v[i])
				s = append(s[:index], s[index+1:]...)
			}
		}
		if len(s) == 0 && notContain(results, v) {
			results = append(results, v)
		}
	}
	return results
}
