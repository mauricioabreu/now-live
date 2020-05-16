package main

import (
	"log"

	"github.com/minio/minio-go"
)

type Store struct {
	Client     *minio.Client
	bucketName string
}

func NewStore(bucketName string) *Store {
	endpoint := "minio_proxy:80"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}

	return &Store{Client: client, bucketName: bucketName}
}

func (s *Store) UploadFile(filePath string) {
	bucketName := "video"
	resource := Resource{File: filePath}

	n, err := s.Client.FPutObject(bucketName, resource.ObjectName("/app/media"), filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", filePath, n)
}

func (s *Store) CreateBucket(bucketName string) error {
	err := s.Client.MakeBucket(bucketName, "us-east-1")
	if err != nil {
		exists, errExists := s.Client.BucketExists(bucketName)
		if errExists == nil && exists {
			return nil
		}
		return errExists
	}
	return err
}
