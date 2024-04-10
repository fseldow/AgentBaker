// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: pkg/proto/nbcontract/v1/clusterconfig.proto

package nbcontractv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterConfig_VM int32

const (
	ClusterConfig_UNSPECIFIED ClusterConfig_VM = 0
	ClusterConfig_STANDARD    ClusterConfig_VM = 1
	ClusterConfig_VMSS        ClusterConfig_VM = 2
)

// Enum value maps for ClusterConfig_VM.
var (
	ClusterConfig_VM_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "STANDARD",
		2: "VMSS",
	}
	ClusterConfig_VM_value = map[string]int32{
		"UNSPECIFIED": 0,
		"STANDARD":    1,
		"VMSS":        2,
	}
)

func (x ClusterConfig_VM) Enum() *ClusterConfig_VM {
	p := new(ClusterConfig_VM)
	*p = x
	return p
}

func (x ClusterConfig_VM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterConfig_VM) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes[0].Descriptor()
}

func (ClusterConfig_VM) Type() protoreflect.EnumType {
	return &file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes[0]
}

func (x ClusterConfig_VM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterConfig_VM.Descriptor instead.
func (ClusterConfig_VM) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP(), []int{0, 0}
}

type LoadBalancerConfig_LoadBalancerSku int32

const (
	LoadBalancerConfig_UNSPECIFIED LoadBalancerConfig_LoadBalancerSku = 0
	LoadBalancerConfig_BASIC       LoadBalancerConfig_LoadBalancerSku = 1
	LoadBalancerConfig_STANDARD    LoadBalancerConfig_LoadBalancerSku = 2
)

// Enum value maps for LoadBalancerConfig_LoadBalancerSku.
var (
	LoadBalancerConfig_LoadBalancerSku_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "BASIC",
		2: "STANDARD",
	}
	LoadBalancerConfig_LoadBalancerSku_value = map[string]int32{
		"UNSPECIFIED": 0,
		"BASIC":       1,
		"STANDARD":    2,
	}
)

func (x LoadBalancerConfig_LoadBalancerSku) Enum() *LoadBalancerConfig_LoadBalancerSku {
	p := new(LoadBalancerConfig_LoadBalancerSku)
	*p = x
	return p
}

func (x LoadBalancerConfig_LoadBalancerSku) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadBalancerConfig_LoadBalancerSku) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes[1].Descriptor()
}

func (LoadBalancerConfig_LoadBalancerSku) Type() protoreflect.EnumType {
	return &file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes[1]
}

func (x LoadBalancerConfig_LoadBalancerSku) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadBalancerConfig_LoadBalancerSku.Descriptor instead.
func (LoadBalancerConfig_LoadBalancerSku) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP(), []int{2, 0}
}

// Cluster Config fields stored in azure.json used by cloud-provider-azure
type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rescource group name
	ResourceGroup string `protobuf:"bytes,1,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	// Location
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// VM type
	VmType ClusterConfig_VM `protobuf:"varint,3,opt,name=vm_type,json=vmType,proto3,enum=nbcontract.v1.ClusterConfig_VM" json:"vm_type,omitempty"` // default to standard for v1.27 and below versions and vmss for v1.28+ versions
	// Primary availability set name
	PrimaryAvailabilitySet string `protobuf:"bytes,4,opt,name=primary_availability_set,json=primaryAvailabilitySet,proto3" json:"primary_availability_set,omitempty"`
	// Primary scale set name
	PrimaryScaleSet string `protobuf:"bytes,5,opt,name=primary_scale_set,json=primaryScaleSet,proto3" json:"primary_scale_set,omitempty"`
	// Cluster network config
	VirtualNetworkConfig *ClusterNetworkConfig `protobuf:"bytes,6,opt,name=virtual_network_config,json=virtualNetworkConfig,proto3" json:"virtual_network_config,omitempty"`
	// Specifiy if it uses instance metadata
	UseInstanceMetadata bool `protobuf:"varint,7,opt,name=use_instance_metadata,json=useInstanceMetadata,proto3" json:"use_instance_metadata,omitempty"` // default to false
	// Load balancer config
	LoadBalancerConfig *LoadBalancerConfig `protobuf:"bytes,8,opt,name=load_balancer_config,json=loadBalancerConfig,proto3" json:"load_balancer_config,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterConfig) GetResourceGroup() string {
	if x != nil {
		return x.ResourceGroup
	}
	return ""
}

func (x *ClusterConfig) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ClusterConfig) GetVmType() ClusterConfig_VM {
	if x != nil {
		return x.VmType
	}
	return ClusterConfig_UNSPECIFIED
}

func (x *ClusterConfig) GetPrimaryAvailabilitySet() string {
	if x != nil {
		return x.PrimaryAvailabilitySet
	}
	return ""
}

func (x *ClusterConfig) GetPrimaryScaleSet() string {
	if x != nil {
		return x.PrimaryScaleSet
	}
	return ""
}

func (x *ClusterConfig) GetVirtualNetworkConfig() *ClusterNetworkConfig {
	if x != nil {
		return x.VirtualNetworkConfig
	}
	return nil
}

func (x *ClusterConfig) GetUseInstanceMetadata() bool {
	if x != nil {
		return x.UseInstanceMetadata
	}
	return false
}

func (x *ClusterConfig) GetLoadBalancerConfig() *LoadBalancerConfig {
	if x != nil {
		return x.LoadBalancerConfig
	}
	return nil
}

type ClusterNetworkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Virtual network name
	VnetName string `protobuf:"bytes,1,opt,name=vnet_name,json=vnetName,proto3" json:"vnet_name,omitempty"`
	// Virtual network resource group
	VnetResourceGroup string `protobuf:"bytes,2,opt,name=vnet_resource_group,json=vnetResourceGroup,proto3" json:"vnet_resource_group,omitempty"`
	// Subnet name
	Subnet string `protobuf:"bytes,3,opt,name=subnet,proto3" json:"subnet,omitempty"`
	// Network security group name
	SecurityGroupName string `protobuf:"bytes,4,opt,name=security_group_name,json=securityGroupName,proto3" json:"security_group_name,omitempty"`
	// Route table name
	RouteTable string `protobuf:"bytes,5,opt,name=route_table,json=routeTable,proto3" json:"route_table,omitempty"`
}

func (x *ClusterNetworkConfig) Reset() {
	*x = ClusterNetworkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterNetworkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterNetworkConfig) ProtoMessage() {}

func (x *ClusterNetworkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterNetworkConfig.ProtoReflect.Descriptor instead.
func (*ClusterNetworkConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterNetworkConfig) GetVnetName() string {
	if x != nil {
		return x.VnetName
	}
	return ""
}

func (x *ClusterNetworkConfig) GetVnetResourceGroup() string {
	if x != nil {
		return x.VnetResourceGroup
	}
	return ""
}

func (x *ClusterNetworkConfig) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *ClusterNetworkConfig) GetSecurityGroupName() string {
	if x != nil {
		return x.SecurityGroupName
	}
	return ""
}

func (x *ClusterNetworkConfig) GetRouteTable() string {
	if x != nil {
		return x.RouteTable
	}
	return ""
}

type LoadBalancerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Load balancer sku, default to basic
	LoadBalancerSku LoadBalancerConfig_LoadBalancerSku `protobuf:"varint,1,opt,name=load_balancer_sku,json=loadBalancerSku,proto3,enum=nbcontract.v1.LoadBalancerConfig_LoadBalancerSku" json:"load_balancer_sku,omitempty"`
	// Specify if master node should be excluded from standard load balancer, default to true
	ExcludeMasterFromStandardLoadBalancer *bool `protobuf:"varint,2,opt,name=exclude_master_from_standard_load_balancer,json=excludeMasterFromStandardLoadBalancer,proto3,oneof" json:"exclude_master_from_standard_load_balancer,omitempty"`
	// Maximum number of load balancer rules, default to 148
	MaxLoadBalancerRuleCount *int32 `protobuf:"varint,3,opt,name=max_load_balancer_rule_count,json=maxLoadBalancerRuleCount,proto3,oneof" json:"max_load_balancer_rule_count,omitempty"`
	// Disable outbound SNAT (Source Network Address Translation) for load balancer, default to false
	DisableOutboundSnat *bool `protobuf:"varint,4,opt,name=disable_outbound_snat,json=disableOutboundSnat,proto3,oneof" json:"disable_outbound_snat,omitempty"`
}

func (x *LoadBalancerConfig) Reset() {
	*x = LoadBalancerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancerConfig) ProtoMessage() {}

func (x *LoadBalancerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancerConfig.ProtoReflect.Descriptor instead.
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP(), []int{2}
}

func (x *LoadBalancerConfig) GetLoadBalancerSku() LoadBalancerConfig_LoadBalancerSku {
	if x != nil {
		return x.LoadBalancerSku
	}
	return LoadBalancerConfig_UNSPECIFIED
}

func (x *LoadBalancerConfig) GetExcludeMasterFromStandardLoadBalancer() bool {
	if x != nil && x.ExcludeMasterFromStandardLoadBalancer != nil {
		return *x.ExcludeMasterFromStandardLoadBalancer
	}
	return false
}

func (x *LoadBalancerConfig) GetMaxLoadBalancerRuleCount() int32 {
	if x != nil && x.MaxLoadBalancerRuleCount != nil {
		return *x.MaxLoadBalancerRuleCount
	}
	return 0
}

func (x *LoadBalancerConfig) GetDisableOutboundSnat() bool {
	if x != nil && x.DisableOutboundSnat != nil {
		return *x.DisableOutboundSnat
	}
	return false
}

var File_pkg_proto_nbcontract_v1_clusterconfig_proto protoreflect.FileDescriptor

var file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x85, 0x04, 0x0a,
	0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x07, 0x76, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x56, 0x4d, 0x52, 0x06, 0x76, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x12, 0x59, 0x0a, 0x16, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x15,
	0x75, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x75, 0x73, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x53, 0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x12, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2d, 0x0a, 0x02, 0x56, 0x4d, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4d,
	0x53, 0x53, 0x10, 0x02, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x76, 0x6e, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x76, 0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x76, 0x6e,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0xf8, 0x03, 0x0a, 0x12, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5d, 0x0a, 0x11, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x6b, 0x75, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x6b, 0x75, 0x52, 0x0f, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x6b, 0x75, 0x12, 0x5e, 0x0a, 0x2a, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x25, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x1c, 0x6d, 0x61, 0x78,
	0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x01, 0x52, 0x18, 0x6d, 0x61, 0x78, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x37,
	0x0a, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x73, 0x6e, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52,
	0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x6e, 0x61, 0x74, 0x88, 0x01, 0x01, 0x22, 0x3b, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x6b, 0x75, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41,
	0x52, 0x44, 0x10, 0x02, 0x42, 0x2d, 0x0a, 0x2b, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x73, 0x6e, 0x61, 0x74, 0x42, 0xbe,
	0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x6e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e,
	0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0d, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x19, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x4e, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescOnce sync.Once
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescData = file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDesc
)

func file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescGZIP() []byte {
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescData)
	})
	return file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDescData
}

var file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_nbcontract_v1_clusterconfig_proto_goTypes = []interface{}{
	(ClusterConfig_VM)(0),                   // 0: nbcontract.v1.ClusterConfig.VM
	(LoadBalancerConfig_LoadBalancerSku)(0), // 1: nbcontract.v1.LoadBalancerConfig.LoadBalancerSku
	(*ClusterConfig)(nil),                   // 2: nbcontract.v1.ClusterConfig
	(*ClusterNetworkConfig)(nil),            // 3: nbcontract.v1.ClusterNetworkConfig
	(*LoadBalancerConfig)(nil),              // 4: nbcontract.v1.LoadBalancerConfig
}
var file_pkg_proto_nbcontract_v1_clusterconfig_proto_depIdxs = []int32{
	0, // 0: nbcontract.v1.ClusterConfig.vm_type:type_name -> nbcontract.v1.ClusterConfig.VM
	3, // 1: nbcontract.v1.ClusterConfig.virtual_network_config:type_name -> nbcontract.v1.ClusterNetworkConfig
	4, // 2: nbcontract.v1.ClusterConfig.load_balancer_config:type_name -> nbcontract.v1.LoadBalancerConfig
	1, // 3: nbcontract.v1.LoadBalancerConfig.load_balancer_sku:type_name -> nbcontract.v1.LoadBalancerConfig.LoadBalancerSku
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_proto_nbcontract_v1_clusterconfig_proto_init() }
func file_pkg_proto_nbcontract_v1_clusterconfig_proto_init() {
	if File_pkg_proto_nbcontract_v1_clusterconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterNetworkConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadBalancerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_nbcontract_v1_clusterconfig_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nbcontract_v1_clusterconfig_proto_depIdxs,
		EnumInfos:         file_pkg_proto_nbcontract_v1_clusterconfig_proto_enumTypes,
		MessageInfos:      file_pkg_proto_nbcontract_v1_clusterconfig_proto_msgTypes,
	}.Build()
	File_pkg_proto_nbcontract_v1_clusterconfig_proto = out.File
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_rawDesc = nil
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_goTypes = nil
	file_pkg_proto_nbcontract_v1_clusterconfig_proto_depIdxs = nil
}
