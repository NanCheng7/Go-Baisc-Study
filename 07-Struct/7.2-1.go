package main

import "fmt"

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemid int) {
	b.items = append(b.items, itemid)
}
func main() {
	b := new(Bag)
	b.Insert(1001)
	b.Insert(666)
	fmt.Println(b.items)
}
