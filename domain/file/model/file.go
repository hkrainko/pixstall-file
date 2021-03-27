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
	FileTypeMessage              = "msg"
	FileTypeCompletion           = "completion"
	FileTypeCommissionRef        = "commission-ref"
	FileTypeCommissionProofCopy  = "commission-proof-copy"
	FileTypeArtworkHidden        = "artwork-hidden"
	FileTypeArtwork              = "artwork"
	FileTypeRoof                 = "roof"
	FileTypeOpenCommission       = "open-commission"
	FileTypeOpenCommissionHidden = "open-commission-hidden"
	FileTypeProfile              = "profile"
)

func (f FileType) GetFileDir() string {
	result := ""
	switch f {
	case FileTypeMessage:
		result = "pri-img/msg/"
		break
	case FileTypeCompletion:
		result = "pri-file/completion/"
		break
	case FileTypeCommissionRef:
		result = "pri-img/commission-ref/"
		break
	case FileTypeCommissionProofCopy:
		result = "pri-img/commission-proof-copy/"
		break
	case FileTypeArtworkHidden:
		result = "pri-img/artwork/"
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
	case FileTypeOpenCommissionHidden:
		result = "img/open-commission-hidden"
		break
	case FileTypeProfile:
		result = "img/profile"
		break
	}

	return result + time.Now().Format("2006/01/02") + "/"
}
