package converter

import (
	"os"
	"image/jpeg"
	"image/png"
	"path/filepath"
)

type FilePathStruct struct {
	Base     string
	Filename string
	Ext      string
}

func NewFile(path string) *FilePathStruct {
	str := &FilePathStruct{
		Base:     filepath.Dir(path),
		Filename: filepath.Base(path[:len(path)-len(filepath.Ext(path))]),
		Ext:      filepath.Ext(path),
	}
	return str
}

func ToConvert(path string, toExt string) error {
	filepath := NewFile(path);
	outputFile := filepath.Base + "/" + filepath.Filename + "." + toExt

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	png.Encode(out, img)
	return nil
}
