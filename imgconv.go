package imgconv

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Convert() {
	// 処理
	var (
		//dir  = flag.String("dir", "./", "Message")
		src  = flag.String("s", "jpg", "Message")
		dest = flag.String("d", "png", "Message")
	)

	// オプションのパース
	flag.Parse()

	path := "./testdata/sample1.jpg"
	fmt.Println(filepath.Ext(path))

	// 対象拡張子のファイルを抽出
	if filepath.Ext(path) == fmt.Sprintf(".%s", *src) {

		// 対象ファイルオープン
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			fmt.Println(err)
		}

		// デコード
		img, _, err := image.Decode(file)
		if err != nil {
			fmt.Println(err)
		}

		out, err := os.Create(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *dest))
		defer out.Close()

		// 変換後拡張子に応じたエンコード
		switch *dest {
		case "jpg", "jpeg":
			jpeg.Encode(out, img, nil)

		case "png":
			png.Encode(out, img)

		case "gif":
			gif.Encode(out, img, nil)
		}
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
