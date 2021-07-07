// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pkg/nfnNotify/proto/nfn.proto

package nfn

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SubscribeContext struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeContext) Reset()         { *m = SubscribeContext{} }
func (m *SubscribeContext) String() string { return proto.CompactTextString(m) }
func (*SubscribeContext) ProtoMessage()    {}
func (*SubscribeContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{0}
}

func (m *SubscribeContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeContext.Unmarshal(m, b)
}
func (m *SubscribeContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeContext.Marshal(b, m, deterministic)
}
func (m *SubscribeContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeContext.Merge(m, src)
}
func (m *SubscribeContext) XXX_Size() int {
	return xxx_messageInfo_SubscribeContext.Size(m)
}
func (m *SubscribeContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeContext.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeContext proto.InternalMessageInfo

func (m *SubscribeContext) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Notification struct {
	CniType string `protobuf:"bytes,1,opt,name=cni_type,json=cniType,proto3" json:"cni_type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Notification_InSync
	//	*Notification_ProviderNwCreate
	//	*Notification_ProviderNwRemove
	//	*Notification_ContainterRtInsert
	//	*Notification_ContainterRtRemove
	//	*Notification_PodAddNetwork
	//	*Notification_PodDelNetwork
	Payload              isNotification_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetCniType() string {
	if m != nil {
		return m.CniType
	}
	return ""
}

type isNotification_Payload interface {
	isNotification_Payload()
}

type Notification_InSync struct {
	InSync *InSync `protobuf:"bytes,2,opt,name=in_sync,json=inSync,proto3,oneof"`
}

type Notification_ProviderNwCreate struct {
	ProviderNwCreate *ProviderNetworkCreate `protobuf:"bytes,3,opt,name=provider_nw_create,json=providerNwCreate,proto3,oneof"`
}

type Notification_ProviderNwRemove struct {
	ProviderNwRemove *ProviderNetworkRemove `protobuf:"bytes,4,opt,name=provider_nw_remove,json=providerNwRemove,proto3,oneof"`
}

type Notification_ContainterRtInsert struct {
	ContainterRtInsert *ContainerRouteInsert `protobuf:"bytes,5,opt,name=containter_rt_insert,json=containterRtInsert,proto3,oneof"`
}

type Notification_ContainterRtRemove struct {
	ContainterRtRemove *ContainerRouteRemove `protobuf:"bytes,6,opt,name=containter_rt_remove,json=containterRtRemove,proto3,oneof"`
}

type Notification_PodAddNetwork struct {
	PodAddNetwork *PodAddNetwork `protobuf:"bytes,7,opt,name=pod_add_network,json=podAddNetwork,proto3,oneof"`
}

type Notification_PodDelNetwork struct {
	PodDelNetwork *PodDelNetwork `protobuf:"bytes,8,opt,name=pod_del_network,json=podDelNetwork,proto3,oneof"`
}

func (*Notification_InSync) isNotification_Payload() {}

func (*Notification_ProviderNwCreate) isNotification_Payload() {}

func (*Notification_ProviderNwRemove) isNotification_Payload() {}

func (*Notification_ContainterRtInsert) isNotification_Payload() {}

func (*Notification_ContainterRtRemove) isNotification_Payload() {}

func (*Notification_PodAddNetwork) isNotification_Payload() {}

func (*Notification_PodDelNetwork) isNotification_Payload() {}

func (m *Notification) GetPayload() isNotification_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Notification) GetInSync() *InSync {
	if x, ok := m.GetPayload().(*Notification_InSync); ok {
		return x.InSync
	}
	return nil
}

func (m *Notification) GetProviderNwCreate() *ProviderNetworkCreate {
	if x, ok := m.GetPayload().(*Notification_ProviderNwCreate); ok {
		return x.ProviderNwCreate
	}
	return nil
}

func (m *Notification) GetProviderNwRemove() *ProviderNetworkRemove {
	if x, ok := m.GetPayload().(*Notification_ProviderNwRemove); ok {
		return x.ProviderNwRemove
	}
	return nil
}

func (m *Notification) GetContainterRtInsert() *ContainerRouteInsert {
	if x, ok := m.GetPayload().(*Notification_ContainterRtInsert); ok {
		return x.ContainterRtInsert
	}
	return nil
}

func (m *Notification) GetContainterRtRemove() *ContainerRouteRemove {
	if x, ok := m.GetPayload().(*Notification_ContainterRtRemove); ok {
		return x.ContainterRtRemove
	}
	return nil
}

func (m *Notification) GetPodAddNetwork() *PodAddNetwork {
	if x, ok := m.GetPayload().(*Notification_PodAddNetwork); ok {
		return x.PodAddNetwork
	}
	return nil
}

func (m *Notification) GetPodDelNetwork() *PodDelNetwork {
	if x, ok := m.GetPayload().(*Notification_PodDelNetwork); ok {
		return x.PodDelNetwork
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Notification) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Notification_InSync)(nil),
		(*Notification_ProviderNwCreate)(nil),
		(*Notification_ProviderNwRemove)(nil),
		(*Notification_ContainterRtInsert)(nil),
		(*Notification_ContainterRtRemove)(nil),
		(*Notification_PodAddNetwork)(nil),
		(*Notification_PodDelNetwork)(nil),
	}
}

type ProviderNetworkCreate struct {
	ProviderNwName       string      `protobuf:"bytes,1,opt,name=provider_nw_name,json=providerNwName,proto3" json:"provider_nw_name,omitempty"`
	Vlan                 *VlanInfo   `protobuf:"bytes,2,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Direct               *DirectInfo `protobuf:"bytes,3,opt,name=direct,proto3" json:"direct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProviderNetworkCreate) Reset()         { *m = ProviderNetworkCreate{} }
func (m *ProviderNetworkCreate) String() string { return proto.CompactTextString(m) }
func (*ProviderNetworkCreate) ProtoMessage()    {}
func (*ProviderNetworkCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{2}
}

func (m *ProviderNetworkCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderNetworkCreate.Unmarshal(m, b)
}
func (m *ProviderNetworkCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderNetworkCreate.Marshal(b, m, deterministic)
}
func (m *ProviderNetworkCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderNetworkCreate.Merge(m, src)
}
func (m *ProviderNetworkCreate) XXX_Size() int {
	return xxx_messageInfo_ProviderNetworkCreate.Size(m)
}
func (m *ProviderNetworkCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderNetworkCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderNetworkCreate proto.InternalMessageInfo

func (m *ProviderNetworkCreate) GetProviderNwName() string {
	if m != nil {
		return m.ProviderNwName
	}
	return ""
}

func (m *ProviderNetworkCreate) GetVlan() *VlanInfo {
	if m != nil {
		return m.Vlan
	}
	return nil
}

func (m *ProviderNetworkCreate) GetDirect() *DirectInfo {
	if m != nil {
		return m.Direct
	}
	return nil
}

type ProviderNetworkRemove struct {
	ProviderNwName       string   `protobuf:"bytes,1,opt,name=provider_nw_name,json=providerNwName,proto3" json:"provider_nw_name,omitempty"`
	VlanLogicalIntf      string   `protobuf:"bytes,2,opt,name=vlan_logical_intf,json=vlanLogicalIntf,proto3" json:"vlan_logical_intf,omitempty"`
	DirectProviderIntf   string   `protobuf:"bytes,3,opt,name=direct_provider_intf,json=directProviderIntf,proto3" json:"direct_provider_intf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderNetworkRemove) Reset()         { *m = ProviderNetworkRemove{} }
func (m *ProviderNetworkRemove) String() string { return proto.CompactTextString(m) }
func (*ProviderNetworkRemove) ProtoMessage()    {}
func (*ProviderNetworkRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{3}
}

func (m *ProviderNetworkRemove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderNetworkRemove.Unmarshal(m, b)
}
func (m *ProviderNetworkRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderNetworkRemove.Marshal(b, m, deterministic)
}
func (m *ProviderNetworkRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderNetworkRemove.Merge(m, src)
}
func (m *ProviderNetworkRemove) XXX_Size() int {
	return xxx_messageInfo_ProviderNetworkRemove.Size(m)
}
func (m *ProviderNetworkRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderNetworkRemove.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderNetworkRemove proto.InternalMessageInfo

func (m *ProviderNetworkRemove) GetProviderNwName() string {
	if m != nil {
		return m.ProviderNwName
	}
	return ""
}

func (m *ProviderNetworkRemove) GetVlanLogicalIntf() string {
	if m != nil {
		return m.VlanLogicalIntf
	}
	return ""
}

func (m *ProviderNetworkRemove) GetDirectProviderIntf() string {
	if m != nil {
		return m.DirectProviderIntf
	}
	return ""
}

type VlanInfo struct {
	VlanId               string   `protobuf:"bytes,1,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	ProviderIntf         string   `protobuf:"bytes,2,opt,name=provider_intf,json=providerIntf,proto3" json:"provider_intf,omitempty"`
	LogicalIntf          string   `protobuf:"bytes,3,opt,name=logical_intf,json=logicalIntf,proto3" json:"logical_intf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlanInfo) Reset()         { *m = VlanInfo{} }
func (m *VlanInfo) String() string { return proto.CompactTextString(m) }
func (*VlanInfo) ProtoMessage()    {}
func (*VlanInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{4}
}

func (m *VlanInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanInfo.Unmarshal(m, b)
}
func (m *VlanInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanInfo.Marshal(b, m, deterministic)
}
func (m *VlanInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanInfo.Merge(m, src)
}
func (m *VlanInfo) XXX_Size() int {
	return xxx_messageInfo_VlanInfo.Size(m)
}
func (m *VlanInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VlanInfo proto.InternalMessageInfo

func (m *VlanInfo) GetVlanId() string {
	if m != nil {
		return m.VlanId
	}
	return ""
}

func (m *VlanInfo) GetProviderIntf() string {
	if m != nil {
		return m.ProviderIntf
	}
	return ""
}

func (m *VlanInfo) GetLogicalIntf() string {
	if m != nil {
		return m.LogicalIntf
	}
	return ""
}

type DirectInfo struct {
	ProviderIntf         string   `protobuf:"bytes,1,opt,name=provider_intf,json=providerIntf,proto3" json:"provider_intf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectInfo) Reset()         { *m = DirectInfo{} }
func (m *DirectInfo) String() string { return proto.CompactTextString(m) }
func (*DirectInfo) ProtoMessage()    {}
func (*DirectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{5}
}

func (m *DirectInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectInfo.Unmarshal(m, b)
}
func (m *DirectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectInfo.Marshal(b, m, deterministic)
}
func (m *DirectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectInfo.Merge(m, src)
}
func (m *DirectInfo) XXX_Size() int {
	return xxx_messageInfo_DirectInfo.Size(m)
}
func (m *DirectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DirectInfo proto.InternalMessageInfo

func (m *DirectInfo) GetProviderIntf() string {
	if m != nil {
		return m.ProviderIntf
	}
	return ""
}

type RouteData struct {
	Dst                  string   `protobuf:"bytes,2,opt,name=dst,proto3" json:"dst,omitempty"`
	Gw                   string   `protobuf:"bytes,3,opt,name=gw,proto3" json:"gw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteData) Reset()         { *m = RouteData{} }
func (m *RouteData) String() string { return proto.CompactTextString(m) }
func (*RouteData) ProtoMessage()    {}
func (*RouteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{6}
}

func (m *RouteData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteData.Unmarshal(m, b)
}
func (m *RouteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteData.Marshal(b, m, deterministic)
}
func (m *RouteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteData.Merge(m, src)
}
func (m *RouteData) XXX_Size() int {
	return xxx_messageInfo_RouteData.Size(m)
}
func (m *RouteData) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteData.DiscardUnknown(m)
}

var xxx_messageInfo_RouteData proto.InternalMessageInfo

func (m *RouteData) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *RouteData) GetGw() string {
	if m != nil {
		return m.Gw
	}
	return ""
}

type ContainerRouteInsert struct {
	ContainerId          string       `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Route                []*RouteData `protobuf:"bytes,2,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContainerRouteInsert) Reset()         { *m = ContainerRouteInsert{} }
func (m *ContainerRouteInsert) String() string { return proto.CompactTextString(m) }
func (*ContainerRouteInsert) ProtoMessage()    {}
func (*ContainerRouteInsert) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{7}
}

func (m *ContainerRouteInsert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRouteInsert.Unmarshal(m, b)
}
func (m *ContainerRouteInsert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRouteInsert.Marshal(b, m, deterministic)
}
func (m *ContainerRouteInsert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRouteInsert.Merge(m, src)
}
func (m *ContainerRouteInsert) XXX_Size() int {
	return xxx_messageInfo_ContainerRouteInsert.Size(m)
}
func (m *ContainerRouteInsert) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRouteInsert.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRouteInsert proto.InternalMessageInfo

func (m *ContainerRouteInsert) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *ContainerRouteInsert) GetRoute() []*RouteData {
	if m != nil {
		return m.Route
	}
	return nil
}

type ContainerRouteRemove struct {
	ContainerId          string       `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Route                []*RouteData `protobuf:"bytes,2,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContainerRouteRemove) Reset()         { *m = ContainerRouteRemove{} }
func (m *ContainerRouteRemove) String() string { return proto.CompactTextString(m) }
func (*ContainerRouteRemove) ProtoMessage()    {}
func (*ContainerRouteRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{8}
}

func (m *ContainerRouteRemove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRouteRemove.Unmarshal(m, b)
}
func (m *ContainerRouteRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRouteRemove.Marshal(b, m, deterministic)
}
func (m *ContainerRouteRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRouteRemove.Merge(m, src)
}
func (m *ContainerRouteRemove) XXX_Size() int {
	return xxx_messageInfo_ContainerRouteRemove.Size(m)
}
func (m *ContainerRouteRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRouteRemove.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRouteRemove proto.InternalMessageInfo

func (m *ContainerRouteRemove) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *ContainerRouteRemove) GetRoute() []*RouteData {
	if m != nil {
		return m.Route
	}
	return nil
}

type PodInfo struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PodInfo) Reset()         { *m = PodInfo{} }
func (m *PodInfo) String() string { return proto.CompactTextString(m) }
func (*PodInfo) ProtoMessage()    {}
func (*PodInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{9}
}

func (m *PodInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodInfo.Unmarshal(m, b)
}
func (m *PodInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodInfo.Marshal(b, m, deterministic)
}
func (m *PodInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodInfo.Merge(m, src)
}
func (m *PodInfo) XXX_Size() int {
	return xxx_messageInfo_PodInfo.Size(m)
}
func (m *PodInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PodInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PodInfo proto.InternalMessageInfo

func (m *PodInfo) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PodInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NetConf struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetConf) Reset()         { *m = NetConf{} }
func (m *NetConf) String() string { return proto.CompactTextString(m) }
func (*NetConf) ProtoMessage()    {}
func (*NetConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{10}
}

func (m *NetConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetConf.Unmarshal(m, b)
}
func (m *NetConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetConf.Marshal(b, m, deterministic)
}
func (m *NetConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetConf.Merge(m, src)
}
func (m *NetConf) XXX_Size() int {
	return xxx_messageInfo_NetConf.Size(m)
}
func (m *NetConf) XXX_DiscardUnknown() {
	xxx_messageInfo_NetConf.DiscardUnknown(m)
}

var xxx_messageInfo_NetConf proto.InternalMessageInfo

func (m *NetConf) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PodAddNetwork struct {
	ContainerId          string       `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Pod                  *PodInfo     `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	Net                  *NetConf     `protobuf:"bytes,3,opt,name=net,proto3" json:"net,omitempty"`
	Route                []*RouteData `protobuf:"bytes,4,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PodAddNetwork) Reset()         { *m = PodAddNetwork{} }
func (m *PodAddNetwork) String() string { return proto.CompactTextString(m) }
func (*PodAddNetwork) ProtoMessage()    {}
func (*PodAddNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{11}
}

func (m *PodAddNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodAddNetwork.Unmarshal(m, b)
}
func (m *PodAddNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodAddNetwork.Marshal(b, m, deterministic)
}
func (m *PodAddNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodAddNetwork.Merge(m, src)
}
func (m *PodAddNetwork) XXX_Size() int {
	return xxx_messageInfo_PodAddNetwork.Size(m)
}
func (m *PodAddNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_PodAddNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_PodAddNetwork proto.InternalMessageInfo

func (m *PodAddNetwork) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *PodAddNetwork) GetPod() *PodInfo {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *PodAddNetwork) GetNet() *NetConf {
	if m != nil {
		return m.Net
	}
	return nil
}

func (m *PodAddNetwork) GetRoute() []*RouteData {
	if m != nil {
		return m.Route
	}
	return nil
}

type PodDelNetwork struct {
	ContainerId          string       `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Pod                  *PodInfo     `protobuf:"bytes,2,opt,name=pod,proto3" json:"pod,omitempty"`
	Net                  *NetConf     `protobuf:"bytes,3,opt,name=net,proto3" json:"net,omitempty"`
	Route                []*RouteData `protobuf:"bytes,4,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PodDelNetwork) Reset()         { *m = PodDelNetwork{} }
func (m *PodDelNetwork) String() string { return proto.CompactTextString(m) }
func (*PodDelNetwork) ProtoMessage()    {}
func (*PodDelNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{12}
}

func (m *PodDelNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodDelNetwork.Unmarshal(m, b)
}
func (m *PodDelNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodDelNetwork.Marshal(b, m, deterministic)
}
func (m *PodDelNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodDelNetwork.Merge(m, src)
}
func (m *PodDelNetwork) XXX_Size() int {
	return xxx_messageInfo_PodDelNetwork.Size(m)
}
func (m *PodDelNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_PodDelNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_PodDelNetwork proto.InternalMessageInfo

func (m *PodDelNetwork) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *PodDelNetwork) GetPod() *PodInfo {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *PodDelNetwork) GetNet() *NetConf {
	if m != nil {
		return m.Net
	}
	return nil
}

func (m *PodDelNetwork) GetRoute() []*RouteData {
	if m != nil {
		return m.Route
	}
	return nil
}

type InSync struct {
	NodeIntfIpAddress    string   `protobuf:"bytes,1,opt,name=node_intf_ip_address,json=nodeIntfIpAddress,proto3" json:"node_intf_ip_address,omitempty"`
	NodeIntfMacAddress   string   `protobuf:"bytes,2,opt,name=node_intf_mac_address,json=nodeIntfMacAddress,proto3" json:"node_intf_mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InSync) Reset()         { *m = InSync{} }
func (m *InSync) String() string { return proto.CompactTextString(m) }
func (*InSync) ProtoMessage()    {}
func (*InSync) Descriptor() ([]byte, []int) {
	return fileDescriptor_a34e3338d09121c4, []int{13}
}

func (m *InSync) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InSync.Unmarshal(m, b)
}
func (m *InSync) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InSync.Marshal(b, m, deterministic)
}
func (m *InSync) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InSync.Merge(m, src)
}
func (m *InSync) XXX_Size() int {
	return xxx_messageInfo_InSync.Size(m)
}
func (m *InSync) XXX_DiscardUnknown() {
	xxx_messageInfo_InSync.DiscardUnknown(m)
}

var xxx_messageInfo_InSync proto.InternalMessageInfo

func (m *InSync) GetNodeIntfIpAddress() string {
	if m != nil {
		return m.NodeIntfIpAddress
	}
	return ""
}

func (m *InSync) GetNodeIntfMacAddress() string {
	if m != nil {
		return m.NodeIntfMacAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*SubscribeContext)(nil), "SubscribeContext")
	proto.RegisterType((*Notification)(nil), "Notification")
	proto.RegisterType((*ProviderNetworkCreate)(nil), "ProviderNetworkCreate")
	proto.RegisterType((*ProviderNetworkRemove)(nil), "ProviderNetworkRemove")
	proto.RegisterType((*VlanInfo)(nil), "VlanInfo")
	proto.RegisterType((*DirectInfo)(nil), "DirectInfo")
	proto.RegisterType((*RouteData)(nil), "RouteData")
	proto.RegisterType((*ContainerRouteInsert)(nil), "ContainerRouteInsert")
	proto.RegisterType((*ContainerRouteRemove)(nil), "ContainerRouteRemove")
	proto.RegisterType((*PodInfo)(nil), "PodInfo")
	proto.RegisterType((*NetConf)(nil), "NetConf")
	proto.RegisterType((*PodAddNetwork)(nil), "PodAddNetwork")
	proto.RegisterType((*PodDelNetwork)(nil), "PodDelNetwork")
	proto.RegisterType((*InSync)(nil), "InSync")
}

func init() {
	proto.RegisterFile("internal/pkg/nfnNotify/proto/nfn.proto", fileDescriptor_a34e3338d09121c4)
}

var fileDescriptor_a34e3338d09121c4 = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xdf, 0x6e, 0xd3, 0x3e,
	0x18, 0xed, 0xbf, 0x35, 0xcd, 0xd7, 0x75, 0xeb, 0xac, 0xee, 0xf7, 0x2b, 0x83, 0x49, 0x25, 0x93,
	0x50, 0x85, 0x44, 0xba, 0x8d, 0x1b, 0x24, 0xb8, 0x19, 0x9d, 0x50, 0x23, 0x41, 0x35, 0x65, 0x88,
	0x1b, 0x2e, 0x22, 0x2f, 0x76, 0x2b, 0x6b, 0xa9, 0x1d, 0xa5, 0xde, 0x4a, 0x1f, 0x80, 0x3b, 0xde,
	0x81, 0x77, 0xe4, 0x09, 0x90, 0x1d, 0x27, 0x69, 0xd7, 0x22, 0xed, 0x02, 0x89, 0x3b, 0xe7, 0x1c,
	0x9f, 0xf3, 0x1d, 0x7f, 0xfe, 0x5a, 0xc3, 0x0b, 0xc6, 0x25, 0x4d, 0x38, 0x8e, 0x06, 0xf1, 0xed,
	0x74, 0xc0, 0x27, 0x7c, 0x2c, 0x24, 0x9b, 0x2c, 0x07, 0x71, 0x22, 0xa4, 0x50, 0xdf, 0xae, 0x5e,
	0x39, 0x03, 0x68, 0x5f, 0xdf, 0xdd, 0xcc, 0xc3, 0x84, 0xdd, 0xd0, 0xa1, 0xe0, 0x92, 0x7e, 0x93,
	0xe8, 0x29, 0xd8, 0x5c, 0x10, 0x1a, 0x70, 0x3c, 0xa3, 0xdd, 0x72, 0xaf, 0xdc, 0xb7, 0xfd, 0x86,
	0x02, 0xc6, 0x78, 0x46, 0x9d, 0x5f, 0x55, 0xd8, 0xd5, 0x5e, 0x2c, 0xc4, 0x92, 0x09, 0x8e, 0x9e,
	0x40, 0x23, 0xe4, 0x2c, 0x90, 0xcb, 0x38, 0xdb, 0x6c, 0x85, 0x9c, 0x7d, 0x5e, 0xc6, 0x14, 0x39,
	0x60, 0x31, 0x1e, 0xcc, 0x97, 0x3c, 0xec, 0x56, 0x7a, 0xe5, 0x7e, 0xf3, 0xdc, 0x72, 0x3d, 0x7e,
	0xbd, 0xe4, 0xe1, 0xa8, 0xe4, 0xd7, 0x99, 0x5e, 0xa1, 0x0f, 0x80, 0xe2, 0x44, 0xdc, 0x33, 0x42,
	0x93, 0x80, 0x2f, 0x82, 0x30, 0xa1, 0x58, 0xd2, 0x6e, 0x55, 0x6f, 0xff, 0xcf, 0xbd, 0x32, 0xd4,
	0x98, 0xca, 0x85, 0x48, 0x6e, 0x87, 0x9a, 0x1d, 0x95, 0xfc, 0x76, 0xa6, 0x19, 0x2f, 0x52, 0xec,
	0xa1, 0x4f, 0x42, 0x67, 0xe2, 0x9e, 0x76, 0x6b, 0xdb, 0x7d, 0x7c, 0xcd, 0xae, 0xfb, 0xa4, 0x18,
	0xf2, 0xa0, 0x13, 0x0a, 0x2e, 0xb1, 0xee, 0x5f, 0x90, 0xc8, 0x80, 0xf1, 0x39, 0x4d, 0x64, 0x77,
	0x47, 0x3b, 0x1d, 0xba, 0xc3, 0x94, 0xa4, 0x89, 0x2f, 0xee, 0x24, 0xf5, 0x34, 0x39, 0x2a, 0xf9,
	0xa8, 0x10, 0xf9, 0x32, 0x45, 0x37, 0xad, 0x4c, 0xa8, 0xfa, 0x56, 0xab, 0x3c, 0xd3, 0x9a, 0x95,
	0x49, 0xf5, 0x06, 0xf6, 0x63, 0x41, 0x02, 0x4c, 0x48, 0xc0, 0xd3, 0x23, 0x74, 0x2d, 0xed, 0xb2,
	0xe7, 0x5e, 0x09, 0x72, 0x41, 0x88, 0x39, 0xd8, 0xa8, 0xe4, 0xb7, 0xe2, 0x55, 0x20, 0x53, 0x12,
	0x1a, 0xe5, 0xca, 0x46, 0xa1, 0xbc, 0xa4, 0xd1, 0xba, 0xb2, 0x00, 0xde, 0xdb, 0x60, 0xc5, 0x78,
	0x19, 0x09, 0x4c, 0x9c, 0xef, 0x65, 0x38, 0xdc, 0x7a, 0x15, 0xa8, 0x0f, 0xed, 0xd5, 0xb6, 0xaf,
	0x8c, 0xcc, 0x5e, 0xd1, 0x5a, 0x35, 0x38, 0xe8, 0x18, 0x6a, 0xf7, 0x11, 0xe6, 0x66, 0x12, 0x6c,
	0xf7, 0x4b, 0x84, 0xb9, 0xc7, 0x27, 0xc2, 0xd7, 0x30, 0x3a, 0x81, 0x3a, 0x61, 0x09, 0x0d, 0xa5,
	0xb9, 0xfb, 0xa6, 0x7b, 0xa9, 0x3f, 0xf5, 0x16, 0x43, 0x39, 0x3f, 0x37, 0x73, 0x98, 0x06, 0x3d,
	0x3e, 0xc7, 0x4b, 0x38, 0x50, 0x05, 0x83, 0x48, 0x4c, 0x59, 0x88, 0xa3, 0x80, 0x71, 0x39, 0xd1,
	0xa1, 0x6c, 0x7f, 0x5f, 0x11, 0x1f, 0x53, 0xdc, 0xe3, 0x72, 0x82, 0x4e, 0xa1, 0x93, 0x56, 0x0e,
	0x72, 0x73, 0xbd, 0xbd, 0xaa, 0xb7, 0xa3, 0x94, 0xcb, 0x02, 0x29, 0x85, 0x73, 0x0b, 0x8d, 0xec,
	0x60, 0xe8, 0x7f, 0xb0, 0x74, 0x25, 0x46, 0x4c, 0x94, 0xba, 0xfa, 0xf4, 0x08, 0x3a, 0x81, 0xd6,
	0xba, 0x5f, 0x5a, 0x7e, 0x37, 0x5e, 0x71, 0x42, 0xcf, 0x61, 0x77, 0x2d, 0x62, 0x5a, 0xb3, 0x19,
	0x15, 0xf1, 0x9c, 0x33, 0x80, 0xa2, 0x49, 0x9b, 0xae, 0xe5, 0x4d, 0x57, 0xe7, 0x15, 0xd8, 0x7a,
	0xda, 0x2e, 0xb1, 0xc4, 0xa8, 0x0d, 0x55, 0x32, 0x97, 0xa6, 0xba, 0x5a, 0xa2, 0x3d, 0xa8, 0x4c,
	0x17, 0xa6, 0x54, 0x65, 0xba, 0x70, 0xbe, 0x42, 0x67, 0xdb, 0xc0, 0xab, 0x70, 0x61, 0x86, 0x17,
	0xe7, 0x6b, 0xe6, 0x98, 0x47, 0x50, 0x0f, 0x76, 0x12, 0xa5, 0xe8, 0x56, 0x7a, 0xd5, 0x7e, 0xf3,
	0x1c, 0xdc, 0xbc, 0xae, 0x9f, 0x12, 0x9b, 0xe6, 0xe6, 0x2e, 0xff, 0x8a, 0xf9, 0x5b, 0xb0, 0xae,
	0x04, 0xd1, 0x8d, 0x79, 0x06, 0xb6, 0x9a, 0x87, 0x79, 0x8c, 0xc3, 0x6c, 0x28, 0x0a, 0x00, 0x21,
	0xa8, 0xe9, 0x69, 0x49, 0xbb, 0xa0, 0xd7, 0xce, 0x31, 0x58, 0x63, 0x2a, 0x87, 0x82, 0x4f, 0x14,
	0x4d, 0xb0, 0xc4, 0x46, 0xa7, 0xd7, 0xce, 0x8f, 0x32, 0xb4, 0xd6, 0x7e, 0x76, 0x8f, 0x89, 0x7c,
	0x04, 0xd5, 0x58, 0x10, 0x33, 0xfe, 0x0d, 0xd7, 0x84, 0xf3, 0x15, 0xa8, 0x38, 0x4e, 0xb3, 0xc9,
	0x6f, 0xb8, 0xa6, 0xb6, 0xaf, 0xc0, 0xe2, 0xa8, 0xb5, 0x3f, 0x1d, 0xd5, 0xc4, 0x29, 0x7e, 0xba,
	0xff, 0x36, 0x4e, 0x04, 0xf5, 0xf4, 0x5f, 0x1e, 0x0d, 0xa0, 0xa3, 0x1f, 0x12, 0x35, 0x8d, 0x01,
	0x8b, 0xd5, 0xdf, 0x57, 0x42, 0xe7, 0x73, 0x13, 0xe7, 0x40, 0x71, 0x6a, 0x28, 0xbd, 0xf8, 0x22,
	0x25, 0xd0, 0x19, 0x1c, 0x16, 0x82, 0x19, 0x0e, 0x73, 0x45, 0x7a, 0x39, 0x28, 0x53, 0x7c, 0xc2,
	0xa1, 0x91, 0x9c, 0xbf, 0x03, 0x3b, 0x7f, 0xdd, 0xd0, 0x00, 0xec, 0xfc, 0x35, 0x43, 0x07, 0xee,
	0xc3, 0x97, 0xed, 0xa8, 0xe5, 0xae, 0x3e, 0x5d, 0xa7, 0xe5, 0x9b, 0xba, 0x7e, 0x05, 0x5f, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xd2, 0x3c, 0x95, 0x2f, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NfnNotifyClient is the client API for NfnNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NfnNotifyClient interface {
	Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error)
}

type nfnNotifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNfnNotifyClient(cc grpc.ClientConnInterface) NfnNotifyClient {
	return &nfnNotifyClient{cc}
}

func (c *nfnNotifyClient) Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NfnNotify_serviceDesc.Streams[0], "/nfnNotify/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &nfnNotifySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NfnNotify_SubscribeClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type nfnNotifySubscribeClient struct {
	grpc.ClientStream
}

func (x *nfnNotifySubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NfnNotifyServer is the server API for NfnNotify service.
type NfnNotifyServer interface {
	Subscribe(*SubscribeContext, NfnNotify_SubscribeServer) error
}

// UnimplementedNfnNotifyServer can be embedded to have forward compatible implementations.
type UnimplementedNfnNotifyServer struct {
}

func (*UnimplementedNfnNotifyServer) Subscribe(req *SubscribeContext, srv NfnNotify_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterNfnNotifyServer(s *grpc.Server, srv NfnNotifyServer) {
	s.RegisterService(&_NfnNotify_serviceDesc, srv)
}

func _NfnNotify_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeContext)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NfnNotifyServer).Subscribe(m, &nfnNotifySubscribeServer{stream})
}

type NfnNotify_SubscribeServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type nfnNotifySubscribeServer struct {
	grpc.ServerStream
}

func (x *nfnNotifySubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

var _NfnNotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nfnNotify",
	HandlerType: (*NfnNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NfnNotify_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/nfnNotify/proto/nfn.proto",
}