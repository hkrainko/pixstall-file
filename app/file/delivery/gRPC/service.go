package gRPC

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/model"
	"pixstall-file/proto"
	"time"
)

type FileService struct {
	fileUseCase  file.UseCase
	proto.UnimplementedFileServiceServer
}

func NewFileService(fileUseCase file.UseCase) FileService {
	return FileService{
		fileUseCase:  fileUseCase,
	}
}

func (f FileService) SaveFile(stream proto.FileService_SaveFileServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	buf := bytes.NewBuffer([]byte{})

	startTime := time.Now()
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	metaData := req.GetMetaData()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			recFile := buf.Bytes()

			path, err := f.fileUseCase.SaveFile(ctx, &recFile, model.FileType(metaData.FileType), metaData.Ext)
			if err != nil {
				return err
			}
			fmt.Printf("SaveFile used: %v(s)", endTime.Sub(startTime))
			return stream.SendAndClose(&proto.SaveFileResponse{
				Path: *path,
			})
		}
		if err != nil {
			return err
		}
		buf.Write(req.GetFile())
	}
}
