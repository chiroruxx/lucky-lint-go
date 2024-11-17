package assign

func lllllllllllllll() (int, int) {
	a := 1 // want `naming 'a' is not lucky`
	aaaaaaal := 1

	// allow ok variable
	_, ok := a, aaaaaaal
	return ok, aaaaaaal
}
