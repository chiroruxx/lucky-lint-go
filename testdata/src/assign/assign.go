package assign

func lllllllllllllll() (int, int) {
	a := 1 // want `naming 'a'\(stroke count 2\) is not lucky`
	aaaaaaal := 1

	// allow ok variable
	_, ok := a, aaaaaaal
	return ok, aaaaaaal
}
