# Context Learning

This repository target learning how to use context and shows in code what kind of problem the context aims to solve.


## References

- [_Concurrency in Go_](http://shop.oreilly.com/product/0636920046189.do), 日本語版:[『Go言語による並行処理』](https://www.oreilly.co.jp/books/9784873118468/)
- [GoDoc Package context](https://golang.org/pkg/context/)

## Package Context

In the [Package context](https://golang.org/pkg/context/) is written.

> Package contest defines Context type, which carries deadlines, cancelation signals, and other request scoped values across API boundaries and between processes.

Go makes it easy to generate goroutine and do parallel processing.
In here, I shows acount the problem as toa how to manage the lifetime og goroutine.

GoRoutineが終了するときの種類

- Goroutineが処理を完了する
- 回復できないエラーにより処理が続けられない場合
- 停止が命令された場合

Contextは生存期間が不明なGoroutineを親と正しく結びつけて必要な時にキャンセルできるようにシグナルを伝えることを想定したPackage



## Test script results


```sh
$ go run .
2019/02/10 11:15:57 Client timeout is [70ms]
2019/02/10 11:15:57 Slow: Timeout [20.7998ms/20ms] Timeout by Server
2019/02/10 11:15:57 Fast: Timeout [20.7998ms/20ms] Timeout by Server
2019/02/10 11:15:57 Fast: Complete [30.5659ms/60ms] <nil>
2019/02/10 11:15:57 Fast: Complete [30.5659ms/80ms] <nil>
2019/02/10 11:15:57 Fast: Complete [30.5659ms/40ms] <nil>
2019/02/10 11:15:57 Slow: Timeout [40.5784ms/40ms] Timeout by Server
2019/02/10 11:15:57 Slow: Timeout [60.6092ms/60ms] Timeout by Server
2019/02/10 11:15:57 Slow: Cancel [70.584ms/80ms] context deadline exceeded
2019/02/10 11:15:57 -- WithProxies --
2019/02/10 11:15:57 FastProxy: through[to:45ms] Slow: Timeout [20.8947ms/20ms] Timeout by Server
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Slow: Timeout [20.8947ms/20ms] Timeout by Server
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Fast: Timeout [21.8919ms/20ms] Timeout by Server
2019/02/10 11:15:57 FastProxy: through[to:45ms] Fast: Timeout [21.8919ms/20ms] Timeout by Server
2019/02/10 11:15:57 FastProxy: through[to:45ms] Fast: Complete [30.8681ms/40ms] <nil>
2019/02/10 11:15:57 FastProxy: through[to:45ms] Fast: Complete [30.8681ms/80ms] <nil>
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Fast: Complete [30.8681ms/40ms] <nil>
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Fast: Complete [30.8681ms/60ms] <nil>
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Fast: Complete [30.8681ms/80ms] <nil>
2019/02/10 11:15:57 FastProxy: through[to:45ms] Fast: Complete [30.8681ms/60ms] <nil>
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Slow: Timeout [40.9924ms/40ms] Timeout by Server
2019/02/10 11:15:57 FastProxy: through[to:45ms] Slow: Timeout [40.9924ms/40ms] Timeout by Server
2019/02/10 11:15:57 FastProxy: through[to:45ms] Slow: Cancel [45.0934ms/80ms] context canceled
2019/02/10 11:15:57 FastProxy: through[to:45ms] Slow: Cancel [45.0934ms/60ms] context canceled
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Slow: Timeout [61.1621ms/60ms] Timeout by Server
2019/02/10 11:15:57 SlowProxy: through[to:75ms] Slow: Cancel [70.1348ms/80ms] context deadline exceeded
```