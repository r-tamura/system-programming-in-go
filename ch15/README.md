# Go言語で知るプロセス（3）

### シグナル Signal
 - プロセス間通信(Inter-Process Communication/IPC)
 - 割り込み(Interrupt)

 システムコール呼び出しでは6つ程度引数がとれるが、シグナルは一つ。システムコールを受けるカーネルは常時起動しているが、
 プロセスは停止している場合もあるので、シグナルはそのような場合も考慮されている。


### シグナルのライフサイクル(The lifecycle of signal)

```
CPUが0除算/メモリ範囲外アクセスなどで発生し -- raise --> カーネルがシグナルを生成 -- send --> 対象プロセス
```

対象プロセスでシグナルを処理(handle)する。デフォルトでは無視かプロセスの終了。