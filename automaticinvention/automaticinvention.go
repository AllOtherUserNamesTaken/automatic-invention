package automaticinvention

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Performance struct {
	Team  string
	Score int64
}

func (p *Performance) String() string {
	return fmt.Sprintf("Performance( Team='%v', Score=%v )", p.Team, p.Score)
}

const (
	LOSS int64 = 0
	DRAW int64 = 1
	WIN  int64 = 3
)

func (First *Performance) RelativeScore(Second *Performance) int64 {
	if First.Score == Second.Score {
		return DRAW
	}
	if First.Score > Second.Score {
		return WIN
	}
	return LOSS
}

func PerformanceFromString(performance string) *Performance {
	performance = strings.TrimSpace(performance)
	last := strings.LastIndex(performance, " ")
	score, _ := strconv.ParseInt(strings.TrimSpace(performance[last:]), 10, 64)
	team := strings.TrimSpace(performance[0:last])
	return &Performance{
		Team:  team,
		Score: score,
	}
}

type Game struct {
	First  *Performance
	Second *Performance
}

func (g *Game) String() string {
	return fmt.Sprintf("Game( First=%v, Second=%v )", g.First, g.Second)
}

func GameFromString(game string) *Game {
	game = strings.TrimSpace(game)
	performances := strings.Split(game, ",")
	First := PerformanceFromString(performances[0])
	Second := PerformanceFromString(performances[1])
	return &Game{
		First:  First,
		Second: Second,
	}
}

func (g *Game) RankedGame() *Game {
	First := g.First.RelativeScore(g.Second)
	Second := g.Second.RelativeScore(g.First)

	g.First.Score = First
	g.Second.Score = Second

	return g
}

type ICounter interface {
	update(gm *Game) *ICounter
	add(p *Performance) *ICounter
}

type Counter struct {
	Cnt map[string]int64
}

func (cnt *Counter) Init() *Counter {
	if len(cnt.Cnt) == 0 {
		cnt.Cnt = make(map[string]int64)
	}
	return cnt
}

func (cnt *Counter) AsArray() []*Performance {
	i := 0
	result := make([]*Performance, len(cnt.Cnt))
	for k, v := range cnt.Cnt {
		result[i] = &Performance{
			Team:  k,
			Score: v,
		}
		i = i + 1
	}
	return result
}

func (cnt *Counter) Add(p *Performance) *Counter {

	val, ok := cnt.Cnt[p.Team]
	if ok {
		cnt.Cnt[p.Team] = val + p.Score
	} else {
		cnt.Cnt[p.Team] = p.Score
	}
	return cnt
}

func (cnt *Counter) Update(g *Game) *Counter {
	return cnt.Add(g.First).Add(g.Second)
}

type CounterSorter struct {
	Performances []*Performance
	by           func(p1, p2 *Performance) bool
}

type By func(p1, p2 *Performance) bool

func (by By) Sort(performances []*Performance) {
	ps := &CounterSorter{
		Performances: performances,
		by:           by,
	}
	sort.Sort(ps)
}

func (s *CounterSorter) Len() int {
	return len(s.Performances)
}

func (s *CounterSorter) Swap(i, j int) {
	s.Performances[i], s.Performances[j] = s.Performances[j], s.Performances[i]
}

func (s *CounterSorter) Less(i, j int) bool {
	return s.by(s.Performances[i], s.Performances[j])
}

func SortByPerformance(p1, p2 *Performance) bool {
	if p1.Score == p2.Score {
		return p1.Team < p2.Team
	}
	return p1.Score > p2.Score
}
