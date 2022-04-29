package fileutils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// IsExist 判断文件或者目录是否存在
// @param {string} path - 文件或者目录的路径
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// ReadFile 读取文件
func ReadFile(path string) ([]byte, error) {
	// 只支持文件读
	if !IsFile(path) {
		return nil, errors.New("file not exists")
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("read file failed")
	}

	return b, nil
}

// CopyFile 拷贝文件,
func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// 如果目标文件夹不存在，则创建
	dstPath := path.Dir(dst)
	if !IsExist(dstPath) {
		if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
			return err
		}
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// ReplaceStringsInFile 替换文件中的字符串
// @param {string} filePath - 文件路径
// @param {string} origin - 源字符串
// @param {string} target - 目标字符串字符串
func ReplaceStringsInFile(filePath string, origin, target string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	output := make([]byte, 0)

	reg := regexp.MustCompile(origin)
	targetByte := []byte(target)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		newLine := reg.ReplaceAll(line, targetByte)
		output = append(output, newLine...)
		output = append(output, byte('\n'))
	}

	err = ReWriteFile(filePath, output)
	if err != nil {
		return err
	}

	return err
}

// ReWriteFile 将新的内容写入文件，写入时会先清空文件的旧内容
// @param {string} filePath - 文件路径
// @param {string} output -  要写入文件的内容
func ReWriteFile(filePath string, content []byte) error {
	var f *os.File
	var err error
	if !IsExist(filePath) {
		dir := filepath.Dir(filePath)
		if !IsExist(dir) {
			_ = os.MkdirAll(dir, os.ModePerm)
		}
		f, err = os.Create(filePath)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
	}
	defer func() {
		_ = f.Close()
	}()

	writer := bufio.NewWriter(f)
	_, err = writer.Write(content)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// AppendToFile 向文件追加内容
// @param {string} content - 要追加到文件的内容
func AppendToFile(fileName string, content []byte) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = f.Close()
	}()

	// 查找文件末尾的偏移量
	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt(content, n)
	if err != nil {
		return err
	}
	return nil
}

// ListFiles 列出目录下的文件
func ListFiles(dirPath string) ([]string, error) {
	files := make([]string, 0)
	err := filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			if strings.HasPrefix(f.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil
}
