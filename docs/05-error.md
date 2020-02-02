## 错误处理
Go中错误处理使用关键字panic和recover，通常还会使用defer来对出现出错之后对一些资源的关闭。
和Java来类比话，panic相当于throw，recover相当于catch，defer相当于finally，这个类比是
根据业务处理来类比的。

### defer
`defer` 用来关闭资源，函数退出之前执行，多个defer之间执行的顺序和定义的顺序相反，是一个栈结构

### panic
`panic` 用来显示抛出错误，参数是error，可以通过`errors.New`来new一个错误或者实现`Error()`方法自定义错误

### recover
`recover()` 可以catch住panic出来错误做相应的业务处理，让程序正常执行，`recover()`只能在`defer`中使用
