package aws_s3

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
	"pixstall-file/domain/image/model"
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

func (a awsS3FileRepository) SaveImage(ctx context.Context, image model.Image, path string) (*string, error) {
	panic("implement me")
}

func (a awsS3FileRepository) SaveImages(ctx context.Context, pathImages []model.Image, path string) ([]string, error) {
	panic("implement me")
}

func (a awsS3FileRepository) SaveFile(ctx context.Context, pathFile model2.File, path string) (*string, error) {
	panic("implement me")
}

func (a awsS3FileRepository) SaveFiles(ctx context.Context, pathImages []model2.File, path string) ([]string, error) {
	panic("implement me")
}

func (a awsS3FileRepository) GetImage(ctx context.Context, path string) (*model.Image, error) {
	panic("implement me")
}

func (a awsS3FileRepository) GetFile(ctx context.Context, path string) (*model2.File, error) {
	panic("implement me")
}