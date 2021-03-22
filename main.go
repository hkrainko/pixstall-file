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
	"net/http/httputil"
	"net/url"
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

	// Proxy
	//proxy := goproxy.NewProxyHttpServer()
	//proxy.Verbose = true
	//proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	//proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("amazonaws.com$"))).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	//	ctx.Logf("%v", "We can see what APIs are being called!")
	//	return req, ctx.Resp
	//})
	//proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	//	ctx.Logf("%v", "We can modify some data coming back!")
	//	return resp
	//})
	//
	//proxy.OnRequest().DoFunc(
	//	func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	//		fmt.Printf("req:%v", req.PostForm)
	//		print("%v", awsS3.Endpoint)
	//		return req, nil
	//	},
	//)
	//
	//err = http.ListenAndServe(":9006", proxy)

	// Gin
	r := gin.Default()
	//userIDExtractor := middleware.NewJWTPayloadsExtractor([]string{"userId"})

	artworkGroup := r.Group("/artwork")

	//imageGroup := fileGroup.Group("/img")
	{
		artworkGroup.GET("/:id", func(c *gin.Context) {
			print("Get")
			//ctx.AbortWithStatus(http.StatusOK)
			//director := func(req *http.Request) {
			//	req = ctx.Request
			//	req.URL.Scheme = "http"
			//	req.URL.Host = "https://pixstall-store-dev.s3.ap-east-1.amazonaws.com/"
			//	req.Host = "https://pixstall-store-dev.s3.ap-east-1.amazonaws.com/"
			//	req.Header["my-header"] = []string{ctx.Request.Header.Get("my-header")}
			//	// Golang camelcases headers
			//	delete(req.Header, "My-Header")
			//}
			//proxy := &httputil.ReverseProxy{Director: director}
			//proxy.ServeHTTP(ctx.Writer, ctx.Request)

			remote, err := url.Parse("https://pixstall-store-dev.s3.ap-east-1.amazonaws.com")
			if err != nil {
				panic(err)
			}

			proxy := httputil.NewSingleHostReverseProxy(remote)
			//Define the director func
			//This is a good place to log, for example
			proxy.Director = func(req *http.Request) {
				req.Header = c.Request.Header
				req.Host = remote.Host
				req.URL.Scheme = remote.Scheme
				req.URL.Host = remote.Host
				//req.URL.Path = c.Param("proxyPath")
				req.RequestURI = c.Request.RequestURI
			}

			proxy.ServeHTTP(c.Writer, c.Request)
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

	err = r.Run(":9007")
	print(err)
}