package gRPC

import (
	"bytes"
	"context"
	"errors"
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
	fileType, err := f.domainFileTypeFormGRPC(metaData.FileType)
	if err != nil {
		return err
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			recFile := buf.Bytes()


			path, err := f.fileUseCase.SaveFile(ctx, &recFile, fileType, metaData.Ext)
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

func (f FileService) domainFileTypeFormGRPC(gFileType proto.MetaData_FileType) (model.FileType, error) {
	switch gFileType {
	case proto.MetaData_Message:
		return model.FileTypeMessage, nil
	case proto.MetaData_Completion:
		return model.FileTypeCompletion, nil
	case proto.MetaData_CommissionRef:
		return model.FileTypeCommissionRef, nil
	case proto.MetaData_CommissionProofCopy:
		return model.FileTypeCommissionProofCopy, nil
	case proto.MetaData_ArtworkHidden:
		return model.FileTypeArtworkHidden, nil
	case proto.MetaData_Artwork:
		return model.FileTypeArtwork, nil
	case proto.MetaData_Roof:
		return model.FileTypeRoof, nil
	case proto.MetaData_OpenCommission:
		return model.FileTypeOpenCommission, nil
	case proto.MetaData_OpenCommissionHidden:
		return model.FileTypeOpenCommissionHidden, nil
	case proto.MetaData_Profile:
		return model.FileTypeProfile, nil
	default:
		return "", errors.New("not found")
	}
}