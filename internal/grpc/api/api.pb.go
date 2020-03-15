// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	Uuid                 string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	StartAt              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=startAt,proto3" json:"startAt,omitempty"`
	Duration             int64                `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	User                 string               `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	NotifyBefore         int64                `protobuf:"varint,7,opt,name=notifyBefore,proto3" json:"notifyBefore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetStartAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *Event) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetNotifyBefore() int64 {
	if m != nil {
		return m.NotifyBefore
	}
	return 0
}

type Day struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Day) Reset()         { *m = Day{} }
func (m *Day) String() string { return proto.CompactTextString(m) }
func (*Day) ProtoMessage()    {}
func (*Day) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *Day) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Day.Unmarshal(m, b)
}
func (m *Day) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Day.Marshal(b, m, deterministic)
}
func (m *Day) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Day.Merge(m, src)
}
func (m *Day) XXX_Size() int {
	return xxx_messageInfo_Day.Size(m)
}
func (m *Day) XXX_DiscardUnknown() {
	xxx_messageInfo_Day.DiscardUnknown(m)
}

var xxx_messageInfo_Day proto.InternalMessageInfo

func (m *Day) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type EventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventsResponse) Reset()         { *m = EventsResponse{} }
func (m *EventsResponse) String() string { return proto.CompactTextString(m) }
func (*EventsResponse) ProtoMessage()    {}
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *EventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsResponse.Unmarshal(m, b)
}
func (m *EventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsResponse.Marshal(b, m, deterministic)
}
func (m *EventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsResponse.Merge(m, src)
}
func (m *EventsResponse) XXX_Size() int {
	return xxx_messageInfo_EventsResponse.Size(m)
}
func (m *EventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventsResponse proto.InternalMessageInfo

func (m *EventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Day)(nil), "Day")
	proto.RegisterType((*EventsResponse)(nil), "EventsResponse")
}

func init() {
	proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784)
}

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x86, 0xe5, 0x4d, 0x9a, 0xee, 0x4e, 0x77, 0xb7, 0x60, 0x71, 0xb0, 0x72, 0x80, 0x28, 0xa7,
	0xc0, 0xc1, 0x45, 0x05, 0x1e, 0x00, 0x54, 0xe0, 0xc4, 0x25, 0x42, 0xe2, 0xec, 0x92, 0x69, 0xb1,
	0x68, 0xe3, 0xc8, 0x9e, 0x20, 0xf5, 0x11, 0x78, 0x48, 0xde, 0x05, 0x31, 0x69, 0x51, 0x41, 0x2a,
	0x37, 0xff, 0xbf, 0x3f, 0xcf, 0xff, 0x8f, 0xe1, 0x9f, 0x69, 0xec, 0xc8, 0x34, 0x56, 0x37, 0xde,
	0x91, 0x4b, 0x8f, 0xe6, 0xce, 0xcd, 0x17, 0x38, 0x62, 0x35, 0x6d, 0x67, 0x23, 0xb2, 0x4b, 0x0c,
	0x64, 0x96, 0x4d, 0x07, 0xe4, 0x6f, 0x02, 0x7a, 0xd7, 0x2f, 0x58, 0x93, 0x94, 0x10, 0xb7, 0xad,
	0xad, 0x94, 0xc8, 0x44, 0xf1, 0xa7, 0xe4, 0xb3, 0x3c, 0x80, 0x1e, 0x59, 0x5a, 0xa0, 0xfa, 0xc5,
	0x66, 0x27, 0xe4, 0x39, 0xf4, 0x03, 0x19, 0x4f, 0x97, 0xa4, 0xa2, 0x4c, 0x14, 0x83, 0x71, 0xaa,
	0xbb, 0x18, 0xbd, 0x89, 0xd1, 0xf7, 0x9b, 0x98, 0x72, 0x83, 0xca, 0x14, 0x7e, 0x57, 0xad, 0x37,
	0x64, 0x5d, 0xad, 0xe2, 0x4c, 0x14, 0x51, 0xf9, 0xa9, 0x65, 0x06, 0x83, 0x0a, 0xc3, 0xa3, 0xb7,
	0x0d, 0x5f, 0xf7, 0x38, 0x6d, 0xdb, 0xe2, 0x76, 0x01, 0xbd, 0x4a, 0xd6, 0xed, 0x02, 0x7a, 0x99,
	0xc3, 0xdf, 0xda, 0x91, 0x9d, 0xad, 0xae, 0x70, 0xe6, 0x3c, 0xaa, 0x3e, 0x4f, 0xfd, 0xe2, 0xe5,
	0x17, 0x10, 0x4d, 0xcc, 0x4a, 0x6a, 0x88, 0x2b, 0x43, 0xc8, 0xcb, 0xfd, 0xdc, 0x97, 0xb9, 0xfc,
	0x14, 0xfe, 0xf3, 0xaf, 0x84, 0x12, 0x43, 0xe3, 0xea, 0x80, 0xf2, 0x10, 0x12, 0x64, 0x47, 0x89,
	0x2c, 0x2a, 0x06, 0xe3, 0x44, 0x33, 0x50, 0xae, 0xdd, 0xf1, 0xab, 0x80, 0xa4, 0x7b, 0x22, 0x0b,
	0x18, 0xde, 0x22, 0x75, 0xe2, 0xc6, 0xf9, 0x8f, 0xfc, 0x58, 0x4f, 0xcc, 0x2a, 0x1d, 0xea, 0x6f,
	0x43, 0x8f, 0x61, 0x6f, 0x9b, 0x7c, 0x40, 0x7c, 0xde, 0x85, 0x9e, 0xc0, 0xfe, 0x36, 0x7a, 0xe7,
	0x6a, 0x7a, 0xda, 0xc1, 0x4e, 0x13, 0xde, 0xeb, 0xec, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x7e,
	0x38, 0xeb, 0x0d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	GetEventsForDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
	GetEventsForMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) GetEventsForDay(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/Events/GetEventsForDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) GetEventsForWeek(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/Events/GetEventsForWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) GetEventsForMonth(ctx context.Context, in *Day, opts ...grpc.CallOption) (*EventsResponse, error) {
	out := new(EventsResponse)
	err := c.cc.Invoke(ctx, "/Events/GetEventsForMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	GetEventsForDay(context.Context, *Day) (*EventsResponse, error)
	GetEventsForWeek(context.Context, *Day) (*EventsResponse, error)
	GetEventsForMonth(context.Context, *Day) (*EventsResponse, error)
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) GetEventsForDay(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForDay not implemented")
}
func (*UnimplementedEventsServer) GetEventsForWeek(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForWeek not implemented")
}
func (*UnimplementedEventsServer) GetEventsForMonth(ctx context.Context, req *Day) (*EventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventsForMonth not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_GetEventsForDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).GetEventsForDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/GetEventsForDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).GetEventsForDay(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_GetEventsForWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).GetEventsForWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/GetEventsForWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).GetEventsForWeek(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_GetEventsForMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).GetEventsForMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Events/GetEventsForMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).GetEventsForMonth(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventsForDay",
			Handler:    _Events_GetEventsForDay_Handler,
		},
		{
			MethodName: "GetEventsForWeek",
			Handler:    _Events_GetEventsForWeek_Handler,
		},
		{
			MethodName: "GetEventsForMonth",
			Handler:    _Events_GetEventsForMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}