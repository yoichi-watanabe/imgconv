// Package main はimgconvコマンドのエントリーポイント
package main

import (
	"os"

	// 自作パッケージ
	"github.com/yoichi-watanabe/imgconv"
)

// エンドポイント
func main() {
	os.Exit(Run(os.Args))
}

// Run はmainパッケージのexecute関数
func Run(args []string) int {
	imgconv.Convert(args)
	return 0
}
