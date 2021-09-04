package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type awss3 struct {
	S3SESSION         *s3.S3
	BUCKET_NAME       string
	REGION            string
	CREDENTIAL_KEY    string
	CREDENTIAL_SECRET string
	OBJECT_URI        string
}

type AmazonS3 interface {
	UploadObjectbyFilePath(string) (*s3.PutObjectOutput, error)
	UploadByteObject([]byte, string) (*s3.PutObjectOutput, error)
	GetObject(string) (*s3.GetObjectOutput, error)
	ListObjects() (*s3.ListObjectsV2Output, error)
	DeleteObject(string) (*s3.DeleteObjectOutput, error)
	GetURI() string
}

func NewS3Session() AmazonS3 {
	return &awss3{
		BUCKET_NAME:       os.Getenv("AWS_BUCKET_NAME"),
		REGION:            os.Getenv("AWS_REGION"),
		CREDENTIAL_KEY:    os.Getenv("AWS_CREDENTIAL_KEY"),
		CREDENTIAL_SECRET: os.Getenv("AWS_CREDENTIAL_SECRET"),
		S3SESSION: s3.New(
			session.Must(
				session.NewSession(
					&aws.Config{
						Region: aws.String(os.Getenv("AWS_REGION")),
						Credentials: credentials.NewStaticCredentials(
							os.Getenv("AWS_CREDENTIAL_KEY"),
							os.Getenv("AWS_CREDENTIAL_SECRET"),
							"",
						),
					},
				),
			),
		),
	}
}

func (as *awss3) UploadObjectbyFilePath(filepath string) (res *s3.PutObjectOutput, err error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return res, err
	}

	f_bytes := bytes.NewReader(f)
	fileType := http.DetectContentType(f)

	as.OBJECT_URI = os.Getenv("AWS_URI_PREFIX") + url.QueryEscape(strings.Split(filepath, "/")[len(strings.Split(filepath, "/"))-1])
	return as.S3SESSION.PutObject(&s3.PutObjectInput{
		Body:        f_bytes,
		Bucket:      aws.String(as.BUCKET_NAME),
		Key:         aws.String(strings.Split(filepath, "/")[len(strings.Split(filepath, "/"))-1]),
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		ContentType: aws.String(fileType),
	})
}

func (as *awss3) UploadByteObject(f []byte, name string) (res *s3.PutObjectOutput, err error) {
	f_bytes := bytes.NewReader(f)
	fileType := http.DetectContentType(f)

	as.OBJECT_URI = os.Getenv("AWS_URI_PREFIX") + name
	return as.S3SESSION.PutObject(&s3.PutObjectInput{
		Body:        f_bytes,
		Bucket:      aws.String(as.BUCKET_NAME),
		Key:         aws.String(name),
		ACL:         aws.String(s3.BucketCannedACLPublicRead),
		ContentType: aws.String(fileType),
	})
}

func (as *awss3) GetObject(filename string) (res *s3.GetObjectOutput, err error) {
	return as.S3SESSION.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(as.BUCKET_NAME),
		Key:    aws.String(filename),
	})
}

func (as *awss3) ListObjects() (res *s3.ListObjectsV2Output, err error) {
	return as.S3SESSION.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(as.BUCKET_NAME),
	})
}

func (as *awss3) DeleteObject(filename string) (res *s3.DeleteObjectOutput, err error) {
	return as.S3SESSION.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(as.BUCKET_NAME),
		Key:    aws.String(filename),
	})
}

func (as *awss3) GetURI() string {
	return as.OBJECT_URI
}
