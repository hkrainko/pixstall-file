package file

import (
	"context"
	model2 "pixstall-file/domain/file/model"
)

type Repo interface {
	SaveFile(ctx context.Context, file model2.File) (*string, error)
	SaveFiles(ctx context.Context, files []model2.File) ([]string, error)
	GetFile(ctx context.Context, path string) (*model2.File, error)
}
