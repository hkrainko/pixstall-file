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

type ImageSizeType string

const (
	ImageSizeTypeSmall = "xs"
	ImageSizeTypeMiddle = "md"
	ImageSizeTypeLarge = "lg"
	ImageSizeTypeRaw = "raw"
)