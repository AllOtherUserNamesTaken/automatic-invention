package automaticinvention

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGameFromString(t *testing.T) {
	actual := GameFromString("Lions 3, Snakes 3")
	expected := new(Game)
	expected.First = new(Performance)
	expected.Second = new(Performance)
	expected.First.Team = "Lions"
	expected.First.Score = 3
	expected.Second.Team = "Snakes"
	expected.Second.Score = 3

	if !cmp.Equal(actual, expected) {
		t.Log(actual, " != ", expected)
		t.Fail()
	}
}

func TestCounter(t *testing.T) {
	game := GameFromString("Lions 3, Snakes 3").RankedGame()
	actual := new(Counter).Init()
	actual = actual.Update(game)
	expected := map[string]int64{
		"Lions":  1,
		"Snakes": 1,
	}

	if !cmp.Equal(actual.Cnt, expected) {
		t.Log(actual.Cnt, " != ", expected)
		t.Fail()
	}
}

func TestSortByPerformanc(t *testing.T) {
	game := GameFromString("Lions 3, Snakes 3").RankedGame()
	cactual := new(Counter).Init()
	cactual = cactual.Update(game)

	carry := cactual.AsArray()
	By(SortByPerformance).Sort(carry)

	expected := []*Performance{
		{"Lions", 1},
		{"Snakes", 1},
	}

	if !cmp.Equal(carry, expected) {
		t.Log(carry, " != ", expected)
		t.Fail()
	}
}
