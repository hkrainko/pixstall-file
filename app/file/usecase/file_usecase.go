package usecase

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	"image"
	"image/jpeg"
	error2 "pixstall-file/domain/error"
	"pixstall-file/domain/file"
	model2 "pixstall-file/domain/file/model"
	image_processing "pixstall-file/domain/image/image-processing"
	"pixstall-file/domain/image/model"
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

func (f fileUseCase) SaveFile(ctx context.Context, fileData *[]byte, fileType model2.FileType) (*string, error) {

	files, err := f.getFiles(fileData, fileType)
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

func (f fileUseCase) getFiles(fileData *[]byte, fileType model2.FileType) (*[]model2.File, error) {
	var files []model2.File

	switch fileType {
	case model2.FileTypeMessage:
		if b, err := f.imageProcessingRepo.ResizeToJpegByte(*fileData, 800, 0); err != nil {
			return nil, err
		} else {
			file := model2.File{
				Data:        b,
				FileType:    model2.FileTypeMessage,
				ContentType: "image/jpeg",
				IsPublic:    false,
				Path:        fileType.GetFileDir() + model.ImageScaleMiddle.PathSuffix(),
				RawPath:     "",
			}

		}




	}


}

func (f fileUseCase) getImagePath(fileType model2.FileType, scale string)