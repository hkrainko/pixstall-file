package usecase

import (
	"context"
	"fmt"
	error2 "pixstall-file/domain/error"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
	image_processing "pixstall-file/domain/image/image-processing"
)

type fileUseCase struct {
	fileRepo file.Repo
	imageProcessingRepo image_processing.Repo
}

func NewFileUseCase(fileRepo file.Repo, imageProcessingRepo image_processing.Repo) file.UseCase {
	return fileUseCase{
		fileRepo: fileRepo,
		imageProcessingRepo: imageProcessingRepo,
	}
}

func (f fileUseCase) SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, ext string) (*string, error) {

	files, err := NewFileFactory(fileType, f.imageProcessingRepo).getFiles(fileData, fileType, ext)
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