package file

import (
	"context"
	model2 "pixstall-file/domain/file/model"
)

type UseCase interface {
	SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, name string, owner string, acl []string) (*string, error)
	IsAccessible(ctx context.Context, accessUserID *string, fileType model2.FileType, prefix string) (*bool, error)
}