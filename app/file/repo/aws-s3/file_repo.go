package aws_s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
)

type awsS3FileRepository struct {
	s3 *s3.S3
}

const (
	BucketName = "pixstall-store-dev"
)

func NewAWSS3FileRepository(s3 *s3.S3) file.Repo {
	return &awsS3FileRepository{
		s3: s3,
	}
}

func (a awsS3FileRepository) SaveFile(ctx context.Context, file model2.File) (*string, error) {
	// convert buffer to reader
	reader := bytes.NewReader(file.Data)

	// use it in `PutObjectInput`
	_, err := a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(BucketName),
		Key:         aws.String(file.Path),
		Body:        reader,
		ContentType: aws.String(file.ContentType),
		ACL:         aws.String("public-read"), //profile should be public accessible
	})

	if err != nil {
		return nil, err
	}
	return &file.Path, nil
}

func (a awsS3FileRepository) SaveFiles(ctx context.Context, files []model2.File) ([]string, error) {
	var resultPaths []string
	for _, dFile := range files {
		// convert buffer to reader
		reader := bytes.NewReader(dFile.Data)

		// use it in `PutObjectInput`
		_, err := a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket: aws.String(BucketName),
			Key:    aws.String(dFile.Path),
			Body:   reader,
			ContentType: aws.String("image"),
			ACL: aws.String("public-read"),  //profile should be public accessible
		})
		if err == nil {
			resultPaths = append(resultPaths, dFile.Path)
		}
	}
	return resultPaths, nil
}

func (a awsS3FileRepository) GetFile(ctx context.Context, path string) (*model2.File, error) {
	panic("implement me")
}
