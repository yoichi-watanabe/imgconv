// Package imgconv は画像ファイル形式変換処理を実行するパッケージ
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

// Cmdoption はコマンド実行時に設定されたオプションを持つ構造体
type Cmdoption struct {
	targetdir *string
	src       *string
	dest      *string
}

//Convert はオプションで指定された拡張子へ画像ファイルを変換する関数
func Convert(args []string) error {

	c := Cmdoption{
		flag.String("t", "./", "Message"),
		flag.String("s", "jpg", "Message"),
		flag.String("d", "png", "Message")}

	// オプションのパース
	flag.Parse()

	// ログ出力
	fmt.Println(fmt.Sprintf("%s配下の拡張子を変換します。 %s → %s", *c.targetdir, *c.src, *c.dest))

	// 対象拡張子のファイルを抽出
	err := filepath.Walk(*c.targetdir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == fmt.Sprintf(".%s", *c.src) {
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

			fmt.Println(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *c.dest) + " を出力します")
			out, err := os.Create(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *c.dest))
			defer out.Close()

			// 変換後拡張子に応じたエンコード
			switch *c.dest {
			case "jpg", "jpeg":
				jpeg.Encode(out, img, nil)

			case "png":
				png.Encode(out, img)

			case "gif":
				gif.Encode(out, img, nil)
			}

			fmt.Println(getFileNameWithoutExt(path) + fmt.Sprintf(".%s", *c.dest) + " を出力しました")

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// getFileNameWithoutExt は拡張子を外したファイル名を返却する関数
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
