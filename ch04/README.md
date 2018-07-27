# What I learned in this chapter

### Builtin binary package
[]byte型と数値の変換を行う機能を持つ

#### func Read(r io.Reader, order ByteOrder, data interface{}) error
io.Readerから指定されたバイトオーダーでdataサイズ分のデータを読み出し、dataへ格納する。
dataは固定長の型へのポインタでなければならない。

