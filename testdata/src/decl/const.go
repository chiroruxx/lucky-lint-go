package decl

const a = 1 // want `naming 'a'\(stroke count 2\) is not lucky`

var aaaaaaal = ""

var (
	b        = 1 // want `naming 'b'\(stroke count 2\) is not lucky`
	bbbbbbbl = 1
)
