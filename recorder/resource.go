package main

import (
	"path/filepath"
	"strings"
)

type Resource struct {
	File string
}

func (r *Resource) ObjectName(root string) string {
	path := strings.Replace(r.File, root, "", 1)
	if path[0] == '/' {
		return path[1:]
	}
	return filepath.Dir(path)
}

func (r *Resource) Path() string {
	return filepath.Base(r.File)
}
