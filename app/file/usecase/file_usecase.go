package usecase

import (
	"context"
	"log"
	error2 "pixstall-file/domain/error"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/acl"
	model2 "pixstall-file/domain/file/model"
	image_processing "pixstall-file/domain/image/image-processing"
)

type fileUseCase struct {
	fileRepo file.Repo
	fileAclRepo acl.Repo
	imageProcessingRepo image_processing.Repo
}

func NewFileUseCase(fileRepo file.Repo, fileAclRepo acl.Repo, imageProcessingRepo image_processing.Repo) file.UseCase {
	return fileUseCase{
		fileRepo: fileRepo,
		fileAclRepo: fileAclRepo,
		imageProcessingRepo: imageProcessingRepo,
	}
}

func (f fileUseCase) SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType, name string) (*string, error) {

	files, err := NewFileFactory(fileType, f.imageProcessingRepo).getFiles(fileData, fileType, name)
	if err != nil || len(*files) <= 0 {
		return nil, error2.UnknownError
	}
	dFiles := *files
	paths, err := f.fileRepo.SaveFiles(ctx, dFiles)
	if err != nil {
		return nil, error2.UnknownError
	}
	log.Printf("SaveFile into paths:%v\n", paths)

	return &dFiles[0].RawPath, nil
}

func (f fileUseCase) IsAccessible(ctx context.Context, userID *string, prefix string, ext string, fullPath string) (*bool, error) {
	result := true
	return &result, nil
}