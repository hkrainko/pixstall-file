package image

import (
	"context"
	model2 "pixstall-file/domain/file/model"
	"pixstall-file/domain/image/model"
)

type UseCase interface {
	SaveImage(ctx context.Context, image model.Image, dir string) (*string, error)
	SaveImages(ctx context.Context, images []model.Image, dir string) ([]string, error)
	IsPublic(ctx context.Context, imgType model2.FileType, prefix string) (*bool, error)
	GetImage(ctx context.Context, userID *string, prefix string, ext string, fullPath string) (*model.Image, error)
}
