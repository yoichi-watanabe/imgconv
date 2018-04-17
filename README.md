# 画像変換コマンド

## Gopher道場第１回の宿題。課題の内容は以下の通り

### 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する  
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換  
- [x] ディレクトリ以下は再帰的に処理する  
- [x] 変換前と変換後の画像形式を指定できる

### 以下を満たすように開発してください
- [x] mainパッケージと分離する  
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う  
- [x] 準標準パッケージ：golang.org/x以下のパッケージ  
- [ ] ユーザ定義型を作ってみる
- [ ] GoDocを生成してみる

---
## コマンド実行方法

### imgconv [ options ] [< args >]

コマンド実行例 : imgconv -s png -d gif -t ./

- -s 変換前画像形式(デフォルト:jpg)
- -d 変換後画像形式(デフォルト:png)
- -t 変換対象ディレクトリ(デフォルト:./)