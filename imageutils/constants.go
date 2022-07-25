package imageutils

type MIME string

const (
	ImageUnknown MIME = "unknown"
	ImageJPEG    MIME = "image/jpeg"
	ImagePNG     MIME = "image/png"
	ImageGIF     MIME = "image/gif"
	ImageTIF     MIME = "image/tiff"
	ImageBMP     MIME = "image/bmp"
)
