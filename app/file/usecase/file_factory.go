package usecase

import (
	"github.com/google/uuid"
	"net/http"
	model2 "pixstall-file/domain/file/model"
	image_processing "pixstall-file/domain/image/image-processing"
	"pixstall-file/domain/image/model"
	"strings"
)

type FileFactory struct {
	fileType            model2.FileType
	imageProcessingRepo image_processing.Repo
}

func NewFileFactory(fileType model2.FileType, imageProcessingRepo image_processing.Repo) FileFactory {
	return FileFactory{
		fileType:            fileType,
		imageProcessingRepo: imageProcessingRepo,
	}
}

func (f FileFactory) getFiles(fileData *[]byte, fileType model2.FileType, name string) (*[]model2.File, error) {
	newName, err := f.getRandomName()
	if err != nil {
		return nil, err
	}
	contentType, err := getFileContentType(fileData)
	if err != nil {
		return nil, err
	}
	dir := f.fileType.GetFileDir()
	rawPath := dir + *newName + f.getFileExt(name)
	var files []model2.File
	for scale, width := range f.getScaleWidthMap(fileType) {
		if b, err := f.imageProcessingRepo.ResizeToJpegByte(*fileData, width, 0); err != nil {
			return nil, err
		} else {
			files = append(files, model2.File{
				ID:          *newName,
				Data:        b,
				FileType:    fileType,
				ContentType: "image/jpeg",
				Path:        dir + *newName + scale.PathSuffix() + ".jpg",
				RawPath:     rawPath,
			})
		}
	}
	files = append(files, model2.File{
		ID:          *newName,
		Data:        *fileData,
		FileType:    fileType,
		ContentType: contentType,
		Path:        rawPath,
		RawPath:     rawPath,
	})
	return &files, nil
}

func (f FileFactory) getScaleWidthMap(fileType model2.FileType) map[model.ImageScale]int {
	result := make(map[model.ImageScale]int)
	switch fileType {
	case model2.FileTypeMessage:
		result[model.ImageScaleMiddle] = 200
		break
	case model2.FileTypeCompletion:
		break
	case model2.FileTypeCommissionRef:
		result[model.ImageScaleMiddle] = 200
		break
	case model2.FileTypeCommissionProofCopy:
		result[model.ImageScaleMiddle] = 200
		break
	case model2.FileTypeArtwork:
		result[model.ImageScaleLarge] = 600
		result[model.ImageScaleMiddle] = 200
		result[model.ImageScaleSmall] = 50
		break
	case model2.FileTypeRoof:
		result[model.ImageScaleLarge] = 1200
		break
	case model2.FileTypeOpenCommission:
		result[model.ImageScaleMiddle] = 300
		result[model.ImageScaleSmall] = 100
		break
	case model2.FileTypeProfile:
		result[model.ImageScaleMiddle] = 200
		result[model.ImageScaleSmall] = 50
		break
	}
	return result
}

func (f FileFactory) getFileExt(name string) string {
	ss := strings.Split(name, ".")
	if len(ss) <= 1 {
		return ""
	}
	return "." + ss[len(ss)-1]
}

func (f FileFactory) getRandomName() (*string, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	result := strings.ReplaceAll(newUUID.String(), "-", "")
	return &result, nil
}

func getFileContentType(fileByte *[]byte) (string, error) {
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(*fileByte)

	return contentType, nil
}
