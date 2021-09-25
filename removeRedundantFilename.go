package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	fileExt = []string{"md", "csv"} // 需要修改的文件类型
)

func main() {
	// 从当前目录开始
	removeRedundantFilename(".")
}

func removeRedundantFilename(baseDir string) {
	pathSep := string(os.PathSeparator)
	dir, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return
	}
	for _, file := range dir {
		// 目录特殊处理
		if file.IsDir() {
			names1Idx := strings.LastIndex(file.Name(), " ")
			if names1Idx != -1 {
				newFilename := file.Name()[:names1Idx]
				// 注意这块，注意补上文件的路径，避免出现找不到文件的错误：baseDir+pathSep+file.Name()
				err = os.Rename(baseDir+pathSep+file.Name(), baseDir+pathSep+newFilename)
				if err != nil {
					fmt.Println("folder rename failed(文件夹重名命失败)", baseDir+pathSep+file.Name())
				} else {
					fmt.Println("folder rename successful(文件重命名成功)", baseDir+pathSep+file.Name())
				}
				removeRedundantFilename(baseDir + pathSep + newFilename)
			} else {
				removeRedundantFilename(baseDir + pathSep + file.Name())
			}

		}
		// 重名命当前文件
		names1Idx := strings.LastIndex(file.Name(), " ")
		names2Idx := strings.LastIndex(file.Name(), ".")
		// 判断是否有多余的文件名
		if names1Idx != -1 && names2Idx != -1 {
			// 检查是否合法文件类型
			if !checkFileExt(file.Name()) {
				continue
			}
			newFilename := file.Name()[:names1Idx] + file.Name()[names2Idx:]
			// 注意这块，注意补上文件的路径，避免出现找不到文件的错误：baseDir+pathSep+file.Name()
			err = os.Rename(baseDir+pathSep+file.Name(), baseDir+pathSep+newFilename)
			if err != nil {
				fmt.Println("file rename failed(文件重名命失败)", baseDir+pathSep+file.Name())
			} else {
				fmt.Println("file rename successful(文件重名命成功)", baseDir+pathSep+newFilename)
			}
		}
	}
}

// 检查是否为目标修改的文件类型
func checkFileExt(filename string) bool {
	for _, v := range fileExt {
		if strings.Contains(filename, v) {
			return true
		}
	}
	return false
}
