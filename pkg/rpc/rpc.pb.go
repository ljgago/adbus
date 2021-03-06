// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package apigrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	api "proto-files/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x71, 0xa4, 0x7b, 0x2f, 0x99, 0x10, 0xc8, 0x1d, 0x01, 0x2a, 0x0e, 0x81, 0xe0, 0xf2,
	0x4f, 0xa8, 0xb1, 0x81, 0x6e, 0x5a, 0x36, 0x55, 0x52, 0x2a, 0xaa, 0xb6, 0x0b, 0x04, 0xed, 0x86,
	0x0d, 0x75, 0x9c, 0xc1, 0x0c, 0x75, 0xed, 0xa9, 0x67, 0x02, 0x8a, 0x50, 0x85, 0xd4, 0x55, 0x57,
	0x95, 0x4a, 0x77, 0x7d, 0x8f, 0xbe, 0x40, 0x5f, 0xa1, 0xaf, 0xd0, 0x37, 0xe8, 0x0b, 0x54, 0x1e,
	0xcf, 0x24, 0xe3, 0x3f, 0x21, 0x15, 0xea, 0xca, 0xd2, 0x7c, 0x87, 0xf3, 0xfb, 0xbe, 0xc3, 0xf1,
	0xc4, 0xa0, 0x18, 0x12, 0xc7, 0x24, 0x61, 0xc0, 0x02, 0x58, 0xb4, 0x3b, 0xed, 0x2e, 0x35, 0x6d,
	0x82, 0xf5, 0xaa, 0x1b, 0x04, 0xae, 0x87, 0x2c, 0x2e, 0xb4, 0xbb, 0x27, 0x16, 0x7a, 0x4b, 0x58,
	0x2f, 0xae, 0xd3, 0xe7, 0x85, 0x68, 0x13, 0x6c, 0xd9, 0xbe, 0x1f, 0x30, 0x9b, 0xe1, 0xc0, 0xa7,
	0x42, 0xbd, 0xc7, 0x1f, 0x4e, 0xc3, 0x45, 0x7e, 0x83, 0x5e, 0xd8, 0xae, 0x8b, 0x42, 0x2b, 0x20,
	0xbc, 0x22, 0xa7, 0x7a, 0x8e, 0x3f, 0x1a, 0x27, 0xd8, 0x43, 0x34, 0x6e, 0x48, 0x70, 0x2c, 0x6d,
	0xff, 0xaa, 0x80, 0xc9, 0x5d, 0x74, 0x8e, 0x1d, 0x44, 0x0f, 0x51, 0x18, 0x3d, 0xe1, 0x15, 0x28,
	0xbd, 0xc0, 0x94, 0x89, 0x53, 0x38, 0x6b, 0xc6, 0x4e, 0x4c, 0x69, 0xd3, 0x7c, 0x12, 0xd9, 0xd4,
	0x17, 0xcc, 0x7e, 0x12, 0x53, 0xa9, 0x3f, 0x40, 0x94, 0x04, 0x3e, 0x45, 0xc6, 0x83, 0xeb, 0xa6,
	0xd1, 0xae, 0x83, 0x29, 0x50, 0x7a, 0xca, 0x18, 0x79, 0x8e, 0x7a, 0xcd, 0x2e, 0x3b, 0x85, 0x63,
	0xa0, 0x0c, 0x8a, 0x2d, 0x64, 0x87, 0x28, 0x7c, 0x76, 0xc1, 0xe0, 0xd8, 0x87, 0x1f, 0x3f, 0xbf,
	0x14, 0xca, 0xb0, 0x64, 0x9d, 0x6f, 0x59, 0x1d, 0x41, 0x74, 0xc1, 0xc4, 0xe3, 0x10, 0xd9, 0x0c,
	0xc5, 0x2d, 0xa1, 0x4a, 0x52, 0x85, 0x03, 0xf4, 0xae, 0x8b, 0x28, 0xd3, 0x17, 0x87, 0xea, 0xc2,
	0xca, 0x2c, 0xa7, 0x54, 0x0c, 0x95, 0xb2, 0xa3, 0x6d, 0xc0, 0xd7, 0xa0, 0xb8, 0x87, 0x84, 0x71,
	0x58, 0x55, 0xba, 0xf4, 0x4f, 0x25, 0x62, 0x3e, 0x5f, 0x14, 0xfd, 0xef, 0xf0, 0xfe, 0x10, 0x56,
	0x94, 0xfe, 0xd6, 0x25, 0xee, 0xbc, 0x87, 0x1e, 0x98, 0x78, 0x45, 0x3a, 0xf9, 0x51, 0x54, 0x21,
	0x2f, 0x4a, 0x52, 0x17, 0xa8, 0x2a, 0x47, 0xcd, 0xe8, 0x19, 0x54, 0x94, 0x07, 0x83, 0x89, 0x5d,
	0xe4, 0xa1, 0x5c, 0x9a, 0x2a, 0xe4, 0xd1, 0x92, 0x7a, 0x32, 0xd8, 0x46, 0x36, 0x58, 0x00, 0x4a,
	0x7b, 0x88, 0xed, 0x7b, 0x76, 0xcf, 0xc3, 0x94, 0xc1, 0x5a, 0x72, 0x3e, 0xf2, 0x5c, 0x82, 0x16,
	0x86, 0xc9, 0x82, 0xb3, 0xc4, 0x39, 0x55, 0x38, 0x97, 0xe6, 0x58, 0x44, 0x12, 0x2e, 0xc1, 0x64,
	0x3c, 0x90, 0x3e, 0xb3, 0x9e, 0x99, 0x55, 0x1a, 0xbb, 0x74, 0x43, 0x85, 0x20, 0x2f, 0x73, 0xf2,
	0x82, 0x3e, 0x9c, 0x1c, 0x0d, 0x16, 0x01, 0x10, 0xad, 0xf8, 0x5e, 0x18, 0x74, 0x09, 0x85, 0xf3,
	0xa9, 0xcd, 0x8f, 0x8f, 0x25, 0xb4, 0x36, 0x44, 0x15, 0x40, 0x9d, 0x03, 0xa7, 0x21, 0x54, 0x81,
	0x6e, 0xdc, 0xf8, 0xa3, 0x06, 0xa0, 0xf2, 0x2a, 0xb5, 0x7a, 0xfc, 0x6f, 0xe1, 0x72, 0xfe, 0x9b,
	0x26, 0x64, 0xc9, 0x5d, 0x19, 0x51, 0x25, 0xf8, 0x2b, 0x9c, 0xbf, 0x08, 0x6b, 0x59, 0xbe, 0x75,
	0xc9, 0x9f, 0xc7, 0xd1, 0xff, 0xf7, 0x93, 0x06, 0xa6, 0xd5, 0x05, 0xec, 0x9b, 0x59, 0x1d, 0xb2,
	0xa1, 0x69, 0x3b, 0x6b, 0x23, 0xeb, 0x92, 0x86, 0xf4, 0x11, 0x86, 0x3e, 0x6b, 0x00, 0x2a, 0xab,
	0x93, 0x37, 0x9b, 0xac, 0x9c, 0x37, 0x9b, 0xbc, 0x2a, 0x61, 0xc5, 0xe4, 0x56, 0xd6, 0xe1, 0xea,
	0x8d, 0x56, 0x06, 0x3b, 0xf9, 0x55, 0x03, 0x33, 0xc9, 0xbd, 0x92, 0xb6, 0xd6, 0x86, 0x6e, 0x5e,
	0xca, 0xd9, 0xfa, 0xe8, 0xc2, 0xa4, 0x39, 0xfd, 0x4f, 0xcd, 0x5d, 0x81, 0xff, 0x9b, 0x4e, 0xf4,
	0x2b, 0x70, 0xd8, 0xf3, 0x1d, 0xe9, 0xeb, 0xae, 0x82, 0xcb, 0xa8, 0xd2, 0xd3, 0x90, 0x1b, 0xdf,
	0x68, 0x70, 0x07, 0x6b, 0x86, 0x71, 0xb3, 0x03, 0xda, 0xf3, 0x9d, 0xe8, 0xa5, 0xe9, 0x1b, 0x78,
	0x89, 0x06, 0x83, 0xc9, 0x1a, 0x50, 0xd4, 0xbf, 0x64, 0x80, 0x21, 0xf9, 0xd6, 0x96, 0x07, 0x19,
	0x9b, 0x9e, 0x07, 0x17, 0x73, 0xd3, 0x37, 0x3d, 0x6f, 0x14, 0x58, 0xdc, 0xba, 0x46, 0xe2, 0x1e,
	0x94, 0x39, 0x3b, 0x00, 0x0c, 0x9a, 0x25, 0x2e, 0x87, 0xc1, 0xf1, 0x28, 0x40, 0x9d, 0x03, 0x74,
	0x63, 0x26, 0x73, 0x0d, 0x49, 0x4a, 0x3f, 0x4c, 0x34, 0xaf, 0xfc, 0x30, 0x42, 0xb9, 0x55, 0x18,
	0x39, 0xb3, 0x7e, 0x98, 0xa8, 0x59, 0x4e, 0x98, 0xe8, 0xf8, 0xf6, 0x61, 0x04, 0xa5, 0xf5, 0x5d,
	0xbb, 0x6e, 0x7e, 0xd3, 0xe0, 0x43, 0x50, 0x6e, 0x46, 0xfd, 0xeb, 0xfb, 0x61, 0x70, 0x86, 0x1c,
	0x66, 0x2c, 0xa5, 0x0e, 0x60, 0xe5, 0x94, 0x31, 0x42, 0x77, 0x2c, 0x2b, 0xf6, 0x81, 0x83, 0xed,
	0x7f, 0x36, 0xcd, 0x4d, 0x73, 0x6b, 0xa3, 0xa0, 0x15, 0xb6, 0x2b, 0x36, 0x21, 0x1e, 0x76, 0xf8,
	0xa7, 0x8e, 0x75, 0x46, 0x03, 0x7f, 0x27, 0x73, 0x72, 0xf4, 0x08, 0x4c, 0x81, 0x62, 0xcb, 0xa6,
	0xd8, 0xe1, 0x9f, 0x1f, 0x85, 0x71, 0x2d, 0xf5, 0x01, 0x02, 0x6a, 0xc9, 0x0f, 0x94, 0xc9, 0xf1,
	0x82, 0x3e, 0x1e, 0x51, 0x8f, 0xdf, 0xa0, 0x5e, 0xbd, 0xd0, 0x9e, 0x4a, 0xd5, 0x1f, 0xfd, 0x67,
	0x13, 0xec, 0x86, 0xc4, 0x69, 0xff, 0xcb, 0x63, 0xdf, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb2,
	0xde, 0x68, 0xfd, 0xdd, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DevicesServiceClient is the client API for DevicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicesServiceClient interface {
	// GET /v1/devices
	ListDevices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.ListDevicesResponse, error)
	// POST /v1/devices
	CreateDevice(ctx context.Context, in *api.CreateDeviceRequest, opts ...grpc.CallOption) (*api.CreateDeviceResponse, error)
	// GET /v1/devices/{id}
	GetDevice(ctx context.Context, in *api.GetDeviceRequest, opts ...grpc.CallOption) (*api.GetDeviceResponse, error)
	// PUT /v1/devices/{id}
	UpdateDevice(ctx context.Context, in *api.UpdateDeviceRequest, opts ...grpc.CallOption) (*api.UpdateDeviceResponse, error)
	// DELETE /v1/devices/{id}
	DeleteDevice(ctx context.Context, in *api.DeleteDeviceRequest, opts ...grpc.CallOption) (*api.DeleteDeviceResponse, error)
	// GET /v1/devices/{id}/playlist
	GetPlaylist(ctx context.Context, in *api.GetPlaylistRequest, opts ...grpc.CallOption) (*api.GetPlaylistResponse, error)
	// PUT /v1/devices/{id}/playlist
	UpdatePlaylist(ctx context.Context, in *api.UpdatePlaylistRequest, opts ...grpc.CallOption) (*api.UpdatePlaylistResponse, error)
	// GET /v1/devices/groups
	ListGroups(ctx context.Context, in *api.ListGroupsRequest, opts ...grpc.CallOption) (*api.ListGroupsResponse, error)
	// GET /v1/devices/groups/{group_id}
	ListDevicesByGroup(ctx context.Context, in *api.ListDevicesByGroupRequest, opts ...grpc.CallOption) (*api.ListDevicesByGroupResponse, error)
	// PUT /v1/devices/groups/{group_id}
	UpdateDevicesByGroup(ctx context.Context, in *api.UpdateDevicesByGroupRequest, opts ...grpc.CallOption) (*api.UpdateDevicesByGroupResponse, error)
	// GET /v1/devices/groups/{group_id}/playlist
	GetPlaylistByGroup(ctx context.Context, in *api.GetPlaylistByGroupRequest, opts ...grpc.CallOption) (*api.GetPlaylistByGroupResponse, error)
	// PUT /v1/devices/groups/{group_id}/playlist
	UpdatePlaylistByGroup(ctx context.Context, in *api.UpdatePlaylistByGroupRequest, opts ...grpc.CallOption) (*api.UpdatePlaylistByGroupResponse, error)
	// POST /v1/devices/groups/{group_id}/sync
	ActionSyncByGroup(ctx context.Context, in *api.ActionSyncByGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// POST /v1/devices/groups/{group_id}/test
	ActionTestByGroup(ctx context.Context, in *api.ActionTestByGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// POST /v1/devices/sync
	ActionSyncAll(ctx context.Context, in *api.ActionSyncAllRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// POST /v1/devices/{id}/sync
	ActionSync(ctx context.Context, in *api.ActionSyncRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// POST /v1/devices/test
	ActionTestAll(ctx context.Context, in *api.ActionTestAllRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// POST /v1/devices/{id}/test
	ActionTest(ctx context.Context, in *api.ActionTestRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type devicesServiceClient struct {
	cc *grpc.ClientConn
}

func NewDevicesServiceClient(cc *grpc.ClientConn) DevicesServiceClient {
	return &devicesServiceClient{cc}
}

func (c *devicesServiceClient) ListDevices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.ListDevicesResponse, error) {
	out := new(api.ListDevicesResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ListDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) CreateDevice(ctx context.Context, in *api.CreateDeviceRequest, opts ...grpc.CallOption) (*api.CreateDeviceResponse, error) {
	out := new(api.CreateDeviceResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) GetDevice(ctx context.Context, in *api.GetDeviceRequest, opts ...grpc.CallOption) (*api.GetDeviceResponse, error) {
	out := new(api.GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) UpdateDevice(ctx context.Context, in *api.UpdateDeviceRequest, opts ...grpc.CallOption) (*api.UpdateDeviceResponse, error) {
	out := new(api.UpdateDeviceResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) DeleteDevice(ctx context.Context, in *api.DeleteDeviceRequest, opts ...grpc.CallOption) (*api.DeleteDeviceResponse, error) {
	out := new(api.DeleteDeviceResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) GetPlaylist(ctx context.Context, in *api.GetPlaylistRequest, opts ...grpc.CallOption) (*api.GetPlaylistResponse, error) {
	out := new(api.GetPlaylistResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/GetPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) UpdatePlaylist(ctx context.Context, in *api.UpdatePlaylistRequest, opts ...grpc.CallOption) (*api.UpdatePlaylistResponse, error) {
	out := new(api.UpdatePlaylistResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/UpdatePlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ListGroups(ctx context.Context, in *api.ListGroupsRequest, opts ...grpc.CallOption) (*api.ListGroupsResponse, error) {
	out := new(api.ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ListDevicesByGroup(ctx context.Context, in *api.ListDevicesByGroupRequest, opts ...grpc.CallOption) (*api.ListDevicesByGroupResponse, error) {
	out := new(api.ListDevicesByGroupResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ListDevicesByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) UpdateDevicesByGroup(ctx context.Context, in *api.UpdateDevicesByGroupRequest, opts ...grpc.CallOption) (*api.UpdateDevicesByGroupResponse, error) {
	out := new(api.UpdateDevicesByGroupResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/UpdateDevicesByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) GetPlaylistByGroup(ctx context.Context, in *api.GetPlaylistByGroupRequest, opts ...grpc.CallOption) (*api.GetPlaylistByGroupResponse, error) {
	out := new(api.GetPlaylistByGroupResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/GetPlaylistByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) UpdatePlaylistByGroup(ctx context.Context, in *api.UpdatePlaylistByGroupRequest, opts ...grpc.CallOption) (*api.UpdatePlaylistByGroupResponse, error) {
	out := new(api.UpdatePlaylistByGroupResponse)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/UpdatePlaylistByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionSyncByGroup(ctx context.Context, in *api.ActionSyncByGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionSyncByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionTestByGroup(ctx context.Context, in *api.ActionTestByGroupRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionTestByGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionSyncAll(ctx context.Context, in *api.ActionSyncAllRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionSyncAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionSync(ctx context.Context, in *api.ActionSyncRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionTestAll(ctx context.Context, in *api.ActionTestAllRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionTestAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) ActionTest(ctx context.Context, in *api.ActionTestRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/adbus.api.DevicesService/ActionTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServiceServer is the server API for DevicesService service.
type DevicesServiceServer interface {
	// GET /v1/devices
	ListDevices(context.Context, *empty.Empty) (*api.ListDevicesResponse, error)
	// POST /v1/devices
	CreateDevice(context.Context, *api.CreateDeviceRequest) (*api.CreateDeviceResponse, error)
	// GET /v1/devices/{id}
	GetDevice(context.Context, *api.GetDeviceRequest) (*api.GetDeviceResponse, error)
	// PUT /v1/devices/{id}
	UpdateDevice(context.Context, *api.UpdateDeviceRequest) (*api.UpdateDeviceResponse, error)
	// DELETE /v1/devices/{id}
	DeleteDevice(context.Context, *api.DeleteDeviceRequest) (*api.DeleteDeviceResponse, error)
	// GET /v1/devices/{id}/playlist
	GetPlaylist(context.Context, *api.GetPlaylistRequest) (*api.GetPlaylistResponse, error)
	// PUT /v1/devices/{id}/playlist
	UpdatePlaylist(context.Context, *api.UpdatePlaylistRequest) (*api.UpdatePlaylistResponse, error)
	// GET /v1/devices/groups
	ListGroups(context.Context, *api.ListGroupsRequest) (*api.ListGroupsResponse, error)
	// GET /v1/devices/groups/{group_id}
	ListDevicesByGroup(context.Context, *api.ListDevicesByGroupRequest) (*api.ListDevicesByGroupResponse, error)
	// PUT /v1/devices/groups/{group_id}
	UpdateDevicesByGroup(context.Context, *api.UpdateDevicesByGroupRequest) (*api.UpdateDevicesByGroupResponse, error)
	// GET /v1/devices/groups/{group_id}/playlist
	GetPlaylistByGroup(context.Context, *api.GetPlaylistByGroupRequest) (*api.GetPlaylistByGroupResponse, error)
	// PUT /v1/devices/groups/{group_id}/playlist
	UpdatePlaylistByGroup(context.Context, *api.UpdatePlaylistByGroupRequest) (*api.UpdatePlaylistByGroupResponse, error)
	// POST /v1/devices/groups/{group_id}/sync
	ActionSyncByGroup(context.Context, *api.ActionSyncByGroupRequest) (*empty.Empty, error)
	// POST /v1/devices/groups/{group_id}/test
	ActionTestByGroup(context.Context, *api.ActionTestByGroupRequest) (*empty.Empty, error)
	// POST /v1/devices/sync
	ActionSyncAll(context.Context, *api.ActionSyncAllRequest) (*empty.Empty, error)
	// POST /v1/devices/{id}/sync
	ActionSync(context.Context, *api.ActionSyncRequest) (*empty.Empty, error)
	// POST /v1/devices/test
	ActionTestAll(context.Context, *api.ActionTestAllRequest) (*empty.Empty, error)
	// POST /v1/devices/{id}/test
	ActionTest(context.Context, *api.ActionTestRequest) (*empty.Empty, error)
}

// UnimplementedDevicesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDevicesServiceServer struct {
}

func (*UnimplementedDevicesServiceServer) ListDevices(ctx context.Context, req *empty.Empty) (*api.ListDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (*UnimplementedDevicesServiceServer) CreateDevice(ctx context.Context, req *api.CreateDeviceRequest) (*api.CreateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (*UnimplementedDevicesServiceServer) GetDevice(ctx context.Context, req *api.GetDeviceRequest) (*api.GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (*UnimplementedDevicesServiceServer) UpdateDevice(ctx context.Context, req *api.UpdateDeviceRequest) (*api.UpdateDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (*UnimplementedDevicesServiceServer) DeleteDevice(ctx context.Context, req *api.DeleteDeviceRequest) (*api.DeleteDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (*UnimplementedDevicesServiceServer) GetPlaylist(ctx context.Context, req *api.GetPlaylistRequest) (*api.GetPlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylist not implemented")
}
func (*UnimplementedDevicesServiceServer) UpdatePlaylist(ctx context.Context, req *api.UpdatePlaylistRequest) (*api.UpdatePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylist not implemented")
}
func (*UnimplementedDevicesServiceServer) ListGroups(ctx context.Context, req *api.ListGroupsRequest) (*api.ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (*UnimplementedDevicesServiceServer) ListDevicesByGroup(ctx context.Context, req *api.ListDevicesByGroupRequest) (*api.ListDevicesByGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevicesByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) UpdateDevicesByGroup(ctx context.Context, req *api.UpdateDevicesByGroupRequest) (*api.UpdateDevicesByGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevicesByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) GetPlaylistByGroup(ctx context.Context, req *api.GetPlaylistByGroupRequest) (*api.GetPlaylistByGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) UpdatePlaylistByGroup(ctx context.Context, req *api.UpdatePlaylistByGroupRequest) (*api.UpdatePlaylistByGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylistByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionSyncByGroup(ctx context.Context, req *api.ActionSyncByGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionSyncByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionTestByGroup(ctx context.Context, req *api.ActionTestByGroupRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionTestByGroup not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionSyncAll(ctx context.Context, req *api.ActionSyncAllRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionSyncAll not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionSync(ctx context.Context, req *api.ActionSyncRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionSync not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionTestAll(ctx context.Context, req *api.ActionTestAllRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionTestAll not implemented")
}
func (*UnimplementedDevicesServiceServer) ActionTest(ctx context.Context, req *api.ActionTestRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionTest not implemented")
}

func RegisterDevicesServiceServer(s *grpc.Server, srv DevicesServiceServer) {
	s.RegisterService(&_DevicesService_serviceDesc, srv)
}

func _DevicesService_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ListDevices(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).CreateDevice(ctx, req.(*api.CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetDevice(ctx, req.(*api.GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).UpdateDevice(ctx, req.(*api.UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).DeleteDevice(ctx, req.(*api.DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/GetPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetPlaylist(ctx, req.(*api.GetPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_UpdatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).UpdatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/UpdatePlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).UpdatePlaylist(ctx, req.(*api.UpdatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ListGroups(ctx, req.(*api.ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ListDevicesByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ListDevicesByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ListDevicesByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ListDevicesByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ListDevicesByGroup(ctx, req.(*api.ListDevicesByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_UpdateDevicesByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdateDevicesByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).UpdateDevicesByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/UpdateDevicesByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).UpdateDevicesByGroup(ctx, req.(*api.UpdateDevicesByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_GetPlaylistByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GetPlaylistByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetPlaylistByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/GetPlaylistByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetPlaylistByGroup(ctx, req.(*api.GetPlaylistByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_UpdatePlaylistByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.UpdatePlaylistByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).UpdatePlaylistByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/UpdatePlaylistByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).UpdatePlaylistByGroup(ctx, req.(*api.UpdatePlaylistByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionSyncByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionSyncByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionSyncByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionSyncByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionSyncByGroup(ctx, req.(*api.ActionSyncByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionTestByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionTestByGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionTestByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionTestByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionTestByGroup(ctx, req.(*api.ActionTestByGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionSyncAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionSyncAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionSyncAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionSyncAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionSyncAll(ctx, req.(*api.ActionSyncAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionSync(ctx, req.(*api.ActionSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionTestAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionTestAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionTestAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionTestAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionTestAll(ctx, req.(*api.ActionTestAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_ActionTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ActionTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).ActionTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adbus.api.DevicesService/ActionTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).ActionTest(ctx, req.(*api.ActionTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevicesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adbus.api.DevicesService",
	HandlerType: (*DevicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDevices",
			Handler:    _DevicesService_ListDevices_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _DevicesService_CreateDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _DevicesService_GetDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _DevicesService_UpdateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _DevicesService_DeleteDevice_Handler,
		},
		{
			MethodName: "GetPlaylist",
			Handler:    _DevicesService_GetPlaylist_Handler,
		},
		{
			MethodName: "UpdatePlaylist",
			Handler:    _DevicesService_UpdatePlaylist_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _DevicesService_ListGroups_Handler,
		},
		{
			MethodName: "ListDevicesByGroup",
			Handler:    _DevicesService_ListDevicesByGroup_Handler,
		},
		{
			MethodName: "UpdateDevicesByGroup",
			Handler:    _DevicesService_UpdateDevicesByGroup_Handler,
		},
		{
			MethodName: "GetPlaylistByGroup",
			Handler:    _DevicesService_GetPlaylistByGroup_Handler,
		},
		{
			MethodName: "UpdatePlaylistByGroup",
			Handler:    _DevicesService_UpdatePlaylistByGroup_Handler,
		},
		{
			MethodName: "ActionSyncByGroup",
			Handler:    _DevicesService_ActionSyncByGroup_Handler,
		},
		{
			MethodName: "ActionTestByGroup",
			Handler:    _DevicesService_ActionTestByGroup_Handler,
		},
		{
			MethodName: "ActionSyncAll",
			Handler:    _DevicesService_ActionSyncAll_Handler,
		},
		{
			MethodName: "ActionSync",
			Handler:    _DevicesService_ActionSync_Handler,
		},
		{
			MethodName: "ActionTestAll",
			Handler:    _DevicesService_ActionTestAll_Handler,
		},
		{
			MethodName: "ActionTest",
			Handler:    _DevicesService_ActionTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
