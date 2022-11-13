package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert die Position von x in der Liste.
// Liefert -1, falls x nicht enthalten ist.
func Find(list []int, x int) int {
	for pos, el := range list {
		if el == x {
			return pos
		}
	}
	return -1
}
