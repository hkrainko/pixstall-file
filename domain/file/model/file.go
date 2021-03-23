package model

import "mime/multipart"

type File struct {
	multipart.File
	Name        string
	ContentType string
	Volume      int64
}
