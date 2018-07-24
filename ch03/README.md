# 低レベルアクセスへの入り口（2）：io.Reader 前編

| 変数     | io.Reader | io.Writer | io.Seeker | io.Closer |
| -------- | :-------: | :-------: | :-------: | :-------: |
| os.Stdin |     o     |     x     |     x     |     o     |
| os.FIle  |     o     |     o     |     o     |     o     |
| net.Conn |     o     |     o     |     x     |     o     |