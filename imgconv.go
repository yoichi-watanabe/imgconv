package imgconv

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Convert() {
	// 処理
	var (
		dir  = flag.String("dir", "./", "Message")
		src  = flag.String("src", ".jpg", "Message")
		dest = flag.String("dest", "./testdata/dest.png", "Message")
	)

	// オプションのパース
	flag.Parse()

	path := "./testdata/sample1.jpg"
	fmt.Println(filepath.Ext(path))

	// 対象拡張子のファイルを抽出
	if filepath.Ext(path) == *src {
		fmt.Println(path)
		fmt.Println("MATCH")

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

		// エンコード
		out, err := os.Create(*dest)
		defer out.Close()

		png.Encode(out, img)
	}

	fmt.Printf("param -dir -> %s\n", *dir)
	fmt.Printf("src -src -> %s\n", *src)
	fmt.Printf("dits -dest -> %s\n", *dest)

}
