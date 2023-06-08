package main

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
)

func main() {

	UploadFile("test-bucket", "testfile")

}

// UploadFile - uploads file on google storage bucket
func UploadFile(fileName string, bucketName string) (bool, error) {

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	bucket := client.Bucket(bucketName)
       
	file := bucket.Object(bucketName+"/"+fileName)

	// Read the file from disk
	data, err := ioutil.ReadFile(fileName)
    	if err != nil {
        	log.Errorf("Error while reading source file %v",err)
		return false , err
    	}

    	// Upload the file to the bucket
    	_, err = file.Upload(context.Background(), bytes.NewReader(data))
    	if err != nil {
        	log.Errorf("Error while uploading file %v",err)
		return false , err
    	}

	log.Infof("File uploaded: %v", fileName)

return true,_
}
