## Why

golang benchmark没有并发测试的支持。

## Usage

指定多少个goroutine，即多少个并发用户。指定worker方法，方法签名返回error，用于统计错误率。

```
import "github.com/XUJiahua/benchmark_parallel"
 
benchmark_parallel.Run(1000, func() error {
		// custom code
		err := something()
		return err
	})

```

结果打印：并发用户数，错误率，平均耗时。

```
1000 goroutines, 0 / 1000 error rate, 1832241696 nanoseconds/op

```

