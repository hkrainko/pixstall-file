package usecase

import (
	"context"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
	"pixstall-file/domain/image/model"
)

type fileUseCase struct {
	fileRepo file.Repo
}

func NewFileUseCase(fileRepo file.Repo) file.UseCase {
	return fileUseCase{
		fileRepo: fileRepo,
	}
}

func (f fileUseCase) SaveImage(ctx context.Context, image model.Image, dir string) (*string, error) {
	panic("implement me")
}

func (f fileUseCase) SaveImages(ctx context.Context, images []model.Image, dir string) ([]string, error) {
	panic("implement me")
}

func (f fileUseCase) SaveFile(ctx context.Context, file model2.File, dir string) (*string, error) {
	panic("implement me")
}

func (f fileUseCase) SaveFiles(ctx context.Context, images []model2.File, dir string) ([]string, error) {
	panic("implement me")
}

func (f fileUseCase) GetImage(ctx context.Context, path string) (*model.Image, error) {
	panic("implement me")
}

func (f fileUseCase) GetFile(ctx context.Context, path string) (*model2.File, error) {
	panic("implement me")
}