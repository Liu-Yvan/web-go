
package main

import (
	"fmt"
	"web-go/log"
)

func main() {
	log.Info("1.log", "hello123") //业务日志
	log.Error("1.log", "hello123456") //错误日志，加上堆栈信息
	log.GetLogger("1.log").Warning("warning ...") //可以调用BeeLogger对象的所有方法
	fmt.Println("end...")
}




