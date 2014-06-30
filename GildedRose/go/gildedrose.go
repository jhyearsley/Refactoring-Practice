package gildedrose

type GildedRose struct {
	Items []Item
}

func (g GildedRose) UpdateQuality() {
	for i := 0; i < len(g.Items); i += 1 {
		if g.Items[i].Name != "Aged Brie" && g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if g.Items[i].Quality > 0 {
				if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
					g.Items[i].Quality = g.Items[i].Quality - 1
				}
			}
		} else {
			if g.Items[i].Quality < 50 {
				g.Items[i].Quality = g.Items[i].Quality + 1

				if g.Items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if g.Items[i].SellIn < 11 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}

					if g.Items[i].SellIn < 6 {
						if g.Items[i].Quality < 50 {
							g.Items[i].Quality = g.Items[i].Quality + 1
						}
					}
				}
			}
		}

		if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
			g.Items[i].SellIn = g.Items[i].SellIn - 1
		}

		if g.Items[i].SellIn < 0 {
			if g.Items[i].Name != "Aged Brie" {
				if g.Items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if g.Items[i].Quality > 0 {
						if g.Items[i].Name != "Sulfuras, Hand of Ragnaros" {
							g.Items[i].Quality = g.Items[i].Quality - 1
						}
					}
				} else {
					g.Items[i].Quality = g.Items[i].Quality - g.Items[i].Quality
				}
			} else {
				if g.Items[i].Quality < 50 {
					g.Items[i].Quality = g.Items[i].Quality + 1
				}
			}
		}
	}
}
