package image

import (
	"context"
	"pixstall-file/domain/image/model"
)

type UseCase interface {
	SaveImage(ctx context.Context, image model.Image, dir string) (*string, error)
	SaveImages(ctx context.Context, images []model.Image, dir string) ([]string, error)
	IsAccessible(ctx context.Context, userID *string, prefix string, ext string, fullPath string) (*bool, error)
}
