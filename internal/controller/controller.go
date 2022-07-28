package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"net/http"
)

type UploadHandler struct {
	mnoClient *minio.Client
}

func NewUploadHandler(a *minio.Client) UploadHandler {
	return UploadHandler{
		mnoClient: a,
	}
}

func (uh *UploadHandler) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename

	// TODO: Rename bucketName with the one you have been created in minio
	uploadInfo, err := uh.mnoClient.PutObject(c.Request.Context(), "trxvalas-partnerid", filename, file, header.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	c.JSON(200, gin.H{"responseCode": "2000000", "responseMessage": "Upload Success"})
}
