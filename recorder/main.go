package main

import (
	"log"
)

func main() {
	store := NewStore()
	watcher := NewWatcher("media")
	watcher.Start()

	done := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case ev := <-watcher.EventStream:
				log.Print(ev.Path())
				store.UploadFile(ev.Path())
			}
		}
	}()

	<-done
	log.Println("Done...")
}
