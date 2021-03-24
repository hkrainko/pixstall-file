package usecase

import (
	"context"
	"fmt"
	error2 "pixstall-file/domain/error"
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

	files, err := getFiles(fileData, fileType)
	if err != nil || len(*files) <= 0 {
		return nil, error2.UnknownError
	}
	dFiles := *files
	paths, err := f.fileRepo.SaveFiles(ctx, dFiles)
	if err != nil {
		return nil, error2.UnknownError
	}
	fmt.Printf("SaveFile into paths:%v", paths)

	return &dFiles[0].RawPath, nil
}

func (f fileUseCase) GetFile(ctx context.Context, userID string, path string) (*model2.File, error) {
	panic("implement me")
}

func getFiles(fileData []byte, fileType model2.FileType) (*[]model2.File, error) {
	switch fileType {
	case model2.FileTypeMessage:



	}


}