package rpcclienticpt

import (
	"context"
	"encoding/json"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const MetadataKeyUid = "uid"

// gateway set uid to grpc metadata -> to rpc srv
func Token2uidInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var uid string
	ctxUid := ctx.Value(MetadataKeyUid)
	if ctxUid != nil {
		uid = ctxUid.(json.Number).String()
	}

	md := metadata.New(map[string]string{MetadataKeyUid: uid})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// rpc service get uid from ctx
func GatewayUidFromCtx(ctx context.Context) (uid int64) {
	uid, _ = ctx.Value(MetadataKeyUid).(json.Number).Int64()
	return
}

// rpc service get uid from ctx
func RpcUidFromCtx(ctx context.Context) (uid int64) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		tmp := md.Get(MetadataKeyUid)
		if len(tmp) > 0 {
			uid, _ = strconv.ParseInt(tmp[0], 10, 64)
			return
		}
	}
	return
}
