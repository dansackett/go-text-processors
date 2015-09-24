package text_processors

import (
	"fmt"
	"regexp"
	"strings"
)

// Match represents a found item in a string of text. It allows users to grab
// the match itself and see the index points in which the match lies in the text
type Match struct {
	Match   string
	Indices []int
}

// MatchResult represents the result of processing. It provides a way to view the
// original text and the matches found within.
type MatchResult struct {
	Text    string
	Matches []*Match
}

// TextProcessor is an interface providing a sane way to run each processor
// type with a simple Run method.
type TextProcessor interface {
	Run(string) (*MatchResult, error)
}

// regexCustom takes in a string of text and a regular expression statement
// returning a MatchResult object. It is the barebones implementation of what
// processors do.
func regexCustom(s string, re string) *MatchResult {
	var res []*Match

	r := regexp.MustCompile(re)
	matches := r.FindAllString(s, -1)
	idxs := r.FindAllStringIndex(s, -1)

	if len(matches) == 0 {
		return &MatchResult{Text: s, Matches: res}
	}

	for i, m := range matches {
		res = append(res, &Match{Match: m, Indices: idxs[i]})
	}

	return &MatchResult{Text: s, Matches: res}
}

// passiveVoice is a private object for finding matches of the passive voice
// in a string of text.
type passiveVoice struct{}

// PassiveVoiceProcessor builds a references to the private passiveVoice
// object giving us a way to use it outside the context of this library.
func PassiveVoiceProcessor() *passiveVoice {
	return &passiveVoice{}
}

// Run satisfies the TextPrcoessor interface for the passiveVoice object
// returning a MatchResult based on what it finds.
func (p *passiveVoice) Run(s string) *MatchResult {
	return regexCustom(s, fmt.Sprintf("\\b(?i)(?:am|are|were|being|is|been|was|be)\\b\\s*(?:[\\w]+ed|%s)\\b", strings.Join(PassiveIrregulars, "|")))
}

// weaselWords is a private object for finding matches of potential weasel
// words in a string of text.
type weaselWords struct{}

// WeaselWordProcessor builds a references to the private weaselWords
// object giving us a way to use it outside the context of this library.
func WeaselWordProcessor() *weaselWords {
	return &weaselWords{}
}

// Run satisfies the TextPrcoessor interface for the weaselWords object
// returning a MatchResult based on what it finds.
func (p *weaselWords) Run(s string) *MatchResult {
	return regexCustom(s, fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Weasels, "|")))
}

// tooWordy is a private object for finding matches of potential wordy phrases
// in a string of text.
type tooWordy struct{}

// TooWordyProcessor builds a references to the private tooWordy
// object giving us a way to use it outside the context of this library.
func TooWordyProcessor() *tooWordy {
	return &tooWordy{}
}

// Run satisfies the TextPrcoessor interface for the tooWordy object
// returning a MatchResult based on what it finds.
func (p *tooWordy) Run(s string) *MatchResult {
	return regexCustom(s, fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Wordy, "|")))
}

// adverbs is a private object for finding matches of adverbs in a string of text.
type adverbs struct{}

// AdverbProcessor builds a references to the private adverbs
// object giving us a way to use it outside the context of this library.
func AdverbProcessor() *adverbs {
	return &adverbs{}
}

// Run satisfies the TextPrcoessor interface for the adverbs object
// returning a MatchResult based on what it finds.
func (p *adverbs) Run(s string) *MatchResult {
	return regexCustom(s, fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Adverbs, "|")))
}

// cliches is a private object for finding matches of cliches in a string of text.
type cliches struct{}

// ClicheProcessor builds a references to the private cliches
// object giving us a way to use it outside the context of this library.
func ClicheProcessor() *cliches {
	return &cliches{}
}

// Run satisfies the TextPrcoessor interface for the cliches object
// returning a MatchResult based on what it finds.
func (p *cliches) Run(s string) *MatchResult {
	return regexCustom(s, fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Cliches, "|")))
}

// lexicalIllusion is a private object for finding matches of lexical
// illusions in a string of text.
type lexicalIllusion struct{}

// LexicalIllusionProcessor builds a references to the private lexicalIllusion
// object giving us a way to use it outside the context of this library.
func LexicalIllusionProcessor() *lexicalIllusion {
	return &lexicalIllusion{}
}

// Run satisfies the TextPrcoessor interface for the lexicalIllusion object
// returning a MatchResult based on what it finds.
func (p *lexicalIllusion) Run(s string) *MatchResult {
	var res []*Match
	var lastMatch string

	initialRes := regexCustom(s, "(\\w+)")

	for _, match := range initialRes.Matches {
		if match.Match == lastMatch {
			res = append(res, &Match{Match: match.Match, Indices: match.Indices})
		}

		lastMatch = strings.ToLower(match.Match)
	}

	return &MatchResult{Text: s, Matches: res}
}
