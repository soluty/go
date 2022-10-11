package xerror

import (
	"context"
	"fmt"

	"github.com/soluty/go/xlog"
)

// HandleError 处理那些不知道该怎么处理的错误，一般是通过日志打印出来，也可以写到其它的地方
func HandleError(ctx context.Context, err error)  {
	fmt.Println(111)
}

func handle(){
	xlog.Info(context.Background(), "abc")
}