package iterations

// Repeate returns the recieved letter two times concated
func Repeate(letter string, letterLength int) string {
	var repeatedLetter string

	for i := 0; i < letterLength; i++ {
		repeatedLetter += letter
	}

	return repeatedLetter
}