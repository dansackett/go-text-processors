package text_processors

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	PassiveRegex  = regexp.MustCompile(fmt.Sprintf("\\b(?i)(?:am|are|were|being|is|been|was|be)\\b\\s*(?:[\\w]+ed|%s)\\b", strings.Join(PassiveIrregulars, "|")))
	WeaselRegex   = regexp.MustCompile(fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Weasels, "|")))
	WordyRegex    = regexp.MustCompile(fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Wordy, "|")))
	AdverbRegex   = regexp.MustCompile(fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Adverbs, "|")))
	ClicheRegex   = regexp.MustCompile(fmt.Sprintf("\\b(?i)(%s)\\b", strings.Join(Cliches, "|")))
	IllusionRegex = regexp.MustCompile("(\\w+)")
)
