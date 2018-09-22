package calg

import "fmt"

func colorable(s string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", s)
}
