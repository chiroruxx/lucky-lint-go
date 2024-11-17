package decl

func c() { // want `naming 'c'\(stroke count 1\) is not lucky`
	return
}

func llllllllllllllc() {
	return
}

type sssssssssssssss int

func (sssssssssssssss sssssssssssssss) a() { // want `naming 'a'\(stroke count 2\) is not lucky`
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
