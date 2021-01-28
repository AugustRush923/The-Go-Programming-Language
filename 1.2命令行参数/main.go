package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println(len(os.Args))
	for _,item := range os.Args[1:] {
		fmt.Println(item)
	}
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// 练习1.1：修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字
	fmt.Println(os.Args[0])
	// 练习1.2：修改 echo 程序，使其打印每个参数的索引和值，每个一行。
	for index, params := range os.Args{
		fmt.Println(index, params)
	}

	// 练习1.3：做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6节讲
	// 解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
}
