package imgconv

import "testing"

// getFileNameWithoutExt のテスト
func TestGetFileNameWithoutExt(t *testing.T) {
	if getFileNameWithoutExt("aaa.png") != "aaa" {
		t.Fatal("getFileNameWithoutExt(\"aaa.png\") shold be \"aaa\", but doesn't match")
	}
}
