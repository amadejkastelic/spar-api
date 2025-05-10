package utils

import (
	"io"
)

func CloseQuetly(v any) {
	if d, ok := v.(io.Closer); ok {
		_ = d.Close()
	}
}
