package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"pixstall-file/domain/file"
	"pixstall-file/domain/file/model"
)

type FileController struct {
	useCase file.UseCase
}

func NewFileController(useCase file.UseCase) FileController {
	return FileController{
		useCase: useCase,
	}
}

func (i FileController)GetFile(ctx *gin.Context) {
	tokenUserID := ctx.GetString("userId")
	fileType := ctx.Param("fileType")
	if fileType == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	fileName := ctx.Param("fileName")
	if fileName == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ext := filepath.Ext(fileName)
	if ext == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	prefix := fileName[0:len(fileName) - len(ext)]
	if prefix == "" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	accessible, err := i.useCase.IsAccessible(ctx, &tokenUserID, model.FileType(fileType), prefix)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !*accessible {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	i.proxy(ctx)
}

func (i FileController)proxy(ctx *gin.Context) {
	remote, err := url.Parse("https://d1kazwr29qo5il.cloudfront.net")
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
