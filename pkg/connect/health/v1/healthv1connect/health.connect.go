// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: health/v1/health.proto

package healthv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/takokun778/template-module/pkg/connect/health/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// HealthServiceName is the fully-qualified name of the HealthService service.
	HealthServiceName = "health.v1.HealthService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// HealthServiceCheckProcedure is the fully-qualified name of the HealthService's Check RPC.
	HealthServiceCheckProcedure = "/health.v1.HealthService/Check"
)

// HealthServiceClient is a client for the health.v1.HealthService service.
type HealthServiceClient interface {
	// チェック
	Check(context.Context, *connect_go.Request[v1.CheckRequest]) (*connect_go.Response[v1.CheckResponse], error)
}

// NewHealthServiceClient constructs a client for the health.v1.HealthService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewHealthServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) HealthServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &healthServiceClient{
		check: connect_go.NewClient[v1.CheckRequest, v1.CheckResponse](
			httpClient,
			baseURL+HealthServiceCheckProcedure,
			opts...,
		),
	}
}

// healthServiceClient implements HealthServiceClient.
type healthServiceClient struct {
	check *connect_go.Client[v1.CheckRequest, v1.CheckResponse]
}

// Check calls health.v1.HealthService.Check.
func (c *healthServiceClient) Check(ctx context.Context, req *connect_go.Request[v1.CheckRequest]) (*connect_go.Response[v1.CheckResponse], error) {
	return c.check.CallUnary(ctx, req)
}

// HealthServiceHandler is an implementation of the health.v1.HealthService service.
type HealthServiceHandler interface {
	// チェック
	Check(context.Context, *connect_go.Request[v1.CheckRequest]) (*connect_go.Response[v1.CheckResponse], error)
}

// NewHealthServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewHealthServiceHandler(svc HealthServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	healthServiceCheckHandler := connect_go.NewUnaryHandler(
		HealthServiceCheckProcedure,
		svc.Check,
		opts...,
	)
	return "/health.v1.HealthService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case HealthServiceCheckProcedure:
			healthServiceCheckHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedHealthServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedHealthServiceHandler struct{}

func (UnimplementedHealthServiceHandler) Check(context.Context, *connect_go.Request[v1.CheckRequest]) (*connect_go.Response[v1.CheckResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("health.v1.HealthService.Check is not implemented"))
}