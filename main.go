package main

import (
	"flag"
	//"github.com/dorako321/imgc"
	"io/ioutil"
	"path/filepath"
	"fmt"
)

func main(){
	flag.Parse()
	//var fileType = flag.String("type", "jpeg", "select file type (default is jpeg)")

	// 処理対象のディレクトリを取得
	directory := flag.Arg(0)
	fileList := getFileList(directory);

	fmt.Println(fileList)

}

func getFileList(dir string) []string {
	// 対象のディレクトリを読み込み
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir(){
			paths = append(paths, getFileList(filepath.Join(dir, file.Name()))...);
		} else {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}
	return paths
}
