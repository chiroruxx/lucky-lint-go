package imprt

import (
	"fmt"
	a "os" // want `naming 'a' is not lucky`
	"strings"
)

func lllllllllllllll() (string, []string) {
	if strings.HasPrefix("a", "b") {
		return "", nil
	}
	return fmt.Sprintf("a"), a.Args
}
