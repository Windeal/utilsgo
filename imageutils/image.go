package imageutils

// GetImageTypeByPrefix : 根据图片前缀编码获取图片类型
func GetImageTypeByPrefix(b *[]byte) MIME {
	// JPEG (jpg)，文件头：FFD8FF
	if len(*b) > 2 && (*b)[0] == 0xFF && (*b)[1] == 0xD8 && (*b)[2] == 0xFF {
		return ImageJPEG
	}

	// PNG (png)，文件头：89504E47
	if len(*b) > 4 && (*b)[0] == 0x89 && (*b)[1] == 0x50 && (*b)[2] == 0x4E && (*b)[3] == 0x47 {
		return ImagePNG
	}

	// GIF (gif)，文件头：47494638
	if len(*b) > 4 && (*b)[0] == 0x47 && (*b)[1] == 0x49 && (*b)[2] == 0x46 && (*b)[3] == 0x38 {
		return ImageGIF
	}

	// TIFF (tif)，文件头：49492A00
	if len(*b) > 4 && (*b)[0] == 0x49 && (*b)[1] == 0x49 && (*b)[2] == 0x2A && (*b)[3] == 0x00 {
		return ImageTIF
	}

	// Windows Bitmap (bmp)，文件头：424D
	if len(*b) > 2 && (*b)[0] == 0x42 && (*b)[1] == 0x4D {
		return ImageBMP
	}

	return ImageUnknown
}
