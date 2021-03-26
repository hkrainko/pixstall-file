//+build wireinject

package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	acl_repo "pixstall-file/app/file/acl/repo/mongo"
	file_deli "pixstall-file/app/file/delivery/http"
	file_repo "pixstall-file/app/file/repo/aws-s3"
	file_usecase "pixstall-file/app/file/usecase"
	image_deli "pixstall-file/app/image/delivery/http"
	image_processing_repo "pixstall-file/app/image/image-processing/repo/imaging"
	image_usecase "pixstall-file/app/image/usecase"
)

func InitImageController(db *mongo.Database, awsS3 *s3.S3) image_deli.ImageController {
	wire.Build(
		image_deli.NewImageController,
		image_usecase.NewImageUseCase,
		file_repo.NewAWSS3FileRepository,
		acl_repo.NewMongoFileAclRepo,
	)
	return image_deli.ImageController{}
}

func InitFileController(db *mongo.Database, awsS3 *s3.S3) file_deli.FileController {
	wire.Build(
		file_deli.NewFileController,
		file_usecase.NewFileUseCase,
		file_repo.NewAWSS3FileRepository,
		image_processing_repo.NewImagingImageProcessingRepo,
	)

	return file_deli.FileController{}
}
