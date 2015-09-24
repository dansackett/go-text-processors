# Golang Text Processors

This library provides basic text processing based heavily off of [this article](http://matt.might.net/articles/shell-scripts-for-passive-voice-weasel-words-duplicates/) and other sources around the web.
It attempts to make helpful suggestions on how your can improve your writing
by following a few guidelines.

## Processors

The following processors have been provided:

### Passive Voice

Passive voice, while sometimes acceptable, can be a roundabout way of saying
something affirmative. It is good to recognize the passive voice and form your
own conclusion about whether it is acceptable or not.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "We were taught how to write novels in English class"
	p := PassiveVoiceProcessor()
	res := p.Run(s)

    // Original text
    fmt.Println(res.Text)

    for _, match := range res.Matches {
        // Show the matched phrase
        fmt.Println(match.Match)
        // Show the indices the phrase is between
        fmt.Println(match.Indices)
    }
}
```

### Weasel Words

A weasel word is an informal term for words and phrases aimed at creating an
impression that a specific and/or meaningful statement has been made, when
only a vague or ambiguous claim has been communicated, enabling the specific
meaning to be denied if the statement is challenged.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "These are a number of reasons why we do it this way."
	p := WeaselWordProcessor()
	res := p.Run(s)

    ...
}
```

### Wordy Phrases

Wordy phrases refer to words and phrases which may be more complex and
unnecessary in your writing because a simpler, more effective, choice can be
used.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "The point I am trying to make is that this is not right."
	p := TooWordyProcessor()
	res := p.Run(s)

    ...
}
```

### Adverbs

Adverbs have a tendency to falsely make your writing stronger while actually
making it more wordy and superfluous.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "He suspiciously crept around the corner"
	p := AdverbProcessor()
	res := p.Run(s)

    ...
}
```

### Cliches

Cliches are common phrases which have been overused to the point that they
lost their original effect.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "She woke up on the wrong side of the bed."
	p := ClicheProcessor()
	res := p.Run(s)

    ...
}
```

### Lexical Illusions

Lexical illusions refer to words that your brain skips when reading as it
tries to correct the errors for you. In this case, it finds words that have
been repeated, whether on the same line or the next line, that your brain
misses when proofreading.

```
package main

import (
    "fmt"
    processors "github.com/dansackett/go-text-processors"
)

func main () {
	s := "Sometimes the the brain skips double words when proofreading."
	p := LexicalIllusionProcessor()
	res := p.Run(s)

    ...
}
```

## Todo

- Create a CLI tool to allow users to process `stdin` and return finding to `stdout`.
- Potentially add support for other processor types (sentence length, sentence starters, etc)
