package imprt

import (
	"fmt"
	a "os" // want `naming 'a'\(stroke count 2\) is not lucky`
	"strings"
)

func lllllllllllllll() (string, []string) {
	if strings.HasPrefix("a", "b") {
		return "", nil
	}
	return fmt.Sprintf("a"), a.Args
}
