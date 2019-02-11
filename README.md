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
$ go run ./
2019/02/11 09:49:41 Client timeout is [70ms]
2019/02/11 09:49:41 Fast: Timeout [20.9447ms/20ms] Timeout by Server
2019/02/11 09:49:41 Slow: Timeout [20.9447ms/20ms] Timeout by Server
2019/02/11 09:49:41 Fast: Complete [30.0891ms/60ms] <nil>
2019/02/11 09:49:41 Fast: Complete [30.0891ms/40ms] <nil>
2019/02/11 09:49:41 Fast: Complete [30.0891ms/80ms] <nil>
2019/02/11 09:49:41 Slow: Timeout [40.2519ms/40ms] Timeout by Server
2019/02/11 09:49:41 Slow: Timeout [60.2142ms/60ms] Timeout by Server
2019/02/11 09:49:41 Slow: Cancel [71.8894ms/80ms] context deadline exceeded
2019/02/11 09:49:41 -- WithProxies --
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Fast: Timeout [20.1622ms/20ms] Timeout by Server
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Slow: Timeout [20.1622ms/20ms] Timeout by Server
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Slow: Timeout [20.1622ms/20ms] Timeout by Server
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Fast: Timeout [20.1622ms/20ms] Timeout by Server
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Fast: Complete [30.9573ms/60ms] <nil>
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Fast: Complete [30.9573ms/40ms] <nil>
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Fast: Complete [30.9573ms/80ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Fast: Complete [30.9573ms/60ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Fast: Complete [30.9573ms/80ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Fast: Complete [30.9573ms/40ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Slow: Timeout [40.8041ms/40ms] Timeout by Server
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Slow: Timeout [40.8834ms/40ms] Timeout by Server
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Slow: Cancel [45.0796ms/80ms] context canceled
2019/02/11 09:49:41 FastProxy: through[TO:45ms] Slow: Cancel [45.0796ms/60ms] context canceled
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Slow: Timeout [61.0672ms/60ms] Timeout by Server
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] Slow: Cancel [70.9447ms/80ms] context deadline exceeded
2019/02/11 09:49:41 -- Heazy Process Simulation --
2019/02/11 09:49:41 FastProxy: through[TO:45ms] l:1ms: Complete [loop: 20] [33.804ms/40ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] l:1ms: Complete [loop: 20] [33.804ms/40ms] <nil>
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] l:1ms: Timeout [41.9723ms/40ms] Timeout by Simulation
2019/02/11 09:49:41 FastProxy: through[TO:45ms] l:1ms: Timeout [41.9723ms/40ms] Timeout by Simulation
2019/02/11 09:49:41 FastProxy: through[TO:45ms] l:1ms: Cancel [45.7827ms/80ms] context canceled
2019/02/11 09:49:41 FastProxy: through[TO:45ms] l:1ms: Cancel [45.7827ms/80ms] context canceled
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] l:1ms: Cancel [70.7212ms/80ms] context deadline exceeded
2019/02/11 09:49:41 SlowProxy: through[TO:75ms] l:1ms: Cancel [70.7212ms/80ms] context deadline exceeded
```