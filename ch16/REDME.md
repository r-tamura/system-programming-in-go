# Go言語と並列処理

# goroutine
goroutineの起動はネイティブスレッドより高速だがそれでも遅い

# チャネル

チャネルの状態と振る舞い

 状態 | バッファなし | バッファあり | 閉じたチャネル
 ----|:---------:|:---------:|:----------
 作り方 | make(chan 型) | make(chan 型, 要素数) | close(宣言済みのチャネル)
 送信　| 受信側が受信するまでブロック | バッファがあればすぐ終了。ない場合はバッファなしと同じ | パニック
 受信 | 送信側がデータをいれるまでブロック | 送信側がデータをいれるまでブロック | デフォルト値が返る, 第二返り値がfalse


# Select文
"ブロックしうる複数のチャネルを並列で読み込み読み込めたものを処理する"。考え方はselect属システムコールと同じ。


# 並列・並行処理の手法のパターン
 1. マルチプロセス
 1. イベント駆動
 1. マルチスレッド

### マルチプロセス (Multiprocessing)
 - メモリ空間は分離されるのでそのあたりは安全
 - CoWとはいえFD等のコピーは発生したりするので起動は遅い
 - コンテキストスイッチにより実行コストは高い

### イベント駆動 (Event driven)
 - 並行処理
 - I/O多重化

### マルチスレッド (Multithreading)
 - 起動・コンテキストスイッチはプロセスよりも早い
