package gRPC

import (
	"context"
	"io"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/model"
	"pixstall-file/proto"
	"time"
)

type FileService struct {
	fileUseCase  file.UseCase
}

func NewFileService(fileUseCase file.UseCase) FileService {
	return FileService{
		fileUseCase:  fileUseCase,
	}
}

func (f FileService) SaveFile(server proto.FileService_SaveFileServer) error {
	var pointCount, featureCount, distance int32
	var lastPoint *pb.Point
	startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		if lastPoint != nil {
			distance += calcDistance(lastPoint, point)
		}
		lastPoint = point
	}
}

func (f FileService) mustEmbedUnimplementedFileServiceServer() {
	panic("implement me")
}

//func (f FileService) SaveFile(ctx context.Context, request *proto.SaveFileRequest) (*proto.SaveFileResponse, error) {
//	path, err := f.fileUseCase.SaveFile(ctx, request.File, model.FileType(request.FileType))
//	if err != nil {
//		return nil, err
//	}
//	return &proto.SaveFileResponse{
//		Path: *path,
//	}, nil
//}
