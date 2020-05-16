package main

import "strings"

type Resource struct {
	Path string
}

func (r *Resource) ObjectName() string {
	return strings.Replace(r.Path, "/live/", "", 1)
}
