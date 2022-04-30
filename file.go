package util

import (
	"path"
	"strings"
)

func Extension(name string) string {
	ext := strings.ToLower(path.Ext(name))
	if ext != "" {
		return ext[1:]
	}

	return ext
}
