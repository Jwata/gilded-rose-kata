package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GildedRose(t *testing.T) {
	item1 := &Item{"Test Item 1", 1, 1}
	item2 := &Item{"Test Item 2", 2, 2}
	testItems := []*Item{item1, item2}

	GildedRose(testItems)

	assert.Equal(t, testItems[0].sellIn, 0)
	assert.Equal(t, testItems[0].quality, 0)
	assert.Equal(t, testItems[1].sellIn, 1)
	assert.Equal(t, testItems[1].quality, 1)
}
