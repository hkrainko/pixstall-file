package acl

import (
	"context"
	"pixstall-file/domain/file/acl/model"
)

type Repo interface {
	AddFileACL(ctx context.Context, fileACL model.FileACL) (*string, error)
	GetFileACL(ctx context.Context, id string) (*model.FileACL, error)
	UpdateFileACL(ctx context.Context, id string, updater model.FileACLUpdater) error
}
