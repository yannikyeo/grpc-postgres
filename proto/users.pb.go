// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
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

type Role int32

const (
	Role_GUEST  Role = 0
	Role_MEMBER Role = 1
	Role_ADMIN  Role = 2
)

var Role_name = map[int32]string{
	0: "GUEST",
	1: "MEMBER",
	2: "ADMIN",
}

var Role_value = map[string]int32{
	"GUEST":  0,
	"MEMBER": 1,
	"ADMIN":  2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role                 `protobuf:"varint,2,opt,name=role,proto3,enum=users.Role" json:"role,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type UserRole struct {
	Role                 Role     `protobuf:"varint,1,opt,name=role,proto3,enum=users.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRole) Reset()         { *m = UserRole{} }
func (m *UserRole) String() string { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()    {}
func (*UserRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *UserRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRole.Unmarshal(m, b)
}
func (m *UserRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRole.Marshal(b, m, deterministic)
}
func (m *UserRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRole.Merge(m, src)
}
func (m *UserRole) XXX_Size() int {
	return xxx_messageInfo_UserRole.Size(m)
}
func (m *UserRole) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRole.DiscardUnknown(m)
}

var xxx_messageInfo_UserRole proto.InternalMessageInfo

func (m *UserRole) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

type AddUserRequest struct {
	Role                 Role     `protobuf:"varint,1,opt,name=role,proto3,enum=users.Role" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (m *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(m, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_GUEST
}

type DeleteUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListUsersRequest struct {
	// Only list users created after this timestamp
	CreatedSince *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_since,json=createdSince,proto3" json:"created_since,omitempty"`
	// Only list users older than this Duration
	OlderThan            *duration.Duration `protobuf:"bytes,2,opt,name=older_than,json=olderThan,proto3" json:"older_than,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetCreatedSince() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedSince
	}
	return nil
}

func (m *ListUsersRequest) GetOlderThan() *duration.Duration {
	if m != nil {
		return m.OlderThan
	}
	return nil
}

func init() {
	proto.RegisterEnum("users.Role", Role_name, Role_value)
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*UserRole)(nil), "users.UserRole")
	proto.RegisterType((*AddUserRequest)(nil), "users.AddUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "users.DeleteUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "users.ListUsersRequest")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xc7, 0xa5, 0x33, 0xd0, 0x1b, 0xa8, 0x8a, 0x25, 0x44, 0xe9, 0x82, 0xa9, 0xc2, 0xa6,
	0x02, 0x91, 0x40, 0xf8, 0x11, 0xd2, 0x2c, 0xd0, 0x8c, 0x5a, 0x21, 0x24, 0xca, 0x22, 0xed, 0x6c,
	0xd8, 0x54, 0x49, 0x7c, 0x49, 0x8c, 0xd2, 0x38, 0xd8, 0x0e, 0x6f, 0xc1, 0x9b, 0xf0, 0x90, 0xc8,
	0x76, 0x52, 0x20, 0x15, 0xb0, 0xf4, 0x3d, 0xe7, 0xf8, 0x5c, 0x7f, 0x09, 0x78, 0x8d, 0x42, 0xa9,
	0x82, 0x5a, 0x0a, 0x2d, 0xe8, 0xa9, 0x3d, 0xcc, 0xce, 0x73, 0x21, 0xf2, 0x12, 0x43, 0x3b, 0x4c,
	0x9b, 0xcf, 0xa1, 0xe6, 0x7b, 0x54, 0x3a, 0xd9, 0xd7, 0xce, 0x37, 0x7b, 0xd8, 0x37, 0xb0, 0x46,
	0x26, 0x9a, 0x8b, 0xca, 0xe9, 0xbe, 0x86, 0xe1, 0xb5, 0x42, 0x49, 0xc7, 0x30, 0xe0, 0x6c, 0x4a,
	0xe6, 0x64, 0x31, 0x8a, 0x07, 0x9c, 0xd1, 0x73, 0x18, 0x4a, 0x51, 0xe2, 0x74, 0x30, 0x27, 0x8b,
	0x71, 0xe4, 0x05, 0xae, 0x3b, 0x16, 0x25, 0xc6, 0x56, 0xa0, 0x17, 0xe0, 0x65, 0x12, 0x13, 0x8d,
	0x3b, 0x53, 0x39, 0xbd, 0x31, 0x27, 0x0b, 0x2f, 0x9a, 0x05, 0xae, 0x2e, 0xe8, 0xea, 0x82, 0x6d,
	0xb7, 0x4f, 0x0c, 0xce, 0x6e, 0x06, 0xfe, 0x13, 0xb8, 0x65, 0x5a, 0xcd, 0x75, 0x87, 0x26, 0xf2,
	0x97, 0x26, 0xff, 0x39, 0x8c, 0x2f, 0x19, 0xb3, 0x7e, 0xfc, 0xda, 0xa0, 0xd2, 0xff, 0x8f, 0x3c,
	0x82, 0xbb, 0x4b, 0x2c, 0x51, 0xe3, 0xef, 0xa9, 0xde, 0x13, 0xfd, 0xef, 0x04, 0x26, 0x1f, 0xb8,
	0xd2, 0xc6, 0xa3, 0x3a, 0xd3, 0x5b, 0xb8, 0xe3, 0xf6, 0x64, 0x3b, 0xc5, 0xab, 0xcc, 0x75, 0xfc,
	0xfb, 0x61, 0xb7, 0xdb, 0xc0, 0xc6, 0xf8, 0xe9, 0x1b, 0x00, 0x51, 0x32, 0x94, 0x3b, 0x5d, 0x24,
	0x95, 0xc5, 0xe7, 0x45, 0x0f, 0x8e, 0xd2, 0xcb, 0xf6, 0x2b, 0xc4, 0x23, 0x6b, 0xde, 0x16, 0x49,
	0xf5, 0x78, 0x01, 0x43, 0x0b, 0x64, 0x04, 0xa7, 0xef, 0xae, 0x57, 0x9b, 0xed, 0xe4, 0x84, 0x02,
	0x9c, 0xad, 0x57, 0xeb, 0xab, 0x55, 0x3c, 0x21, 0x66, 0x7c, 0xb9, 0x5c, 0xbf, 0xff, 0x38, 0x19,
	0x44, 0x3f, 0x08, 0x78, 0x66, 0xeb, 0x0d, 0xca, 0x6f, 0x3c, 0x43, 0x1a, 0xc2, 0xcd, 0x96, 0x10,
	0xbd, 0xd7, 0xc2, 0xf8, 0x93, 0xd8, 0xac, 0x63, 0x64, 0x66, 0xfe, 0x09, 0x7d, 0x05, 0xf0, 0x8b,
	0x0f, 0x9d, 0xb6, 0xe2, 0x11, 0xb2, 0xe3, 0xd8, 0xe8, 0x00, 0x8c, 0xde, 0x6f, 0xb5, 0x3e, 0xc2,
	0x5e, 0xe8, 0x19, 0xb9, 0x7a, 0xfd, 0xe9, 0x65, 0xce, 0x75, 0xd1, 0xa4, 0x41, 0x26, 0xf6, 0xe1,
	0x17, 0x51, 0x24, 0x55, 0x2a, 0x93, 0x8a, 0x15, 0x42, 0x2a, 0x1d, 0xe6, 0xb2, 0xce, 0x9e, 0xd6,
	0x42, 0xe9, 0x5c, 0xa2, 0x72, 0xff, 0xe9, 0x85, 0xbd, 0x20, 0x3d, 0xb3, 0x87, 0x17, 0x3f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x00, 0xa5, 0xf2, 0xd2, 0xf9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (UserService_ListUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/users.UserService/ListUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceListUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_ListUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceListUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceListUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*User, error)
	ListUsers(*ListUsersRequest, UserService_ListUsersServer) error
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListUsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).ListUsers(m, &userServiceListUsersServer{stream})
}

type UserService_ListUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceListUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceListUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListUsers",
			Handler:       _UserService_ListUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "users.proto",
}
