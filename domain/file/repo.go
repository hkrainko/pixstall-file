package file

import (
	"context"
	model2 "pixstall-file/domain/file/model"
	"pixstall-file/domain/image/model"
)

type Repo interface {
	SaveImage(ctx context.Context, image model.Image, dir string) (*string, error)
	SaveImages(ctx context.Context, images []model.Image, dir string) ([]string, error)
	SaveFile(ctx context.Context, file model2.File, dir string) (*string, error)
	SaveFiles(ctx context.Context, images []model2.File, dir string) ([]string, error)
	GetImage(ctx context.Context, path string) (*model.Image, error)
	GetFile(ctx context.Context, path string) (*model2.File, error)
}
