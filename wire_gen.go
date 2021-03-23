// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
	http2 "pixstall-file/app/file/delivery/http"
	"pixstall-file/app/file/repo/aws-s3"
	usecase2 "pixstall-file/app/file/usecase"
	"pixstall-file/app/image/delivery/http"
	"pixstall-file/app/image/usecase"
)

// Injectors from wire.go:

func InitImageController(db *mongo.Database, awsS3 *s3.S3) http.ImageController {
	repo := aws_s3.NewAWSS3FileRepository(awsS3)
	useCase := usecase.NewImageUseCase(repo)
	imageController := http.NewImageController(useCase)
	return imageController
}

func InitFileController(db *mongo.Database, awsS3 *s3.S3) http2.FileController {
	repo := aws_s3.NewAWSS3FileRepository(awsS3)
	useCase := usecase2.NewFileUseCase(repo)
	fileController := http2.NewFileController(useCase)
	return fileController
}
