package main

import (
	"log"
	"regexp"
)

var (
	regexTempl = regexp.MustCompile(`^.*\.(m3u8|ts)$`)
)

func main() {
	store := NewStore()
	watcher := NewWatcher("/app/media")
	watcher.Start()

	done := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case ev := <-watcher.EventStream:
				if regexTempl.MatchString(ev.Path()) {
					log.Print(ev.Path())
					go store.UploadFile(ev.Path())
				}
			}
		}
	}()

	<-done
	log.Println("Done...")
}
