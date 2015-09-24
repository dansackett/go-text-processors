package text_processors

import (
	"testing"
)

func TestRegexCustom(t *testing.T) {
	re := "p([a-z]+)ch"
	s := "peach pinch punch"

	res := regexCustom(s, re)

	if res == nil {
		t.Errorf("regexCustom returned an empty response")
	}

	if len(res.Matches) != 3 {
		t.Errorf("regexCustom found %d matches, there should be 3", len(res.Matches))
	}
}

func TestRegexCustomEmpty(t *testing.T) {
	re := "p([a-z]+)ch"
	s := "pizza"

	res := regexCustom(s, re)

	if res == nil {
		t.Errorf("regexCustom returned an empty response")
	}

	if len(res.Matches) != 0 {
		t.Errorf("regexCustom found %d matches, there should be 0", len(res.Matches))
	}
}

func TestPassiveVoiceProcessor(t *testing.T) {
	s := "We were taught how to write novels in English class"
	p := PassiveVoiceProcessor()
	res := p.Run(s)

	if res == nil {
		t.Errorf("PassiveVoiceProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("PassiveVoiceProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "were taught" {
			t.Errorf("PassiveVoiceProcessor found \"%s\" as a passive match when it should have found \"were taught\"", match)
		}

		if indices[0] != 3 && indices[1] != 14 {
			t.Errorf("PassiveVoiceProcessor found indices [%d %d] and should have found [3 14]", indices[0], indices[1])
		}
	}
}

func TestWeaselWordProcessor(t *testing.T) {
	s := "These are a number of reasons why we do it this way."
	p := WeaselWordProcessor()
	res := p.Run(s)

	if res == nil {
		t.Errorf("WeaselWordProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("WeaselWordProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "are a number" {
			t.Errorf("WeaselWordProcessor found \"%s\" as a weasel match when it should have found \"are a number\"", match)
		}

		if indices[0] != 6 && indices[1] != 18 {
			t.Errorf("WeaselWordProcessor found indices [%d %d] and should have found [6 18]", indices[0], indices[1])
		}
	}
}

func TestTooWordyProcessor(t *testing.T) {
	s := "The point I am trying to make is that this is not right."
	p := TooWordyProcessor()
	res := p.Run(s)

	if res == nil {
		t.Errorf("TooWordyProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("TooWordyProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "The point I am trying to make" {
			t.Errorf("TooWordyProcessor found \"%s\" as a wordy phrase when it should have found \"The point I am trying to make\"", match)
		}

		if indices[0] != 0 && indices[1] != 29 {
			t.Errorf("TooWordyProcessor found indices [%d %d] and should have found [0 29]", indices[0], indices[1])
		}
	}
}

func TestAdverbProcessor(t *testing.T) {
	s := "He suspiciously crept around the corner."
	p := AdverbProcessor()
	res := p.Run(s)

	if res == nil {
		t.Errorf("AdverbProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("AdverbProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "suspiciously" {
			t.Errorf("AdverbProcessor found \"%s\" as an adverb when it should have found \"suspiciously\"", match)
		}

		if indices[0] != 3 && indices[1] != 15 {
			t.Errorf("AdverbProcessor found indices [%d %d] and should have found [3 15]", indices[0], indices[1])
		}
	}
}

func TestClicheProcessor(t *testing.T) {
	s := "She woke up on the wrong side of the bed."
	p := ClicheProcessor()
	res := p.Run(s)

	if res == nil {
		t.Errorf("ClicheProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("ClicheProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "wrong side of the bed" {
			t.Errorf("ClicheProcessor found \"%s\" as a cliche it should have found \"wrong side of the bed\"", match)
		}

		if indices[0] != 19 && indices[1] != 40 {
			t.Errorf("ClicheProcessor found indices [%d %d] and should have found [19 40]", indices[0], indices[1])
		}
	}
}

func TestLexicalIllusionProcessor(t *testing.T) {
	s1 := "Sometimes the the brain skips double words when proofreading."
	p := LexicalIllusionProcessor()
	res := p.Run(s1)

	if res == nil {
		t.Errorf("LexicalIllusionProcessor returned an empty response")
	}

	if len(res.Matches) != 1 {
		t.Errorf("LexicalIllusionProcessor found %d matches, there should be 1", len(res.Matches))
	}

	if len(res.Matches) > 0 {
		match := res.Matches[0].Match
		indices := res.Matches[0].Indices

		if match != "the" {
			t.Errorf("LexicalIllusionProcessor found \"%s\" as a lexical illusion it should have found \"the\"", match)
		}

		if indices[0] != 14 && indices[1] != 17 {
			t.Errorf("LexicalIllusionProcessor found indices [%d %d] and should have found [14 17]", indices[0], indices[1])
		}
	}
}
