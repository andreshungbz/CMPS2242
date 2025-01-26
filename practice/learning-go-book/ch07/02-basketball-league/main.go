package main

import (
	"io"
	"os"
	"sort"
)

// 1 - types

type Team struct {
	Name string
	Players []string
}

type League struct {
	Name string
	Teams map[string]Team
	Wins map[string]int
}

// 2 - methods

func (l *League) MatchResult(t1 string, t1Score int, t2 string, t2Score int) {
	// handle non existent team name
	if _, ok := l.Teams[t1]; !ok {
		return
	}
	if _, ok := l.Teams[t2]; !ok {
		return
	}

	if t1Score == t2Score {
		return
	}

	if t1Score > t2Score {
		l.Wins[t1]++
	}

	if t2Score > t1Score {
		l.Wins[t2]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))

	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

// 3 - interface

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)

	RankPrinter(l, os.Stdout)
}