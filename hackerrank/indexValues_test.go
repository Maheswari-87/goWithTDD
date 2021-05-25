package main

import (
	"reflect"
	"testing"
)

func TestIndex(t *testing.T) {
	/*t.Run("Getting odd index", func(t *testing.T) {
		got := GetOddStrings([]string{"Hacker", "Rank"})
		want := []string{"akr", "ak"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Getting odd index", func(t *testing.T) {
		got := GetEvenStrings([]string{"Hacker", "Rank"})
		want := []string{"Hce", "Rn"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})*/
	got := EvenOdd([]string{"Hacker", "Rank"})
	want := []string{"Hce", "akr", "Rn", "ak"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
