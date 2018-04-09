package main

import (
	"fmt"
	"sort"
)

type SortList []map[string]interface{}

func (self SortList) Len() int {
	return len(self)
}

func (self SortList) Less(i, j int) bool {
	var iItem = self[i]
	var jItem = self[j]
	var iSortVal = iItem["sortVal"].(int)
	var jSortVal = jItem["sortVal"].(int)
	return iSortVal > jSortVal
}

func (self SortList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func testSort() {
	var dataList = make([]map[string]interface{}, 0)
	var item1 = make(map[string]interface{})
	item1["sortVal"] = 1
	var item2 = make(map[string]interface{})
	item2["sortVal"] = 2
	var item3 = make(map[string]interface{})
	item3["sortVal"] = 3
	dataList = append(dataList, item1, item2, item3)

	//排序前
	fmt.Println(dataList)
	sort.Sort(SortList(dataList))
	//排序后
	fmt.Println(dataList)
}
