package freqanalysis

import (
	"sort"
	"strings"
	"unicode"
)

type WordStat struct {
	word  string
	count int
}

type WordStats []WordStat

func (ws WordStats) Len() int           { return len(ws) }
func (ws WordStats) Less(i, j int) bool { return ws[i].count < ws[j].count }
func (ws WordStats) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }

func filter(str string) []string {

	list := []string{}
	word := strings.Builder{}

	for _, s := range str {
		if unicode.IsLetter(s) || unicode.IsDigit(s) {
			word.WriteRune(s)
		} else {
			if word.Len() > 0 {
				list = append(list, word.String())
				word.Reset()
			}
		}
	}

	if word.Len() > 0 {
		list = append(list, word.String())
	}

	return list
}

func Freq(str string) WordStats {

	statsTemp := make(map[string]int)

	// words := strings.Split(str, " ")
	words := filter(str)

	for _, word := range words {
		word = strings.ToLower(word)
		statsTemp[word] += 1
	}

	stats := make(WordStats, len(statsTemp))
	i := 0
	for word, count := range statsTemp {
		stats[i] = WordStat{word, count}
		i++
	}

	sort.Sort(sort.Reverse(stats))

	if len(stats) < 10 {
		return stats
	}
	return stats[:10]
}
