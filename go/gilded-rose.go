package main

const MaxQuality = 50

const QualitySulfuras = 80

type ItemInterface interface {
	Name() string
	SellIn() int
	Quality() int
	Update()
}

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) Name() string {
	return item.name
}

func (item *Item) SellIn() int {
	return item.sellIn
}

func (item *Item) Quality() int {
	return item.quality
}

func (item *Item) Update() {
	item.decreaseQuality(1)

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.decreaseQuality(1)
	}
}

func (item *Item) decreaseQuality(amt int) {
	quality := item.quality - amt
	if quality < 0 {
		quality = 0
	}
	item.quality = quality
}

func (item *Item) incrementQuality() {
	if item.quality < MaxQuality {
		item.quality += 1
	}
}

type AgedBrie struct {
	*Item
}

func NewAgedBrie(name string, sellIn, quality int) *AgedBrie {
	return &AgedBrie{&Item{name, sellIn, quality}}
}

func (item *AgedBrie) Update() {
	item.incrementQuality()
	item.sellIn -= 1
	if item.sellIn < 0 {
		item.incrementQuality()
	}
}

type BackstagePasses struct {
	*Item
}

func NewBackStagePasses(name string, sellIn, quality int) *BackstagePasses {
	return &BackstagePasses{&Item{name, sellIn, quality}}
}

func (item *BackstagePasses) Update() {
	item.incrementQuality()
	if item.sellIn < 11 {
		item.incrementQuality()
	}
	if item.sellIn < 6 {
		item.incrementQuality()
	}

	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = 0
	}
}

type Sulfuras struct {
	*Item
}

func NewSulfuras(name string) *Sulfuras {
	return &Sulfuras{&Item{name, 0, QualitySulfuras}}
}

func (item *Sulfuras) Update() {}

type ConjuredItem struct {
	*Item
}

func NewConjuredItem(name string, sellIn, quality int) *ConjuredItem {
	return &ConjuredItem{&Item{name, sellIn, quality}}
}

func (item *ConjuredItem) Update() {
	item.decreaseQuality(2)

	item.sellIn -= 1

	if item.sellIn < 0 {
		item.decreaseQuality(2)
	}
}

var items = []ItemInterface{
	&Item{"+5 Dexterity Vest", 10, 20},
	NewAgedBrie("Aged Brie", 2, 0),
	&Item{"Elixir of the Mongoose", 5, 7},
	NewSulfuras("Sulfuras, Hand of Ragnaros"),
	NewBackStagePasses("Backstage passes to a TAFKAL80ETC concert", 15, 20),
	NewConjuredItem("Conjured Mana Cake", 3, 6),
}

func main() {
	// fmt.Println("OMGHAI!")

	// fmt.Println("Before update")
	// printItems(items)

	GildedRose(items)

	// fmt.Println("After update")
	// printItems(items)
}

// func printItems(items []ItemInterface) {
// 	for _, item := range items {
// 		fmt.Println(item.Name(), item.SellIn(), item.Quality())
// 	}
// }

func GildedRose(items []ItemInterface) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		item.Update()
	}
}
