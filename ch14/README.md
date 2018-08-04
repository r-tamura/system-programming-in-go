# Go言語で知るプロセス（2）

### Goが提供するAPI
 - os.Process: 低レベル構造体
 - exec.Cmd: os.Processよりも高機能

親プロセスからStdinPipe(), StdoutPipe(), StderrPipe()で子プロセスの標準入力、標準出力、標準エラーにつながることができる

### 子プロセス機能のオプション
子プロセス起動のオプションとして
 - 子プロセスルートディレクトリ(chroot)
 - 子プロセスユーザID, グループID
 - 補助グループなどのCrdential構造体
 - デバッガAPI Ptraceフラグ
 - セッションIDやプロセスグループIDの初期化フラグ
 - 疑似ターミナル設定
 - フォアグランド動作フラグ
などがある。

子プロセス起動のオプションがPOSIZ系OSとWindowsで異なる。(WindowsではHideWindowがあるなど)
WindowsではCreateProcess APIを使い、Linuxではclone, unshareのシステムコールを使う。


### fork/exec
C言語での子プロセス生成
fork関数/exec関数を使うのが慣習。
execve関数は親プロセスのコマンドライン引数と環境変数を引き継ぐ。

### 子プロセスのコマンドライン引数
POSIX系OSでは複数の引数が指定された場合、子プロセスのプログラムへは分解され終端文字区切りで渡されるが、
Windowsでは1つの文字列として渡されるので各プログラムごとにエスケープやワイルドカードの処理を行う必要がある。