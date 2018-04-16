package main

import (
    "flag"
	"fmt"
	"os"
)

// エンドポイント
func main() {
	os.Exit(Run(os.Args))
}


// execute関数
func Run(args []string) int {
	// 処理
	var (
		dir   = flag.String("dir", "./", "Message")
		from   = flag.String("from", "jpg", "Message")
		to   = flag.String("to", "png", "Message")
    )

    flag.Parse()

	fmt.Printf("param -dir -> %s\n", *dir)
	fmt.Printf("from -from -> %s\n", *from)
	fmt.Printf("to -to -> %s\n", *to)
	
	return 0
}