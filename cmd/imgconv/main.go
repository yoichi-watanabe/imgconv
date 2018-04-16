package main

import (
	"os"

	"github.com/yoichi-watanabe/imgconv"
)

// エンドポイント
func main() {
	os.Exit(Run(os.Args))
}

// execute関数
func Run(args []string) int {
	imgconv.Convert()
	return 0
}
