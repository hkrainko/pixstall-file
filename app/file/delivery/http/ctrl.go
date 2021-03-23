package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-file/domain/file"
)

type FileController struct {
	useCase file.UseCase
}

func NewFileController(useCase file.UseCase) FileController {
	return FileController{
		useCase: useCase,
	}
}

func (i FileController)GetCompletionFile(ctx *gin.Context) {

}

