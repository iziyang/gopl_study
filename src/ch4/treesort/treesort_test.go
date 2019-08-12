package treesort_test

import (
	"ch4/treesort"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {  // 如果数据没有排序
		t.Errorf("not sorted: %v", data)
	}
}
