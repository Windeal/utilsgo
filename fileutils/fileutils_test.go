package fileutils

import (
	"testing"
)

func TestReplaceStringsInFile(t *testing.T) {
	type args struct {
		filePath string
		origin   string
		target   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "测试下文本替换",
			args: args{
				filePath: "./replace.txt",
				origin:   "that",
				target:   "this",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReplaceStringsInFile(tt.args.filePath, tt.args.origin, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("ReplaceStringsInFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
