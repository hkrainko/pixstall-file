package model

type File struct {
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
	FileTypeArtwork          = "artwork"
	FileTypeRoof             = "roof"
	FileTypeOpenCommission   = "open-commission"
	FileTypeProfile          = "profile"
)
