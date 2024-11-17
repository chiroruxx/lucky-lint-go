package decl

func c() { // want `naming 'c' is not lucky`
	return
}

func llllllllllllllc() {
	return
}

type sssssssssssssss int

func (sssssssssssssss sssssssssssssss) a() { // want `naming 'a' is not lucky`
	return
}
func (sssssssssssssss sssssssssssssss) ssssssssssssssl() {
	return
}

// allow receiver name
func (s sssssssssssssss) ssssssssssssssc() {
}

// allow init func
func init() {
}
