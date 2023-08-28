package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibonaqi() func() int {
	// 闭包实现斐波那契数列
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(fileName string) {
	// 写入文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close() // 先写的 defer， 最后执行这个defer，关闭文件

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 倒数第二个 defer，先执行这个 defer，将缓存的数据写入文件
	f := fibonaqi()
	for i := 0; i < 50; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	writeFile("fib.txt")
}
