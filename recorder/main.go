package main

import (
	"log"

	"github.com/minio/minio-go/v6"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	storageClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", storageClient)

	// Watch files
	watcher := NewWatcher("media")
	watcher.Start()

	done := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case ev := <-watcher.EventStream:
				log.Print(ev.Path())
			}
		}
	}()

	<-done
	log.Println("Done...")
}
