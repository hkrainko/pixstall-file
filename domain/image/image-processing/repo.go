package image_processing

type Repo interface {
	ResizeToJpegByte(byte []byte, width int, height int) ([]byte, error)
}