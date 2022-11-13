package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine Kette von drei Zahlen enthält,
// deren Summe 42 ist.
func ContainsSum(list []int) bool {
	if len(list) < 3 {
		return false
	}
	counter := list[0] + list[1] + list[2]
	for pos, el := range list[3:] {
		if counter == 42 {
			return true
		}
		counter += el
		counter -= list[pos]
	}
	return counter == 42
}
