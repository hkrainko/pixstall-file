//+build wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"pixstall-file/app/file/delivery/http"
	file_repo "pixstall-file/app/file/repo/aws-s3"
	file_usecase "pixstall-file/app/file/usecase"
)

func InitFileController(db *mongo.Database, awsS3 *s3.S3) http.FileController {
	wire.Build(
		http.NewFileController,
		file_usecase.NewFileUseCase,
		file_repo.NewAWSS3FileRepository,
	)

	return http.FileController{}

}
