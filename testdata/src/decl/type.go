package decl

type s int    // want `naming 's'\(stroke count 1\) is not lucky`
type sa = int // want `naming 'sa'\(stroke count 3\) is not lucky`
type ssssssssssssssl int
