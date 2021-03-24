package usecase

import (
	"context"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
)

type fileUseCase struct {
	fileRepo file.Repo
}

func NewFileUseCase(fileRepo file.Repo) file.UseCase {
	return fileUseCase{
		fileRepo: fileRepo,
	}
}

func (f fileUseCase) SaveFile(ctx context.Context, fileData []byte, fileType model2.FileType) (*string, error) {



	f.fileRepo.SaveFile()
}

func (f fileUseCase) GetFile(ctx context.Context, userID string, path string) (*model2.File, error) {
	panic("implement me")
}