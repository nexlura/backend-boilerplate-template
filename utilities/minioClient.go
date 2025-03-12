package utilities

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

func getMinioClient() *minio.Client {
	endpoint := "minio.minio-dev:9000"                            // LOCAL -> "localhost:9000"                                  //PROD -> "minio.minio-dev:9000"
	accessKeyID := "WYS0K71Pdz5QLzoe9aJl"                         // LOCAL -> "iaPxveQEkbyPSlhxBUJm"                         //PROD-> "WYS0K71Pdz5QLzoe9aJl"
	secretAccessKey := "TQWVb1nr4aG1RTBLkjnQPgA3wfYvhNHkvbvt0APs" // LOCAL -> "iEzdHljhjStFxKO9tHBDooHsNihG5uM3zm73QRZk" //PROD-> "TQWVb1nr4aG1RTBLkjnQPgA3wfYvhNHkvbvt0APs"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln("minio client err: ", err.Error())
		return nil
	}

	//log.Printf("minioClient connected:: %v", minioClient) // minioClient is now set up

	return minioClient
}

func MinioUpload(file *multipart.FileHeader, bucket string) (string, error) {
	c := getMinioClient()

	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	// generate unique id for the upload file
	id := GenerateUUID()
	fileExt := filepath.Ext(file.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(file.Filename), filepath.Ext(file.Filename))
	now := time.Now()
	fileName := id + "_" + originalFileName + "_" + fmt.Sprintf("%v", now.Unix()) + fileExt

	// trigger minio to  upload
	uploadInfo, err := c.PutObject(ctx, bucket, fileName, reader, file.Size, minio.PutObjectOptions{
		ContentType: file.Header["Content-Type"][0],
		UserMetadata: map[string]string{
			"mediaIdentity": id,
			"mediaType":     bucket,
		},
	})
	if err != nil {
		return "", err
	}

	var filePath string
	// Prepare the file path
	if uploadInfo.Key != "" {
		filePath = fmt.Sprintf("%s/%s/%s", "https://media.propati.xyz", bucket, fileName)
	}

	return filePath, nil
}

func MinioDownload(bucket, id string) (*minio.Object, minio.ObjectInfo, error) {
	c := getMinioClient()

	object, err := c.GetObject(ctx, bucket, id, minio.GetObjectOptions{})
	if err != nil {
		return nil, minio.ObjectInfo{}, err
	}

	defer object.Close()
	stats, err := object.Stat()
	if err != nil {
		return nil, minio.ObjectInfo{}, err
	}

	return object, stats, nil
}

func MinioDelete(bucket, id string) error {
	c := getMinioClient()

	err := c.RemoveObject(ctx, bucket, id, minio.RemoveObjectOptions{})

	if err != nil {
		return err
	}

	return nil
}

func MinioList(bucket string) ([]interface{}, error) {
	c := getMinioClient()

	ct, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Query the list of objects
	objectsChannel := c.ListObjects(ct, bucket, minio.ListObjectsOptions{})
	var objects []interface{}

	// Iterate ove the object channel
	for object := range objectsChannel {
		// Throw error if any
		if object.Err != nil {
			return nil, object.Err
		}

		// Append the object
		objects = append(objects, object)
	}

	return objects, nil
}

func MinioRemoveMedia(media string, bucket string) error {
	// get the media link from the fetch result
	mediaId := GetMediaIdFromLink(media, bucket)

	// delete the key from minio
	if deleteErr := MinioDelete(bucket, mediaId); deleteErr != nil {
		return deleteErr
	}

	return nil
}

func MinioRemoveMultipleMedia(bucket string, media []string) error {
	// delete the current uploaded files
	if len(media) > 0 {
		// loop over images and delete
		for _, item := range media {
			if deleteErr := MinioRemoveMedia(item, bucket); deleteErr != nil {
				return deleteErr
			}
		}
	}

	return nil
}
