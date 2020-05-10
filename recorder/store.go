package main

import (
	"log"
	"path/filepath"

	"github.com/minio/minio-go"
)

type Store struct {
	Client *minio.Client
}

func NewStore() *Store {
	endpoint := "minio1:9000"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", client)
	return &Store{Client: client}
}

func (s *Store) UploadFile(filePath string) {
	bucketName := "video"
	location := "us-east-1"

	err := s.Client.MakeBucket(bucketName, location)
	if err != nil {
		exists, errExists := s.Client.BucketExists(bucketName)
		if errExists == nil && exists {
			log.Printf("Bucket %s already exists\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Bucket successfully created %s\n", bucketName)
	}

	objectName := filepath.Base(filePath)
	contentType := "application/octet-stream"

	n, err := s.Client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
