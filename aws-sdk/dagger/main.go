// module for interact with the AWS SDK

package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsSdk struct{}

// Upload a file to a AWS S3 bucket
// Example usage:
// dagger call upload-bucket \
// --region="us-west-2"  \
// --bucket="mybucket" \
// --file=myfile \
// --access-key env:AWS_ACCESS_KEY --secret-key env:AWS_SECRET_KEY
func (m *AwsSdk) UploadBucket(
	ctx context.Context,
	// AWS access key
	accessKey *Secret,
	// AWS secret key
	secretKey *Secret,
	// AWS S3 region
	region string,
	// AWS S3 bucket name
	bucket string,
	// File to upload
	file *File,
	// Timeout for the operation
	// +optional
	// +default=60
	timeout int,

) (string, error) {

	accessKeyPlainText, err := accessKey.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	secretKeyPlainText, err := secretKey.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyPlainText, secretKeyPlainText, ""),
	}))

	svc := s3.New(sess)
	var cancelFn func()

	t := time.Duration(timeout) * time.Second
	if t > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, t)
	}

	if cancelFn != nil {
		defer cancelFn()
	}

	fileName, err := file.Name(ctx)
	if err != nil {
		return "", err
	}
	file.Export(ctx, fileName)

	readFile, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:             aws.String(bucket),
		Body:               bytes.NewReader(readFile),
		Key:                aws.String(fileName),
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
			return "", err
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
			return "", err
		}
	}

	return "Succesfully upload file " + fileName + " to " + bucket + " s3 bucket", nil
}
