package pitaya

import (
	"context"
	"github.com/golang/protobuf/proto"
)

type AsyncCallback func(res interface{}, err error)
type AsyncRoutine func() (res interface{}, err error)

type CallbackTask struct {
	Callback AsyncCallback
	Res      interface{}
	Err      error
}

func AsyncTask(routine AsyncRoutine, callback AsyncCallback) {
	go func() {
		res, err := routine()
		handlerService.AppendTask(CallbackTask{Callback: callback, Res: res, Err: err})
	}()
}

func AsyncRPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message, callback AsyncCallback) {
	go func() {
		err := RPC(ctx, routeStr, reply, arg)
		handlerService.AppendTask(CallbackTask{Callback: callback, Res: reply, Err: err})
	}()
}
