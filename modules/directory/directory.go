package directoryUtil

import (
	"io/ioutil"
	"path/filepath"
)

func GetFileList(path string, ext string) []string {
	// 対象のディレクトリを読み込み
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir(){
			paths = append(paths, GetFileList(filepath.Join(path, file.Name()), ext)...);
		} else if "." + ext == filepath.Ext(file.Name()) {
			paths = append(paths, filepath.Join(path, file.Name()))
		}
	}
	return paths
}