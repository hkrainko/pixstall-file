package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"pixstall-file/domain/image"
)

type ImageController struct {
	useCase image.UseCase
}

func NewImageController(useCase image.UseCase) ImageController {
	return ImageController{
		useCase: useCase,
	}
}

func (i ImageController)GetImage(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	imgName, exist := ctx.GetQuery("imgName")
	if !exist {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ext := filepath.Ext(imgName)
	if ext == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	prefix := imgName[0:len(imgName) - len(ext)]
	if prefix == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	img, err := i.useCase.GetImage(ctx, &tokenUserID, prefix, ext, ctx.FullPath())
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Data(http.StatusOK, img.ContentType, img.Data)
}

func (i ImageController)proxy(ctx *gin.Context) {
	remote, err := url.Parse("https://pixstall-store-dev.s3.ap-east-1.amazonaws.com")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		//req.URL.Path = c.Param("proxyPath")
		req.RequestURI = ctx.Request.RequestURI
	}

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}