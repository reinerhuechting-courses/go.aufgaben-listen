package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste x mehr als einmal enthält.
func ContainsMultiple(list []int, x int) bool {
	counter := 0
	for _, el := range list {
		if el == x {
			counter++
			if counter > 1 {
				return true
			}
		}
	}
	return false
}
