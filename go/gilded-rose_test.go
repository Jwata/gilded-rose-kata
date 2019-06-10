package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestItem struct{ mock.Mock }

func (item *TestItem) Name() string { return "Mock item" }

func (item *TestItem) SellIn() int { return 1 }

func (item *TestItem) Quality() int { return 1 }

func (item *TestItem) Update() { item.Called() }

func Test_GildedRose(t *testing.T) {
	item1 := &TestItem{}
	item2 := &TestItem{}
	item1.On("Update").Return(nil)
	item2.On("Update").Return(nil)

	testItems := []ItemInterface{item1, item2}
	GildedRose(testItems)

	item1.AssertCalled(t, "Update")
	item2.AssertCalled(t, "Update")
}

func Test_Item_Update_DecreasesSellInByOne(t *testing.T) {
	sellin := 2
	item := &Item{"", sellin, 2}
	item.Update()

	assert.Equal(t, sellin-1, item.SellIn())
}

func Test_Item_Update_DecreasesQualityByOneBeforePassingSellInDate(t *testing.T) {
	sellin := 2
	quality := 2
	item := &Item{"", sellin, quality}
	item.Update()

	assert.Equal(t, quality-1, item.Quality())
}

func Test_Item_Update_DecreasesQualityByTwoAfterPassingSellInDate(t *testing.T) {
	sellin := -1
	quality := 3
	item := &Item{"", sellin, quality}
	item.Update()

	assert.Equal(t, quality-2, item.Quality())
}

func Test_Item_Update_QualityNeverBecomeNegative(t *testing.T) {
	itemBeforeSellIn := &Item{"", 1, 0}
	itemBeforeSellIn.Update()

	assert.Equal(t, 0, itemBeforeSellIn.Quality())

	itemAfterSellIn := &Item{"", -1, 1}
	itemAfterSellIn.Update()

	assert.Equal(t, 0, itemAfterSellIn.Quality())
}

func Test_AgedBrie_Update_IncreasesQualityByOneBeforeSellInDate(t *testing.T) {
	quality := 1
	agedBrie := NewAgedBrie("Aged Brie", 1, quality)
	agedBrie.Update()

	assert.Equal(t, quality+1, agedBrie.Quality())
}

func Test_AgedBrie_Update_IncreasesQualityByTwoAfterSellInDate(t *testing.T) {
	quality := 1
	agedBrie := NewAgedBrie("Aged Brie", -1, quality)
	agedBrie.Update()

	assert.Equal(t, quality+2, agedBrie.Quality())
}

func Test_AgedBrid_Update_QualityNeverBecomeMoreThan50(t *testing.T) {
	agedBrieBeforeSellIn := NewAgedBrie("Aged Brie", 1, MaxQuality)
	agedBrieBeforeSellIn.Update()

	assert.Equal(t, MaxQuality, agedBrieBeforeSellIn.Quality())

	agedBrieAfterSellIn := NewAgedBrie("Aged Brie", -1, MaxQuality-1)
	agedBrieAfterSellIn.Update()

	assert.Equal(t, MaxQuality, agedBrieAfterSellIn.Quality())
}

func Test_Sulfuras_Update_NeverChangeSellInAndQuality(t *testing.T) {
	sulfuras := NewSulfuras("Sulfuras, Hand of Ragnaros")
	sulfuras.Update()

	assert.Equal(t, 0, sulfuras.SellIn())
	assert.Equal(t, QualitySulfuras, sulfuras.Quality())
}

func Test_BackStagePasses_Update_IncreasesQualityByOneWhenMoreThan10Days(t *testing.T) {
	quality := 1

	bsPasses := &Item{"Backstage passes to a TAFKAL80ETC concert", 11, quality}
	bsPasses.Update()

	assert.Equal(t, quality+1, bsPasses.Quality())
}

func Test_BackStagePasses_Update_IncreasesQualityByTwoWhen10DaysOrLess(t *testing.T) {
	quality := 1

	for i := 6; i <= 10; i++ {
		bsPasses := &Item{"Backstage passes to a TAFKAL80ETC concert", i, quality}
		bsPasses.Update()
		assert.Equal(t, quality+2, bsPasses.Quality())
	}
}

func Test_BackStagePasses_Update_IncreasesQualityByThreeWhen5DaysOrLess(t *testing.T) {
	quality := 1

	for i := 1; i <= 5; i++ {
		bsPasses := &Item{"Backstage passes to a TAFKAL80ETC concert", i, quality}
		bsPasses.Update()
		assert.Equal(t, quality+3, bsPasses.Quality())
	}
}

func Test_BackStagePasses_Update_DropsQualityToZeroAfterTheSellInDate(t *testing.T) {
	quality := 10
	bsPasses := &Item{"Backstage passes to a TAFKAL80ETC concert", 0, quality}
	bsPasses.Update()

	assert.Equal(t, 0, bsPasses.Quality())
}

func Test_ConjuredItem_Update_DecreasesQualityTwiceFast(t *testing.T) {
	quality := 5
	conjuredItem := NewConjuredItem("Conjured Mana Cake", 1, quality)

	expected := quality - 2
	conjuredItem.Update()
	assert.Equal(t, expected, conjuredItem.Quality())

	expected -= 4
	conjuredItem.Update()
	assert.Equal(t, expected, conjuredItem.Quality())
}
