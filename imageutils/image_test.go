package imageutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetImageTypeByPrefix(t *testing.T) {
	type args struct {
		b        *[]byte
		filename string
	}
	tests := []struct {
		name string
		args args
		want MIME
	}{
		// TODO: Add test cases.
		{
			name: "./testfile/bmp.bmp",
			args: args{
				b: nil,
			},
			want: ImageBMP,
		},
		{
			name: "./testfile/gif.gif",
			args: args{
				b: nil,
			},
			want: ImageGIF,
		},
		{
			name: "./testfile/jpeg.jpg",
			args: args{
				b: nil,
			},
			want: ImageJPEG,
		},
		{
			name: "./testfile/png.png",
			args: args{
				b: nil,
			},
			want: ImagePNG,
		},
		{
			name: "./testfile/tiff.tif",
			args: args{
				b: nil,
			},
			want: ImageTIF,
		},
	}
	for _, tt := range tests {

		fp, err := os.OpenFile(tt.name, os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
		if err != nil {
			// 如果有错误返回错误内容
			t.Errorf(err.Error())
		}
		//defer fp.Close()

		b, err := ioutil.ReadAll(fp)
		if err != nil {
			t.Errorf(err.Error())
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := GetImageTypeByPrefix(&b); got != tt.want {
				t.Errorf("GetImageTypeBySuffix() = %v, want %v", got, tt.want)
			}
		})

	}
}
