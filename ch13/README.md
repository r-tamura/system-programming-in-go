# Go言語で知るプロセス（1）

プロセスに含まれるもの
 - プロセスID
 - プロセスグループID, セッショングループID
 - ユーザID, グループID
 - 実行ユーザID, 実行グループID
 - カレントディレクトリ
 - ファイルディスクリプタ

プロセスグループ(ジョブ)
パイプで実行されたプロセス群は同一プロセスグループ

セッションプロセスグループ
同一ターミナルから起動したプロセス・同一キーボードで出力先が同一のプロセスは同一セッショングループ

### User ID and Group ID
Windowsでは[GetTokenInformation][GetTokenInformation]というWindows APIがあるが、GoはそのAPIを実装していない


```
fmt.Printf("User ID: %d\n", os.Getuid())
fmt.Printf("Group ID: %d\n", os.Getgid())
groups, _ := os.Getgroups()
fmt.Printf("Subgroup IDs: %v\n", groups)
// User ID: -1
// Group ID: -1
// Subgroup IDs: []
```

[GetTokenInformation]: https://msdn.microsoft.com/en-us/library/windows/desktop/aa446671(v=vs.85).aspx

### Process I/O
プロセスには入力と出力がある
 - コマンドライン引数
 - 環境変数
 - 終了コード