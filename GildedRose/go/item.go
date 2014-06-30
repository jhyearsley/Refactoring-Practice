package gildedrose

import "strconv"

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (i Item) String() string {
	return i.Name + ", " + strconv.Itoa(i.SellIn) + ", " + strconv.Itoa(i.Quality)
}
