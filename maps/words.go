package maps

import (
	"fmt"
	"strings"
)

func count(text string) map[string]int {
	// makes our paragraph into an array of lowercase words.
	words := strings.Fields(strings.ToLower(text))
	// initializes an empty map (of our return type)
	freq := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Trim(word, `.,:"-`)
		freq[word]++
	}

	return freq
}

func Words() {
	text := `As far as eye could reach he saw nothing but the stems of the
	great plants about him receding in the violet shade, and far overhead
	the multiple transparency of huge leaves filtering the sunshine to the
	solemn splendour of twilight in which he walked. Whenever he felt able
	he ran again; the ground continued soft and springy, covered with the
	same resilient weed which was the first thing his hands had touched in
	Malacandra. Once or twice a small red creature scuttled across his
	path, but otherwise there seemed to be no life stirring in the wood;
	nothing to fear -- except the fact of wandering unprovisioned and alone
	in a forest of unknown vegetation thousands or millions of miles
	beyond the reach or knowledge of man.`

	frequency := count(text)

	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%d %v\n", count, word)
		}
	}
}
