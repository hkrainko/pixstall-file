package imaging

import (
	"bufio"
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	_ "image/jpeg"
	image_processing "pixstall-file/domain/image/image-processing"
)

type imagingImageProcessingRepo struct {

}

func NewImagingImageProcessingRepo() image_processing.Repo {
	return &imagingImageProcessingRepo{

	}
}

func (i imagingImageProcessingRepo) ResizeToJpegByte(byte []byte, width int, height int) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(byte))
	if err != nil {
		return nil, err
	}

	dstImage := imaging.Resize(img, width, height, imaging.Lanczos)

	var b bytes.Buffer
	if err := imaging.Encode(bufio.NewWriter(&b), dstImage, imaging.JPEG); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
