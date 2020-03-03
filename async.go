package pitaya

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/topfreegames/pitaya/service"
)



func AsyncTask(routine service.AsyncRoutine, callback service.AsyncCallback) {
	go func() {
		res, err := routine()
		handlerService.AppendTask(service.CallbackTask{Callback: callback, Res: res, Err: err})
	}()
}

func AsyncRPC(ctx context.Context, routeStr string, reply proto.Message, arg proto.Message, callback service.AsyncCallback) {
	go func() {
		err := RPC(ctx, routeStr, reply, arg)
		handlerService.AppendTask(service.CallbackTask{Callback: callback, Res: reply, Err: err})
	}()
}
