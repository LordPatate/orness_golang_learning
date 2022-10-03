package exercise

import (
	"reflect"
	"testing"
)

var fishypedia = map[string]*Fish{
	"cod":    &Fish{12, 75, "Atlantic"},
	"salmon": &Fish{5, 70, "Atlantic"},
	"bass":   &Fish{5, 50, "Mediterranean"},
	"mako":   &Fish{180, 300, "Atlantic"},
}

func TestMorningFishing(t *testing.T) {
	fishesCaught := []string{
		"cod", "salmon", "salmon", "boot",
	}
	expected := map[int][]string{
		12: {"cod"},
		5:  {"salmon", "salmon"},
	}

	result := GetFishNamesByWeight(fishesCaught, fishypedia)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got: %v, want: %v\n", result, expected)
	}
}

func TestEveningFishing(t *testing.T) {
	fishesCaught := []string{
		"salmon", "mako", "salmon", "bass",
	}
	expected := map[int][]string{
		5:   {"salmon", "salmon", "bass"},
		180: {"mako"},
	}

	result := GetFishNamesByWeight(fishesCaught, fishypedia)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got: %v, want: %v\n", result, expected)
	}
}
