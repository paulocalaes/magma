//
//Copyright 2022 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: dp/protos/enodebd_dp.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CBSDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FccId         string `protobuf:"bytes,2,opt,name=fcc_id,json=fccId,proto3" json:"fcc_id,omitempty"`
	SerialNumber  string `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	MinPower      int64  `protobuf:"varint,4,opt,name=min_power,json=minPower,proto3" json:"min_power,omitempty"`
	MaxPower      int64  `protobuf:"varint,5,opt,name=max_power,json=maxPower,proto3" json:"max_power,omitempty"`
	AntennaGain   int64  `protobuf:"varint,6,opt,name=antenna_gain,json=antennaGain,proto3" json:"antenna_gain,omitempty"`
	NumberOfPorts int64  `protobuf:"varint,7,opt,name=number_of_ports,json=numberOfPorts,proto3" json:"number_of_ports,omitempty"`
}

func (x *CBSDRequest) Reset() {
	*x = CBSDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_enodebd_dp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBSDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBSDRequest) ProtoMessage() {}

func (x *CBSDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_enodebd_dp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBSDRequest.ProtoReflect.Descriptor instead.
func (*CBSDRequest) Descriptor() ([]byte, []int) {
	return file_dp_protos_enodebd_dp_proto_rawDescGZIP(), []int{0}
}

func (x *CBSDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CBSDRequest) GetFccId() string {
	if x != nil {
		return x.FccId
	}
	return ""
}

func (x *CBSDRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *CBSDRequest) GetMinPower() int64 {
	if x != nil {
		return x.MinPower
	}
	return 0
}

func (x *CBSDRequest) GetMaxPower() int64 {
	if x != nil {
		return x.MaxPower
	}
	return 0
}

func (x *CBSDRequest) GetAntennaGain() int64 {
	if x != nil {
		return x.AntennaGain
	}
	return 0
}

func (x *CBSDRequest) GetNumberOfPorts() int64 {
	if x != nil {
		return x.NumberOfPorts
	}
	return 0
}

type CBSDStateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels                  []*LteChannel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	RadioEnabled              bool          `protobuf:"varint,2,opt,name=radio_enabled,json=radioEnabled,proto3" json:"radio_enabled,omitempty"`
	CarrierAggregationEnabled bool          `protobuf:"varint,3,opt,name=carrier_aggregation_enabled,json=carrierAggregationEnabled,proto3" json:"carrier_aggregation_enabled,omitempty"`
	// TODO for backwards compatibility only. Remove once enodebd part is updated
	Channel *LteChannel `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *CBSDStateResult) Reset() {
	*x = CBSDStateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_enodebd_dp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBSDStateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBSDStateResult) ProtoMessage() {}

func (x *CBSDStateResult) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_enodebd_dp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBSDStateResult.ProtoReflect.Descriptor instead.
func (*CBSDStateResult) Descriptor() ([]byte, []int) {
	return file_dp_protos_enodebd_dp_proto_rawDescGZIP(), []int{1}
}

func (x *CBSDStateResult) GetChannels() []*LteChannel {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *CBSDStateResult) GetRadioEnabled() bool {
	if x != nil {
		return x.RadioEnabled
	}
	return false
}

func (x *CBSDStateResult) GetCarrierAggregationEnabled() bool {
	if x != nil {
		return x.CarrierAggregationEnabled
	}
	return false
}

func (x *CBSDStateResult) GetChannel() *LteChannel {
	if x != nil {
		return x.Channel
	}
	return nil
}

type LteChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LowFrequencyHz  int64   `protobuf:"varint,1,opt,name=low_frequency_hz,json=lowFrequencyHz,proto3" json:"low_frequency_hz,omitempty"`
	HighFrequencyHz int64   `protobuf:"varint,2,opt,name=high_frequency_hz,json=highFrequencyHz,proto3" json:"high_frequency_hz,omitempty"`
	MaxEirpDbmMhz   float32 `protobuf:"fixed32,3,opt,name=max_eirp_dbm_mhz,json=maxEirpDbmMhz,proto3" json:"max_eirp_dbm_mhz,omitempty"`
}

func (x *LteChannel) Reset() {
	*x = LteChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dp_protos_enodebd_dp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LteChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LteChannel) ProtoMessage() {}

func (x *LteChannel) ProtoReflect() protoreflect.Message {
	mi := &file_dp_protos_enodebd_dp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LteChannel.ProtoReflect.Descriptor instead.
func (*LteChannel) Descriptor() ([]byte, []int) {
	return file_dp_protos_enodebd_dp_proto_rawDescGZIP(), []int{2}
}

func (x *LteChannel) GetLowFrequencyHz() int64 {
	if x != nil {
		return x.LowFrequencyHz
	}
	return 0
}

func (x *LteChannel) GetHighFrequencyHz() int64 {
	if x != nil {
		return x.HighFrequencyHz
	}
	return 0
}

func (x *LteChannel) GetMaxEirpDbmMhz() float32 {
	if x != nil {
		return x.MaxEirpDbmMhz
	}
	return 0
}

var File_dp_protos_enodebd_dp_proto protoreflect.FileDescriptor

var file_dp_protos_enodebd_dp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x6f, 0x64,
	0x65, 0x62, 0x64, 0x5f, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a,
	0x0b, 0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x63, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x63, 0x63, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x5f, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x61, 0x6e, 0x74, 0x65, 0x6e, 0x6e, 0x61, 0x47, 0x61, 0x69, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f,
	0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0xc6, 0x01, 0x0a, 0x0f, 0x43, 0x42, 0x53, 0x44, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x61, 0x64, 0x69,
	0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x1b, 0x63, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x63,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0x8b, 0x01, 0x0a, 0x0a, 0x4c, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x68, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6c, 0x6f, 0x77, 0x46, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x7a, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x69, 0x67, 0x68,
	0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x68, 0x7a, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x68, 0x69, 0x67, 0x68, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x48, 0x7a, 0x12, 0x27, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x69, 0x72, 0x70,
	0x5f, 0x64, 0x62, 0x6d, 0x5f, 0x6d, 0x68, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x6d, 0x61, 0x78, 0x45, 0x69, 0x72, 0x70, 0x44, 0x62, 0x6d, 0x4d, 0x68, 0x7a, 0x32, 0xd7, 0x01,
	0x0a, 0x09, 0x44, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x42, 0x53, 0x44, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x43, 0x42,
	0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x42, 0x53, 0x44,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a,
	0x0c, 0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0c, 0x2e,
	0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x42,
	0x53, 0x44, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0e, 0x43, 0x42, 0x53, 0x44, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0c, 0x2e, 0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x43, 0x42, 0x53, 0x44, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x6c, 0x69, 0x6e,
	0x71, 0x75, 0x69, 0x73, 0x68, 0x12, 0x0c, 0x2e, 0x43, 0x42, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x42, 0x53, 0x44, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2f, 0x64, 0x70, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dp_protos_enodebd_dp_proto_rawDescOnce sync.Once
	file_dp_protos_enodebd_dp_proto_rawDescData = file_dp_protos_enodebd_dp_proto_rawDesc
)

func file_dp_protos_enodebd_dp_proto_rawDescGZIP() []byte {
	file_dp_protos_enodebd_dp_proto_rawDescOnce.Do(func() {
		file_dp_protos_enodebd_dp_proto_rawDescData = protoimpl.X.CompressGZIP(file_dp_protos_enodebd_dp_proto_rawDescData)
	})
	return file_dp_protos_enodebd_dp_proto_rawDescData
}

var file_dp_protos_enodebd_dp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dp_protos_enodebd_dp_proto_goTypes = []interface{}{
	(*CBSDRequest)(nil),     // 0: CBSDRequest
	(*CBSDStateResult)(nil), // 1: CBSDStateResult
	(*LteChannel)(nil),      // 2: LteChannel
}
var file_dp_protos_enodebd_dp_proto_depIdxs = []int32{
	2, // 0: CBSDStateResult.channels:type_name -> LteChannel
	2, // 1: CBSDStateResult.channel:type_name -> LteChannel
	0, // 2: DPService.GetCBSDState:input_type -> CBSDRequest
	0, // 3: DPService.CBSDRegister:input_type -> CBSDRequest
	0, // 4: DPService.CBSDDeregister:input_type -> CBSDRequest
	0, // 5: DPService.CBSDRelinquish:input_type -> CBSDRequest
	1, // 6: DPService.GetCBSDState:output_type -> CBSDStateResult
	1, // 7: DPService.CBSDRegister:output_type -> CBSDStateResult
	1, // 8: DPService.CBSDDeregister:output_type -> CBSDStateResult
	1, // 9: DPService.CBSDRelinquish:output_type -> CBSDStateResult
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dp_protos_enodebd_dp_proto_init() }
func file_dp_protos_enodebd_dp_proto_init() {
	if File_dp_protos_enodebd_dp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dp_protos_enodebd_dp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBSDRequest); i {
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
		file_dp_protos_enodebd_dp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBSDStateResult); i {
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
		file_dp_protos_enodebd_dp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LteChannel); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dp_protos_enodebd_dp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dp_protos_enodebd_dp_proto_goTypes,
		DependencyIndexes: file_dp_protos_enodebd_dp_proto_depIdxs,
		MessageInfos:      file_dp_protos_enodebd_dp_proto_msgTypes,
	}.Build()
	File_dp_protos_enodebd_dp_proto = out.File
	file_dp_protos_enodebd_dp_proto_rawDesc = nil
	file_dp_protos_enodebd_dp_proto_goTypes = nil
	file_dp_protos_enodebd_dp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DPServiceClient is the client API for DPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DPServiceClient interface {
	GetCBSDState(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDRegister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDDeregister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
	CBSDRelinquish(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error)
}

type dPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDPServiceClient(cc grpc.ClientConnInterface) DPServiceClient {
	return &dPServiceClient{cc}
}

func (c *dPServiceClient) GetCBSDState(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/GetCBSDState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDRegister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDDeregister(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDDeregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dPServiceClient) CBSDRelinquish(ctx context.Context, in *CBSDRequest, opts ...grpc.CallOption) (*CBSDStateResult, error) {
	out := new(CBSDStateResult)
	err := c.cc.Invoke(ctx, "/DPService/CBSDRelinquish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DPServiceServer is the server API for DPService service.
type DPServiceServer interface {
	GetCBSDState(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDRegister(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDDeregister(context.Context, *CBSDRequest) (*CBSDStateResult, error)
	CBSDRelinquish(context.Context, *CBSDRequest) (*CBSDStateResult, error)
}

// UnimplementedDPServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDPServiceServer struct {
}

func (*UnimplementedDPServiceServer) GetCBSDState(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCBSDState not implemented")
}
func (*UnimplementedDPServiceServer) CBSDRegister(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDRegister not implemented")
}
func (*UnimplementedDPServiceServer) CBSDDeregister(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDDeregister not implemented")
}
func (*UnimplementedDPServiceServer) CBSDRelinquish(context.Context, *CBSDRequest) (*CBSDStateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CBSDRelinquish not implemented")
}

func RegisterDPServiceServer(s *grpc.Server, srv DPServiceServer) {
	s.RegisterService(&_DPService_serviceDesc, srv)
}

func _DPService_GetCBSDState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).GetCBSDState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/GetCBSDState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).GetCBSDState(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDRegister(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDDeregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDDeregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDDeregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDDeregister(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DPService_CBSDRelinquish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CBSDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DPServiceServer).CBSDRelinquish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DPService/CBSDRelinquish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DPServiceServer).CBSDRelinquish(ctx, req.(*CBSDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DPService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DPService",
	HandlerType: (*DPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCBSDState",
			Handler:    _DPService_GetCBSDState_Handler,
		},
		{
			MethodName: "CBSDRegister",
			Handler:    _DPService_CBSDRegister_Handler,
		},
		{
			MethodName: "CBSDDeregister",
			Handler:    _DPService_CBSDDeregister_Handler,
		},
		{
			MethodName: "CBSDRelinquish",
			Handler:    _DPService_CBSDRelinquish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dp/protos/enodebd_dp.proto",
}
