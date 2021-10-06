package main

import (
	"testing"

	"github.com/glycerine/goconvey/convey"
)

func TestLRUCache_Put(t *testing.T) {

	convey.Convey("t1", func() {
		tests := [][]int{
			{6, 14},
			{3, 1},
			{3},
			{10, 11},
			{8},
			{2, 14},
			{1},
			{5},
			{4},
			{11, 4},
			{12, 24},
			{5, 18},
			{13},
			{7, 23},
			{8},
			{12},
			{3, 27},
			{2, 12},
			{5},
			{2, 9},
			{13, 4},
			{8, 18},
			{1, 7},
			{6},
		}
		cache := NewLRUCache(10)
		for _, tt := range tests {
			if len(tt) == 2 {
				cache.Put(tt[0], tt[1])
			} else {
				cache.Get(tt[0])
			}
		}
		convey.So(cache.Get(6), convey.ShouldAlmostEqual, -1)
	})

}
