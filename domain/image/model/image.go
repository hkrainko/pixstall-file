package model

import "pixstall-file/domain/file/model"

type Image struct {
	model.File
	Size
}

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}

type ImageScale string

const (
	ImageScaleSmall ImageScale = "xs"
	ImageScaleMiddle ImageScale = "md"
	ImageScaleLarge ImageScale = "lg"
	ImageScaleRaw ImageScale = "raw"
)

func (i ImageScale) PathSuffix() string {
	if i == ImageScaleRaw {
		return ""
	}
	return "_" + string(i)
}