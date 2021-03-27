package file

import (
	"context"
	model2 "pixstall-file/domain/file/model"
)

type UseCase interface {
	SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, name string) (*string, error)
	IsAccessible(ctx context.Context, userID *string, prefix string, ext string, fullPath string) (*bool, error)
}