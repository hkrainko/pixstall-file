package model

type File struct {
	Data        []byte
	Name        string
	ContentType string
	Volume      int64
}

type FileType string

const (
	FileTypeMessage = "msg"
	FileTypeCompletion = "completion"
	FileTypeCommissionRefImg = "commission-ref-img"
	FileTypeArtwork = "artwork"
	FileTypeRoof = "roof"
	FileTypeOpenCommission = "open-commission"
	FileTypeProfile = "profile"
)