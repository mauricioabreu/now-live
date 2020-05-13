package main

import (
	"io"
	"log"

	"github.com/minio/minio-go"
)

type Store struct {
	Client *minio.Client
}

func NewStore() *Store {
	endpoint := "minio_proxy:80"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", client)
	return &Store{Client: client}
}

func (s *Store) GetObject(resource string) ([]byte, error) {
	reader, err := s.Client.GetObject("video", resource, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	stat, err := reader.Stat()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, stat.Size)
	_, err = reader.Read(buf)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	return buf, nil
}
