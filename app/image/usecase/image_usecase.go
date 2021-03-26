package usecase

import (
	"context"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/acl"
	"pixstall-file/domain/image"
	"pixstall-file/domain/image/model"
)

type imageUseCase struct {
	fileRepo    file.Repo
	fileACLRepo acl.Repo
}

func NewImageUseCase(fileRepo file.Repo, fileACLRepo acl.Repo) image.UseCase {
	return imageUseCase{
		fileRepo:    fileRepo,
		fileACLRepo: fileACLRepo,
	}
}

func (i imageUseCase) SaveImage(ctx context.Context, image model.Image, dir string) (*string, error) {
	panic("implement me")
}

func (i imageUseCase) SaveImages(ctx context.Context, images []model.Image, dir string) ([]string, error) {
	panic("implement me")
}

func (i imageUseCase) IsAccessible(ctx context.Context, userID *string, prefix string, ext string, fullPath string) (*bool, error) {
	result := true
	return &result, nil
}
