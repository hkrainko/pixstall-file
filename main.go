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

	// Gin
	r := gin.Default()
	//userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	imgGroup := r.Group("/pri-img")
	{
		ctrl := InitImageController(db, awsS3)
		imgGroup.GET("/:imgType/:yyyy/:mm/:dd/:imgName", ctrl.GetImage)

		//imgGroup.GET("/:id", func(c *gin.Context) {
		//
		//	out, err := awsS3.GetObject(&s3.GetObjectInput{
		//		Bucket:                     aws.String(BucketName),
		//		Key:                        aws.String(c.Request.RequestURI),
		//	})
		//	if err != nil {
		//		if aerr, ok := err.(awserr.Error); ok {
		//			switch aerr.Code() {
		//			case s3.ErrCodeNoSuchKey:
		//				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
		//			default:
		//				fmt.Println(aerr.Error())
		//			}
		//		} else {
		//			// Print the error, cast err to awserr.Error to get the Code and
		//			// Message from an error.
		//			fmt.Println(err.Error())
		//		}
		//		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		//		return
		//	}
		//	b, err := ioutil.ReadAll(out.Body)
		//	defer out.Body.Close()
		//	if err != nil {
		//		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		//		return
		//	}
		//	c.Data(http.StatusOK, *out.ContentType, b)
		//	return




			//remote, err := url.Parse("https://pixstall-store-dev.s3.ap-east-1.amazonaws.com")
			//if err != nil {
			//	panic(err)
			//}
			//
			//proxy := httputil.NewSingleHostReverseProxy(remote)
			////Define the director func
			////This is a good place to log, for example
			//proxy.Director = func(req *http.Request) {
			//	req.Header = c.Request.Header
			//	req.Host = remote.Host
			//	req.URL.Scheme = remote.Scheme
			//	req.URL.Host = remote.Host
			//	//req.URL.Path = c.Param("proxyPath")
			//	req.RequestURI = c.Request.RequestURI
			//}
			//
			//proxy.ServeHTTP(c.Writer, c.Request)
		//})


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

	completionFileGroup := r.Group("/file")
	{
		ctrl := InitFileController(db, awsS3)
		completionFileGroup.GET("/completion/*action", ctrl.GetCompletionFile)
	}

	err = r.Run(":9007")
	print(err)
}