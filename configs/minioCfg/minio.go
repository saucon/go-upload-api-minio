package minioCfg

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/saucon/go-upload-api-minio/configs/env"
	"log"
)

func NewMinio(config env.ServerConfig) *minio.Client {
	endpoint := config.MinioConfig.Host
	accessKeyID := config.MinioConfig.AccessKey
	secretAccessKey := config.MinioConfig.SecretAccessKey
	useSSL := config.MinioConfig.UseSsl

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	return minioClient
}
