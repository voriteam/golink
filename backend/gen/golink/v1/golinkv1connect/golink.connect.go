// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: golink/v1/golink.proto

package golinkv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/nownabe/golink/backend/gen/golink/v1"
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
	// GolinkServiceName is the fully-qualified name of the GolinkService service.
	GolinkServiceName = "golink.v1.GolinkService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GolinkServiceCreateGolinkProcedure is the fully-qualified name of the GolinkService's
	// CreateGolink RPC.
	GolinkServiceCreateGolinkProcedure = "/golink.v1.GolinkService/CreateGolink"
	// GolinkServiceGetGolinkProcedure is the fully-qualified name of the GolinkService's GetGolink RPC.
	GolinkServiceGetGolinkProcedure = "/golink.v1.GolinkService/GetGolink"
	// GolinkServiceListGolinksProcedure is the fully-qualified name of the GolinkService's ListGolinks
	// RPC.
	GolinkServiceListGolinksProcedure = "/golink.v1.GolinkService/ListGolinks"
	// GolinkServiceListGolinksByUrlProcedure is the fully-qualified name of the GolinkService's
	// ListGolinksByUrl RPC.
	GolinkServiceListGolinksByUrlProcedure = "/golink.v1.GolinkService/ListGolinksByUrl"
	// GolinkServiceUpdateGolinkProcedure is the fully-qualified name of the GolinkService's
	// UpdateGolink RPC.
	GolinkServiceUpdateGolinkProcedure = "/golink.v1.GolinkService/UpdateGolink"
	// GolinkServiceDeleteGolinkProcedure is the fully-qualified name of the GolinkService's
	// DeleteGolink RPC.
	GolinkServiceDeleteGolinkProcedure = "/golink.v1.GolinkService/DeleteGolink"
	// GolinkServiceAddOwnerProcedure is the fully-qualified name of the GolinkService's AddOwner RPC.
	GolinkServiceAddOwnerProcedure = "/golink.v1.GolinkService/AddOwner"
	// GolinkServiceRemoveOwnerProcedure is the fully-qualified name of the GolinkService's RemoveOwner
	// RPC.
	GolinkServiceRemoveOwnerProcedure = "/golink.v1.GolinkService/RemoveOwner"
	// GolinkServiceGetMeProcedure is the fully-qualified name of the GolinkService's GetMe RPC.
	GolinkServiceGetMeProcedure = "/golink.v1.GolinkService/GetMe"
)

// GolinkServiceClient is a client for the golink.v1.GolinkService service.
type GolinkServiceClient interface {
	CreateGolink(context.Context, *connect_go.Request[v1.CreateGolinkRequest]) (*connect_go.Response[v1.CreateGolinkResponse], error)
	GetGolink(context.Context, *connect_go.Request[v1.GetGolinkRequest]) (*connect_go.Response[v1.GetGolinkResponse], error)
	ListGolinks(context.Context, *connect_go.Request[v1.ListGolinksRequest]) (*connect_go.Response[v1.ListGolinksResponse], error)
	ListGolinksByUrl(context.Context, *connect_go.Request[v1.ListGolinksByUrlRequest]) (*connect_go.Response[v1.ListGolinksByUrlResponse], error)
	UpdateGolink(context.Context, *connect_go.Request[v1.UpdateGolinkRequest]) (*connect_go.Response[v1.UpdateGolinkResponse], error)
	DeleteGolink(context.Context, *connect_go.Request[v1.DeleteGolinkRequest]) (*connect_go.Response[v1.DeleteGolinkResponse], error)
	AddOwner(context.Context, *connect_go.Request[v1.AddOwnerRequest]) (*connect_go.Response[v1.AddOwnerResponse], error)
	RemoveOwner(context.Context, *connect_go.Request[v1.RemoveOwnerRequest]) (*connect_go.Response[v1.RemoveOwnerResponse], error)
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewGolinkServiceClient constructs a client for the golink.v1.GolinkService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGolinkServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GolinkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &golinkServiceClient{
		createGolink: connect_go.NewClient[v1.CreateGolinkRequest, v1.CreateGolinkResponse](
			httpClient,
			baseURL+GolinkServiceCreateGolinkProcedure,
			opts...,
		),
		getGolink: connect_go.NewClient[v1.GetGolinkRequest, v1.GetGolinkResponse](
			httpClient,
			baseURL+GolinkServiceGetGolinkProcedure,
			opts...,
		),
		listGolinks: connect_go.NewClient[v1.ListGolinksRequest, v1.ListGolinksResponse](
			httpClient,
			baseURL+GolinkServiceListGolinksProcedure,
			opts...,
		),
		listGolinksByUrl: connect_go.NewClient[v1.ListGolinksByUrlRequest, v1.ListGolinksByUrlResponse](
			httpClient,
			baseURL+GolinkServiceListGolinksByUrlProcedure,
			opts...,
		),
		updateGolink: connect_go.NewClient[v1.UpdateGolinkRequest, v1.UpdateGolinkResponse](
			httpClient,
			baseURL+GolinkServiceUpdateGolinkProcedure,
			opts...,
		),
		deleteGolink: connect_go.NewClient[v1.DeleteGolinkRequest, v1.DeleteGolinkResponse](
			httpClient,
			baseURL+GolinkServiceDeleteGolinkProcedure,
			opts...,
		),
		addOwner: connect_go.NewClient[v1.AddOwnerRequest, v1.AddOwnerResponse](
			httpClient,
			baseURL+GolinkServiceAddOwnerProcedure,
			opts...,
		),
		removeOwner: connect_go.NewClient[v1.RemoveOwnerRequest, v1.RemoveOwnerResponse](
			httpClient,
			baseURL+GolinkServiceRemoveOwnerProcedure,
			opts...,
		),
		getMe: connect_go.NewClient[v1.GetMeRequest, v1.GetMeResponse](
			httpClient,
			baseURL+GolinkServiceGetMeProcedure,
			opts...,
		),
	}
}

// golinkServiceClient implements GolinkServiceClient.
type golinkServiceClient struct {
	createGolink     *connect_go.Client[v1.CreateGolinkRequest, v1.CreateGolinkResponse]
	getGolink        *connect_go.Client[v1.GetGolinkRequest, v1.GetGolinkResponse]
	listGolinks      *connect_go.Client[v1.ListGolinksRequest, v1.ListGolinksResponse]
	listGolinksByUrl *connect_go.Client[v1.ListGolinksByUrlRequest, v1.ListGolinksByUrlResponse]
	updateGolink     *connect_go.Client[v1.UpdateGolinkRequest, v1.UpdateGolinkResponse]
	deleteGolink     *connect_go.Client[v1.DeleteGolinkRequest, v1.DeleteGolinkResponse]
	addOwner         *connect_go.Client[v1.AddOwnerRequest, v1.AddOwnerResponse]
	removeOwner      *connect_go.Client[v1.RemoveOwnerRequest, v1.RemoveOwnerResponse]
	getMe            *connect_go.Client[v1.GetMeRequest, v1.GetMeResponse]
}

// CreateGolink calls golink.v1.GolinkService.CreateGolink.
func (c *golinkServiceClient) CreateGolink(ctx context.Context, req *connect_go.Request[v1.CreateGolinkRequest]) (*connect_go.Response[v1.CreateGolinkResponse], error) {
	return c.createGolink.CallUnary(ctx, req)
}

// GetGolink calls golink.v1.GolinkService.GetGolink.
func (c *golinkServiceClient) GetGolink(ctx context.Context, req *connect_go.Request[v1.GetGolinkRequest]) (*connect_go.Response[v1.GetGolinkResponse], error) {
	return c.getGolink.CallUnary(ctx, req)
}

// ListGolinks calls golink.v1.GolinkService.ListGolinks.
func (c *golinkServiceClient) ListGolinks(ctx context.Context, req *connect_go.Request[v1.ListGolinksRequest]) (*connect_go.Response[v1.ListGolinksResponse], error) {
	return c.listGolinks.CallUnary(ctx, req)
}

// ListGolinksByUrl calls golink.v1.GolinkService.ListGolinksByUrl.
func (c *golinkServiceClient) ListGolinksByUrl(ctx context.Context, req *connect_go.Request[v1.ListGolinksByUrlRequest]) (*connect_go.Response[v1.ListGolinksByUrlResponse], error) {
	return c.listGolinksByUrl.CallUnary(ctx, req)
}

// UpdateGolink calls golink.v1.GolinkService.UpdateGolink.
func (c *golinkServiceClient) UpdateGolink(ctx context.Context, req *connect_go.Request[v1.UpdateGolinkRequest]) (*connect_go.Response[v1.UpdateGolinkResponse], error) {
	return c.updateGolink.CallUnary(ctx, req)
}

// DeleteGolink calls golink.v1.GolinkService.DeleteGolink.
func (c *golinkServiceClient) DeleteGolink(ctx context.Context, req *connect_go.Request[v1.DeleteGolinkRequest]) (*connect_go.Response[v1.DeleteGolinkResponse], error) {
	return c.deleteGolink.CallUnary(ctx, req)
}

// AddOwner calls golink.v1.GolinkService.AddOwner.
func (c *golinkServiceClient) AddOwner(ctx context.Context, req *connect_go.Request[v1.AddOwnerRequest]) (*connect_go.Response[v1.AddOwnerResponse], error) {
	return c.addOwner.CallUnary(ctx, req)
}

// RemoveOwner calls golink.v1.GolinkService.RemoveOwner.
func (c *golinkServiceClient) RemoveOwner(ctx context.Context, req *connect_go.Request[v1.RemoveOwnerRequest]) (*connect_go.Response[v1.RemoveOwnerResponse], error) {
	return c.removeOwner.CallUnary(ctx, req)
}

// GetMe calls golink.v1.GolinkService.GetMe.
func (c *golinkServiceClient) GetMe(ctx context.Context, req *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return c.getMe.CallUnary(ctx, req)
}

// GolinkServiceHandler is an implementation of the golink.v1.GolinkService service.
type GolinkServiceHandler interface {
	CreateGolink(context.Context, *connect_go.Request[v1.CreateGolinkRequest]) (*connect_go.Response[v1.CreateGolinkResponse], error)
	GetGolink(context.Context, *connect_go.Request[v1.GetGolinkRequest]) (*connect_go.Response[v1.GetGolinkResponse], error)
	ListGolinks(context.Context, *connect_go.Request[v1.ListGolinksRequest]) (*connect_go.Response[v1.ListGolinksResponse], error)
	ListGolinksByUrl(context.Context, *connect_go.Request[v1.ListGolinksByUrlRequest]) (*connect_go.Response[v1.ListGolinksByUrlResponse], error)
	UpdateGolink(context.Context, *connect_go.Request[v1.UpdateGolinkRequest]) (*connect_go.Response[v1.UpdateGolinkResponse], error)
	DeleteGolink(context.Context, *connect_go.Request[v1.DeleteGolinkRequest]) (*connect_go.Response[v1.DeleteGolinkResponse], error)
	AddOwner(context.Context, *connect_go.Request[v1.AddOwnerRequest]) (*connect_go.Response[v1.AddOwnerResponse], error)
	RemoveOwner(context.Context, *connect_go.Request[v1.RemoveOwnerRequest]) (*connect_go.Response[v1.RemoveOwnerResponse], error)
	GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error)
}

// NewGolinkServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGolinkServiceHandler(svc GolinkServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	golinkServiceCreateGolinkHandler := connect_go.NewUnaryHandler(
		GolinkServiceCreateGolinkProcedure,
		svc.CreateGolink,
		opts...,
	)
	golinkServiceGetGolinkHandler := connect_go.NewUnaryHandler(
		GolinkServiceGetGolinkProcedure,
		svc.GetGolink,
		opts...,
	)
	golinkServiceListGolinksHandler := connect_go.NewUnaryHandler(
		GolinkServiceListGolinksProcedure,
		svc.ListGolinks,
		opts...,
	)
	golinkServiceListGolinksByUrlHandler := connect_go.NewUnaryHandler(
		GolinkServiceListGolinksByUrlProcedure,
		svc.ListGolinksByUrl,
		opts...,
	)
	golinkServiceUpdateGolinkHandler := connect_go.NewUnaryHandler(
		GolinkServiceUpdateGolinkProcedure,
		svc.UpdateGolink,
		opts...,
	)
	golinkServiceDeleteGolinkHandler := connect_go.NewUnaryHandler(
		GolinkServiceDeleteGolinkProcedure,
		svc.DeleteGolink,
		opts...,
	)
	golinkServiceAddOwnerHandler := connect_go.NewUnaryHandler(
		GolinkServiceAddOwnerProcedure,
		svc.AddOwner,
		opts...,
	)
	golinkServiceRemoveOwnerHandler := connect_go.NewUnaryHandler(
		GolinkServiceRemoveOwnerProcedure,
		svc.RemoveOwner,
		opts...,
	)
	golinkServiceGetMeHandler := connect_go.NewUnaryHandler(
		GolinkServiceGetMeProcedure,
		svc.GetMe,
		opts...,
	)
	return "/golink.v1.GolinkService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GolinkServiceCreateGolinkProcedure:
			golinkServiceCreateGolinkHandler.ServeHTTP(w, r)
		case GolinkServiceGetGolinkProcedure:
			golinkServiceGetGolinkHandler.ServeHTTP(w, r)
		case GolinkServiceListGolinksProcedure:
			golinkServiceListGolinksHandler.ServeHTTP(w, r)
		case GolinkServiceListGolinksByUrlProcedure:
			golinkServiceListGolinksByUrlHandler.ServeHTTP(w, r)
		case GolinkServiceUpdateGolinkProcedure:
			golinkServiceUpdateGolinkHandler.ServeHTTP(w, r)
		case GolinkServiceDeleteGolinkProcedure:
			golinkServiceDeleteGolinkHandler.ServeHTTP(w, r)
		case GolinkServiceAddOwnerProcedure:
			golinkServiceAddOwnerHandler.ServeHTTP(w, r)
		case GolinkServiceRemoveOwnerProcedure:
			golinkServiceRemoveOwnerHandler.ServeHTTP(w, r)
		case GolinkServiceGetMeProcedure:
			golinkServiceGetMeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGolinkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGolinkServiceHandler struct{}

func (UnimplementedGolinkServiceHandler) CreateGolink(context.Context, *connect_go.Request[v1.CreateGolinkRequest]) (*connect_go.Response[v1.CreateGolinkResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.CreateGolink is not implemented"))
}

func (UnimplementedGolinkServiceHandler) GetGolink(context.Context, *connect_go.Request[v1.GetGolinkRequest]) (*connect_go.Response[v1.GetGolinkResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.GetGolink is not implemented"))
}

func (UnimplementedGolinkServiceHandler) ListGolinks(context.Context, *connect_go.Request[v1.ListGolinksRequest]) (*connect_go.Response[v1.ListGolinksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.ListGolinks is not implemented"))
}

func (UnimplementedGolinkServiceHandler) ListGolinksByUrl(context.Context, *connect_go.Request[v1.ListGolinksByUrlRequest]) (*connect_go.Response[v1.ListGolinksByUrlResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.ListGolinksByUrl is not implemented"))
}

func (UnimplementedGolinkServiceHandler) UpdateGolink(context.Context, *connect_go.Request[v1.UpdateGolinkRequest]) (*connect_go.Response[v1.UpdateGolinkResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.UpdateGolink is not implemented"))
}

func (UnimplementedGolinkServiceHandler) DeleteGolink(context.Context, *connect_go.Request[v1.DeleteGolinkRequest]) (*connect_go.Response[v1.DeleteGolinkResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.DeleteGolink is not implemented"))
}

func (UnimplementedGolinkServiceHandler) AddOwner(context.Context, *connect_go.Request[v1.AddOwnerRequest]) (*connect_go.Response[v1.AddOwnerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.AddOwner is not implemented"))
}

func (UnimplementedGolinkServiceHandler) RemoveOwner(context.Context, *connect_go.Request[v1.RemoveOwnerRequest]) (*connect_go.Response[v1.RemoveOwnerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.RemoveOwner is not implemented"))
}

func (UnimplementedGolinkServiceHandler) GetMe(context.Context, *connect_go.Request[v1.GetMeRequest]) (*connect_go.Response[v1.GetMeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("golink.v1.GolinkService.GetMe is not implemented"))
}
