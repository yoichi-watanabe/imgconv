package imgconv

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

//Convert はオプションで指定された拡張子へ画像ファイルを変換する関数
func Convert() error {
	var (
		targetdir = flag.String("t", "./", "Message")
		src       = flag.String("s", "jpg", "Message")
		dest      = flag.String("d", "png", "Message")
	)

	// オプションのパース
	flag.Parse()

	// 対象拡張子のファイルを抽出
	err := filepath.Walk(*targetdir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fmt.Sprintf(".%s", *src) {
			// 処理
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

			fmt.Println(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *dest) + " を出力します")
			out, err := os.Create(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *dest))
			defer out.Close()

			// 変換後拡張子に応じたエンコード
			switch *dest {
			case "jpg", "jpeg":
				jpeg.Encode(out, img, nil)

			case "png":
				png.Encode(out, img)
			}

			fmt.Println(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *dest) + " を出力しました")

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
