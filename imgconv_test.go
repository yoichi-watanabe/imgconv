package imgconv

import (
	"fmt"
	"strings"
	"testing"
)

var table = [...]string{
	"imgconv",
	"imgconv,-s,png,-d,jpg",
	"imgconv,-s,jpg,-d,gif",
}

func TestConvert(t *testing.T) {

	for i, test := range table {
		args := strings.Split(test, ",")

		fmt.Println("%i 回目のテストを実施します", i)
		fmt.Println(args)
		err := Convert(args)

		if err != nil {
			fmt.Println("処理中にエラーが発生しました。テスト失敗です。")
		}
	}

}

// getFileNameWithoutExt のテスト
func TestGetFileNameWithoutExt(t *testing.T) {
	if getFileNameWithoutExt("aaa.png") != "aaa" {
		t.Fatal("getFileNameWithoutExt(\"aaa.png\") shold be \"aaa\", but doesn't match")
	}
}
