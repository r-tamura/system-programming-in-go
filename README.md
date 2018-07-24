# system-programming-in-go
「Goならわかるシステムプログラミング」写経レポジトリ


# Setup
  - Install Go Lang
  - Install Delve (a debugger for Go)
# Memos
  delve: Go言語のデバッグツール

### VSCodeから`dlv debug`実行時のdebugファイルがデバッグ終了時に削除されない
2018/7/24現在 `vscode-go`には自動削除する機能がないので、`launch.json`で`output`で出力ファイルパスを指定する。

```
{
  "output": "/tmp/delve/debug" // /tmp/delveディレクトリにdebugファイルが生成される
}
```

今後のアップデートで修正されるかもhttps://github.com/Microsoft/vscode-go/issues/1345#issuecomment-406125202