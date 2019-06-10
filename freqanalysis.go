package freqanalysis

import (
	"sort"
	"strings"
)

type WordStat struct {
	word  string
	count int
}

type WordStats []WordStat

func (ws WordStats) Len() int           { return len(ws) }
func (ws WordStats) Less(i, j int) bool { return ws[i].count < ws[j].count }
func (ws WordStats) Swap(i, j int)      { ws[i], ws[j] = ws[j], ws[i] }

func Freq(str string) WordStats {

	statsTemp := make(map[string]int)

	words := strings.Split(str, " ")
	for _, word := range words {
		statsTemp[word] = statsTemp[word] + 1
	}

	stats := make(WordStats, len(statsTemp))
	i := 0
	for word, count := range statsTemp {
		stats[i] = WordStat{word, count}
		i++
	}

	sort.Sort(sort.Reverse(stats))

	return stats
}
