package hscan

import (
	"fmt"
	"testing"
	"time"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("90f2c9c53f66540e67349e0ab83d8cd0", "../main/Top304Thousand-probable-v2.txt")
	want := "p@ssword"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGenHashMaps(t *testing.T) {
	start := time.Now()
	GenHashMaps("../main/Top304Thousand-probable-v2.txt")
	duration := time.Since(start)

	fmt.Println(duration.Nanoseconds())

	// nonoseconds before any performance updates: 715893000 or 2355 per password
	// nanoseconds with subroutines:			   619698700 or 2038 per password
}
