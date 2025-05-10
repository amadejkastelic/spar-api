package format

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatFloat(f float64, sep string) string {
	if f == 0 {
		return fmt.Sprintf("0%s00", sep)
	}
	return strings.Replace(strconv.FormatFloat(f, 'f', 2, 64), ".", sep, 1)
}
