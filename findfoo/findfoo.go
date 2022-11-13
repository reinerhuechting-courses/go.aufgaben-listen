package findfoo

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste die Zahl x enthält.
func Contains(list []int, x int) bool {
	for _, el := range list {
		if el == x {
			return true
		}
	}
	return false
}

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

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Falls es eine Stelle gibt, an der x zwei Mal
// hintereinander vorkommt, liefert die Funktion
// diese Stelle.
// Liefert -1, falls die Situation nicht auftritt.
func FindPair(list []int, x int) int {
	if len(list) < 2 {
		return -1
	}

	for pos, el := range list[:len(list)-1] {
		if el == x && list[pos+1] == x {
			return pos
		}
	}
	return -1
}
