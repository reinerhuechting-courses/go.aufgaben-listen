package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine aufsteigende Kette von mindestens
// drei Zahlen enthält.
// Im Gegensatz zu `ContainsChain()` müssen die Zahlen
// hier nicht direkt aufeinanderfolgen.
func ContainsChain2(list []int) bool {
	if len(list) < 3 {
		return false
	}
	counter := 1
	for pos, el := range list[:len(list)-1] {
		if el < list[pos+1] {
			counter++
		} else {
			counter = 1
		}
		if counter >= 3 {
			return true
		}
	}
	return false
}
