# go-upload-api-minio

## Minio setup
To create a MinIO container with persistent storage, you need to map local persistent directories from the host OS to virtual config. 
To do this, run the below commands
```sh
mkdir -p ~/minio/data

docker run \
-p 9000:9000 \
-p 9001:9001 \
--name minio1 \
-v ~/minio/data:/data \
-e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
-e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
quay.io/minio/minio server /data --console-address ":9001"
```

