package gildedrose

import "testing"

func TestGildedRose(t *testing.T) {
	items := []Item{Item{"foo", 0, 0}}
	app := GildedRose{items}

	app.UpdateQuality()
	if app.Items[0].Name != "fixme" {
		t.Fatalf("Fix me")
	}
}
