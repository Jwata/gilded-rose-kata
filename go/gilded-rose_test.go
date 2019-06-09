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

func Test_Item_Update_DecrementsSellInAndQuality(t *testing.T) {
	item := &Item{"Test Item", 10, 10}
	item.Update()

	assert.Equal(t, item.SellIn(), 9)
	assert.Equal(t, item.Quality(), 9)
}
