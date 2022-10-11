package xerror

import (
	"context"

	"github.com/soluty/go/xlog"
)

func handle(){
	xlog.Info(context.Background(), "abc")
}