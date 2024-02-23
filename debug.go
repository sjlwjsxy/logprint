package logprint //声明是属于哪个包的，包名最好和上级目录是一样的，名称最好是小写字母并且最好不要带下划线且见名知意

import (
	"fmt"
	"time"
)

func Debug(msg interface{}) { //接收任何类型的值
	t := time.Now() //标记当前时间
	fmt.Printf("[debug] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
	//做格式化，debug类型的日志，打印时间、内容、t表示再将时间传过来并格式化，如果想精确到微妙就加000，不需要就不加
}
