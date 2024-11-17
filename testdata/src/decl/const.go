package decl

const a = 1 // want `naming 'a' is not lucky`

var aaaaaaal = ""

var (
	b        = 1 // want `naming 'b' is not lucky`
	bbbbbbbl = 1
)
