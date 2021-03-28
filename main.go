package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"pixstall-file/app/middleware"
	"pixstall-file/proto"
	"time"
)

const (
	BucketName = "pixstall-store-dev"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//AWS s3
	awsAccessKey := "AKIA5BWICLKRWX6ARSEF"
	awsSecret := "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"
	token := ""
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(endpoints.ApEast1RegionID),
			CredentialsChainVerboseErrors: aws.Bool(true),
			Credentials:                   creds,
			//DisableRestProtocolURICleaning: aws.Bool(true),
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	awsS3 := s3.New(sess)

	//Mongo
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer cancel()
	defer func() {
		if err = dbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	db := dbClient.Database("pixstall-file")

	////gRPC
	lis, err := net.Listen("tcp", ":50052")
	s := grpc.NewServer()
	proto.RegisterFileServiceServer(s, InitFileStoreService(db, awsS3))
	go s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Gin
	r := gin.Default()
	userIDExtractor := middleware.NewJWTPayloadsExtractor("userId")

	imgGroup := r.Group("/img")
	{
		ctrl := InitImageController(db, awsS3)
		imgGroup.GET("/:imgType/:yyyy/:mm/:dd/:imgName", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetImage)
	}

	fileGroup := r.Group("/file")
	{
		ctrl := InitFileController(db, awsS3)
		fileGroup.GET("/:fileType/:yyyy/:mm/:dd/:fileName", userIDExtractor.ExtractPayloadsFromJWTInHeader, ctrl.GetFile)
	}

	err = r.Run(":9007")
	print(err)
}