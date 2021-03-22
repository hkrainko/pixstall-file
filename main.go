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
	"net/http"
	"time"
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
		},
		//Profile:                 "default", //[default], use [prod], [uat]
		//SharedConfigState:       session.SharedConfigEnable,
	}))
	_ = s3.New(sess)

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
	_ = dbClient.Database("pixstall-file")


	r := gin.Default()

	//userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	apiGroup := r.Group("/file")

	imageGroup := apiGroup.Group("/img")
	{
		imageGroup.GET("/:id", func(ctx *gin.Context) {
			print("Get")
			ctx.AbortWithStatus(http.StatusOK)
		})


		//ctrl := InitArtistController(db, awsS3)
		//// Artist
		//artistGroup.GET("/:id", ctrl.GetArtist)
		//artistGroup.GET("/:id/details", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetArtistDetails)
		//artistGroup.PATCH("/:id", userIDExtractor.ExtractPayloadsFromJWT, ctrl.UpdateArtist)
		//// Open Commission
		//artistGroup.GET("/:id/open-commissions", ctrl.GetOpenCommissionsForArtist)
		//artistGroup.GET("/:id/open-commissions/details", userIDExtractor.ExtractPayloadsFromJWT, ctrl.GetOpenCommissionsDetailsForArtist)
		//artistGroup.POST("/:id/open-commissions", userIDExtractor.ExtractPayloadsFromJWT, ctrl.AddOpenCommissionForArtist)
	}

	err = r.Run(":9006")
	print(err)
}