package main

import (
	"fmt"
	"log"

	"github.com/rjeczalik/notify"
)

type Watcher struct {
	Path        string
	EventStream chan notify.EventInfo
}

func NewWatcher(path string) *Watcher {
	return &Watcher{
		Path:        fmt.Sprintf("%s/...", path),
		EventStream: make(chan notify.EventInfo, 1),
	}
}

func (w *Watcher) Start() {
	if err := notify.Watch(w.Path, w.EventStream, notify.Create); err != nil {
		log.Fatal(err)
	}
}
