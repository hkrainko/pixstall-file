package gRPC

import (
	"context"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/model"
	"pixstall-file/proto"
)

type FileService struct {
	fileUseCase  file.UseCase
}

func NewFileService(fileUseCase file.UseCase) FileService {
	return FileService{
		fileUseCase:  fileUseCase,
	}
}

func (f FileService) SaveFile(ctx context.Context, request *proto.SaveFileRequest) (*proto.SaveFileResponse, error) {
	path, err := f.fileUseCase.SaveFile(ctx, request.File, model.FileType(request.FileType))
	if err != nil {
		return nil, err
	}
	return &proto.SaveFileResponse{
		Path: *path,
	}, nil
}
