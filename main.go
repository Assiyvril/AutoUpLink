package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func checkPath(path string) (bool, string, string) {
	// 检测 path 文件是否存在
	// 若不存在，报错
	// 路径字符串必须以“WeChat.exe”结尾
	defer func() {
		// 处理数组越界
		if err := recover(); err != nil {
			return
		}
	}()

	ending := path[len(path)-10:]

	if ending != "WeChat.exe" {

		return false, "必须以WeChat.exe结尾", ""
	}
	_, err := os.Stat(path)
	if err != nil {

		return false, "路径不存在", ""
	}
	return true, "路径正确", path
}

func readInputPath() string {
	// 读取并检查输入的路径
	defaultPath := `C:\Program Files (x86)\Tencent\WeChat\WeChat.exe`
	fmt.Println("请输入微信路径，若不输入，则使用默认路径：", defaultPath)
	inputReader := bufio.NewReader(os.Stdin)
	inputPath, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入出错了", err)
		return ""
	}
	// 去除 \r\n
	inputPath = strings.Replace(inputPath, "\n", "", -1)
	inputPath = strings.Replace(inputPath, "\r", "", -1)
	if inputPath == "" {
		_, _, defaultCheckPath := checkPath(defaultPath)
		if defaultCheckPath == "" {
			fmt.Println("默认路径不可用，必须输入路径")
			return ""
		} else {
			return defaultCheckPath
		}
	} else {
		checkResult, checkMsg, checkedPath := checkPath(inputPath)
		if !checkResult {
			fmt.Println(checkMsg)
			return ""
		} else {

			return checkedPath
		}
	}
}

func readInputCount() int {
	fmt.Println("请输入要打开的数量：")
	inputReader := bufio.NewReader(os.Stdin)
	inputCount, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入出错了", err)
		return 0
	}
	// 去除 \r\n
	inputCount = strings.Replace(inputCount, "\n", "", -1)
	inputCount = strings.Replace(inputCount, "\r", "", -1)
	if inputCount == "" {
		fmt.Println("必须输入数量")
		return 0
	}
	// 转换为 int
	count, err := strconv.Atoi(inputCount)
	if err != nil {
		fmt.Println("输入的数量不是数字")
		return 0
	}
	return count
}

func executeCommand(path string, execResult *chan bool) {
	// 执行命令
	cmd := exec.Command(path)
	err := cmd.Start()
	if err != nil {
		fmt.Println("创建进程出错了", err)
		*execResult <- false
	}
	*execResult <- true
}

//func createWorker(path string) chan bool {
//	var execResult chan bool
//	go executeCommand(path, &execResult)
//	return execResult
//}

func main() {
	fmt.Println("*****************************************************************************************************************")
	fmt.Println("*****************************************************************************************************************")
	fmt.Println("*****************************************************************************************************************")
	fmt.Println("*****************！！！！！ --- 使用多开程序之前，必须关闭已经运行的微信 --- ！！！！！************************")
	path := func() string {
		for {
			checkPath := readInputPath()
			if checkPath != "" {
				return checkPath
			}
		}
	}()

	count := func() int {
		for {
			checkCount := readInputCount()
			if checkCount != 0 {
				return checkCount
			}
		}
	}()

	channels := make([]chan bool, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan bool)
		go executeCommand(path, &channels[i])
	}
	// 循环检测，直到 channels 中的所有 channel 都返回 true
	successCount := 0
	failCount := 0
	//fmt.Println("channels: ", channels)
	for {
		if successCount+failCount == count {
			break
		}
		for _, channel := range channels {
			select {
			case result := <-channel:
				if result {
					successCount++
					fmt.Println("成功 +1, ", successCount)
					close(channel)

				} else {
					failCount++
					fmt.Println("失败 +1, ", failCount)
					close(channel)
				}
			default:
				//fmt.Println("无论如何,跳过")
				continue
			}
		}
		//暂停100毫秒
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("成功", successCount, "个;", "失败", failCount, "个")
}
