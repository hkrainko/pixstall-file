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
	FileTypeMessage             = "msg"
	FileTypeCompletion          = "completion"
	FileTypeCommissionRef       = "commission-ref"
	FileTypeCommissionProofCopy = "commission-proof-copy"
	FileTypeArtwork             = "artwork"
	FileTypeRoof                = "roof"
	FileTypeOpenCommission      = "open-commission"
	FileTypeProfile             = "profile"
)

func (f FileType) IsValid() bool {
	switch f {
	case FileTypeMessage,
		FileTypeCompletion,
		FileTypeCommissionRef,
		FileTypeCommissionProofCopy,
		FileTypeArtwork,
		FileTypeRoof,
		FileTypeOpenCommission,
		FileTypeProfile:
		return true
	default:
		return false
	}
}

func (f FileType) GetFileDir() string {
	result := ""
	switch f {
	case FileTypeMessage:
		result = "img/msg/"
		break
	case FileTypeCompletion:
		result = "file/completion/"
		break
	case FileTypeCommissionRef:
		result = "img/commission-ref/"
		break
	case FileTypeCommissionProofCopy:
		result = "img/commission-proof-copy/"
		break
	case FileTypeArtwork:
		result = "img/artwork/"
		break
	case FileTypeRoof:
		result = "img/roof/"
		break
	case FileTypeOpenCommission:
		result = "img/open-commission"
		break
	case FileTypeProfile:
		result = "img/profile"
		break
	}

	return result + time.Now().Format("2006/01/02") + "/"
}
