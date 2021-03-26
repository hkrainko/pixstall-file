package file

import (
	"context"
	model2 "pixstall-file/domain/file/model"
)

type UseCase interface {
	SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, ext string) (*string, error)
	GetFile(ctx context.Context, userID string, path string) (*model2.File, error)
}