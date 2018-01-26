# Log
**描述**

简单的封装了beego的日志库，在原有的基础上增加了日志的配置和错误的堆栈信息，使其符合web的开发习惯。

 **配置说明**

```php
dir 保存的目录
maxlines 每个文件保存的最大行数，默认值 100000
maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
daily 是否按照每天 logrotate，默认是 true
maxdays 文件最多保存多少天，默认保存 30 天
rotate 是否开启 logrotate，默认是 true
async 是否缓存后输出，默认是 false
console 是否输出到控制， 默认是 false
```

 **使用**

```php
package main

import (
	"web-go/log"
)

func main() {
	log.Info("1.log", "hello123") //业务日志
	log.Error("1.log", "hello123456") //错误日志，加上堆栈信息
	log.GetLogger("1.log").Warning("warning ...") //可以调用BeeLogger对象的所有方法
}
```

**输出结果：**
```php
2018/01/26 16:39:12 [I] hello123
2018/01/26 16:39:12 [E] hello123456
web-go/log.Error
	D:/Go/GOPATH/src/web-go/log/log.go:126
main.main
	D:/Go/GOPATH/src/web-go/main.go:11
runtime.main
	D:/Go/src/runtime/proc.go:185
runtime.goexit
	D:/Go/src/runtime/asm_amd64.s:2197
2018/01/26 16:39:12 [W] warning ...
```

**其他**

   beego的日志模块请参考：https://beego.me/docs/module/logs.md
