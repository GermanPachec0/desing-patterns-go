package main

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvacuationAlgo(lru)

	cache.add("d", "4")
	fifo := &Fifo{}
	cache.setEvacuationAlgo(fifo)
	cache.add("e", "5")
	cache.add("f", "6")
}
