package model

import (
	"time"
)

type File struct {
	ID          string
	Data        []byte
	FileType    FileType
	ContentType string
	IsPublic    bool
	Path        string
	RawPath     string
}

type FileType string

const (
	FileTypeMessage          = "msg"
	FileTypeCompletion       = "completion"
	FileTypeCommissionRefImg = "commission-ref-img"
	FileTypeArtworkHidden    = "artwork-hidden"
	FileTypeArtwork          = "artwork"
	FileTypeRoof             = "roof"
	FileTypeOpenCommission   = "open-commission"
	FileTypeProfile          = "profile"
)

func (f FileType) GetFileDir() string {
	dateFormat := time.Now().Format("2006/01/02")
	result := ""
	switch f {
	case FileTypeMessage:
		result = "pri-img/msg/"
	case FileTypeCompletion:
		result = "pri-file/completion/"
	case FileTypeCommissionRefImg:
		result = "pri-img/commission-ref/"
	case FileTypeArtworkHidden:
		result = "pri-img/artwork/"
	case FileTypeArtwork:
		result = "img/artwork/"
	case FileTypeRoof:
		result = "img/roof/"
	case FileTypeOpenCommission:
		result = "img/open-commission"
	case FileTypeProfile:
		result = "img/profile"
	}

	return result + dateFormat + "/"
}
