package main

import (
	"flag"
	"./modules/converter"
	"./modules/directory"
	"os"
)


var (
	fileFrom = flag.String("to", "jpg", "select file type (default is jpg)")
	fileTo = flag.String("from", "png", "select file type (default is png)")
)

func main(){
	flag.Parse()

	// 処理対象のディレクトリを取得
	directory := flag.Arg(0)
	fileList := directoryUtil.GetFileList(directory, *fileFrom);

	// オプションに沿って処理を実行
	for _, file := range fileList {
		err := converter.ToConvert(file, *fileTo)
		if err != nil {
			os.Exit(1)
		}
	}

	os.Exit(0)
}


