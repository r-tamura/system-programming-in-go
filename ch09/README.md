# Unixドメインソケット (Windows名前付きパイプ)
 - 同一ホスト内の通信のみ可能
 - TCPと比較して(HTTPで)数十倍高速
 - ファイルI/OやTCPソケットでは非同期API syscall.Syscallを使用するが、UNIXソケットドメインは高速なので同期式syscall.RawSyscallを利用する