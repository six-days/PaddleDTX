// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Algorithm is a list of algorithms offered
type Algorithm int32

const (
	Algorithm_LINEAR_REGRESSION_VL Algorithm = 0
	Algorithm_LOGIC_REGRESSION_VL  Algorithm = 1
)

var Algorithm_name = map[int32]string{
	0: "LINEAR_REGRESSION_VL",
	1: "LOGIC_REGRESSION_VL",
}

var Algorithm_value = map[string]int32{
	"LINEAR_REGRESSION_VL": 0,
	"LOGIC_REGRESSION_VL":  1,
}

func (x Algorithm) String() string {
	return proto.EnumName(Algorithm_name, int32(x))
}

func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

// TaskType defines types of task
type TaskType int32

const (
	TaskType_LEARN   TaskType = 0
	TaskType_PREDICT TaskType = 1
)

var TaskType_name = map[int32]string{
	0: "LEARN",
	1: "PREDICT",
}

var TaskType_value = map[string]int32{
	"LEARN":   0,
	"PREDICT": 1,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}

func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

// RegMode regulation mode for training
type RegMode int32

const (
	RegMode_Reg_None  RegMode = 0
	RegMode_Reg_Lasso RegMode = 1
	RegMode_Reg_Ridge RegMode = 2
)

var RegMode_name = map[int32]string{
	0: "Reg_None",
	1: "Reg_Lasso",
	2: "Reg_Ridge",
}

var RegMode_value = map[string]int32{
	"Reg_None":  0,
	"Reg_Lasso": 1,
	"Reg_Ridge": 2,
}

func (x RegMode) String() string {
	return proto.EnumName(RegMode_name, int32(x))
}

func (RegMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

// TrainParams lists all the parameters for training
type TrainParams struct {
	Label                string   `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	LabelName            string   `protobuf:"bytes,2,opt,name=labelName,proto3" json:"labelName,omitempty"`
	RegMode              RegMode  `protobuf:"varint,3,opt,name=regMode,proto3,enum=common.RegMode" json:"regMode,omitempty"`
	RegParam             float64  `protobuf:"fixed64,4,opt,name=regParam,proto3" json:"regParam,omitempty"`
	Alpha                float64  `protobuf:"fixed64,5,opt,name=alpha,proto3" json:"alpha,omitempty"`
	Amplitude            float64  `protobuf:"fixed64,6,opt,name=amplitude,proto3" json:"amplitude,omitempty"`
	Accuracy             int64    `protobuf:"varint,7,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	IsTagPart            bool     `protobuf:"varint,8,opt,name=isTagPart,proto3" json:"isTagPart,omitempty"`
	IdName               string   `protobuf:"bytes,9,opt,name=idName,proto3" json:"idName,omitempty"`
	BatchSize            int64    `protobuf:"varint,10,opt,name=BatchSize,proto3" json:"BatchSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainParams) Reset()         { *m = TrainParams{} }
func (m *TrainParams) String() string { return proto.CompactTextString(m) }
func (*TrainParams) ProtoMessage()    {}
func (*TrainParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *TrainParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainParams.Unmarshal(m, b)
}
func (m *TrainParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainParams.Marshal(b, m, deterministic)
}
func (m *TrainParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainParams.Merge(m, src)
}
func (m *TrainParams) XXX_Size() int {
	return xxx_messageInfo_TrainParams.Size(m)
}
func (m *TrainParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainParams.DiscardUnknown(m)
}

var xxx_messageInfo_TrainParams proto.InternalMessageInfo

func (m *TrainParams) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *TrainParams) GetLabelName() string {
	if m != nil {
		return m.LabelName
	}
	return ""
}

func (m *TrainParams) GetRegMode() RegMode {
	if m != nil {
		return m.RegMode
	}
	return RegMode_Reg_None
}

func (m *TrainParams) GetRegParam() float64 {
	if m != nil {
		return m.RegParam
	}
	return 0
}

func (m *TrainParams) GetAlpha() float64 {
	if m != nil {
		return m.Alpha
	}
	return 0
}

func (m *TrainParams) GetAmplitude() float64 {
	if m != nil {
		return m.Amplitude
	}
	return 0
}

func (m *TrainParams) GetAccuracy() int64 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *TrainParams) GetIsTagPart() bool {
	if m != nil {
		return m.IsTagPart
	}
	return false
}

func (m *TrainParams) GetIdName() string {
	if m != nil {
		return m.IdName
	}
	return ""
}

func (m *TrainParams) GetBatchSize() int64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

// TrainModels is final result of distributed training
type TrainModels struct {
	Thetas               map[string]float64 `protobuf:"bytes,1,rep,name=thetas,proto3" json:"thetas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Xbars                map[string]float64 `protobuf:"bytes,2,rep,name=xbars,proto3" json:"xbars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Sigmas               map[string]float64 `protobuf:"bytes,3,rep,name=sigmas,proto3" json:"sigmas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Label                string             `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	IsTagPart            bool               `protobuf:"varint,5,opt,name=isTagPart,proto3" json:"isTagPart,omitempty"`
	IdName               string             `protobuf:"bytes,6,opt,name=idName,proto3" json:"idName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TrainModels) Reset()         { *m = TrainModels{} }
func (m *TrainModels) String() string { return proto.CompactTextString(m) }
func (*TrainModels) ProtoMessage()    {}
func (*TrainModels) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *TrainModels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainModels.Unmarshal(m, b)
}
func (m *TrainModels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainModels.Marshal(b, m, deterministic)
}
func (m *TrainModels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainModels.Merge(m, src)
}
func (m *TrainModels) XXX_Size() int {
	return xxx_messageInfo_TrainModels.Size(m)
}
func (m *TrainModels) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainModels.DiscardUnknown(m)
}

var xxx_messageInfo_TrainModels proto.InternalMessageInfo

func (m *TrainModels) GetThetas() map[string]float64 {
	if m != nil {
		return m.Thetas
	}
	return nil
}

func (m *TrainModels) GetXbars() map[string]float64 {
	if m != nil {
		return m.Xbars
	}
	return nil
}

func (m *TrainModels) GetSigmas() map[string]float64 {
	if m != nil {
		return m.Sigmas
	}
	return nil
}

func (m *TrainModels) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *TrainModels) GetIsTagPart() bool {
	if m != nil {
		return m.IsTagPart
	}
	return false
}

func (m *TrainModels) GetIdName() string {
	if m != nil {
		return m.IdName
	}
	return ""
}

// TaskParams lists all the parameters in a task
type TaskParams struct {
	Algo                 Algorithm    `protobuf:"varint,1,opt,name=algo,proto3,enum=common.Algorithm" json:"algo,omitempty"`
	TaskType             TaskType     `protobuf:"varint,2,opt,name=taskType,proto3,enum=common.TaskType" json:"taskType,omitempty"`
	TrainParams          *TrainParams `protobuf:"bytes,3,opt,name=trainParams,proto3" json:"trainParams,omitempty"`
	ModelTaskID          string       `protobuf:"bytes,4,opt,name=modelTaskID,proto3" json:"modelTaskID,omitempty"`
	ModelParams          *TrainModels `protobuf:"bytes,5,opt,name=modelParams,proto3" json:"modelParams,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TaskParams) Reset()         { *m = TaskParams{} }
func (m *TaskParams) String() string { return proto.CompactTextString(m) }
func (*TaskParams) ProtoMessage()    {}
func (*TaskParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *TaskParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskParams.Unmarshal(m, b)
}
func (m *TaskParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskParams.Marshal(b, m, deterministic)
}
func (m *TaskParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskParams.Merge(m, src)
}
func (m *TaskParams) XXX_Size() int {
	return xxx_messageInfo_TaskParams.Size(m)
}
func (m *TaskParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskParams.DiscardUnknown(m)
}

var xxx_messageInfo_TaskParams proto.InternalMessageInfo

func (m *TaskParams) GetAlgo() Algorithm {
	if m != nil {
		return m.Algo
	}
	return Algorithm_LINEAR_REGRESSION_VL
}

func (m *TaskParams) GetTaskType() TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskType_LEARN
}

func (m *TaskParams) GetTrainParams() *TrainParams {
	if m != nil {
		return m.TrainParams
	}
	return nil
}

func (m *TaskParams) GetModelTaskID() string {
	if m != nil {
		return m.ModelTaskID
	}
	return ""
}

func (m *TaskParams) GetModelParams() *TrainModels {
	if m != nil {
		return m.ModelParams
	}
	return nil
}

// TrainTaskResult defines final result of training
type TrainTaskResult struct {
	TaskID               string   `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Model                []byte   `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	ErrMsg               string   `protobuf:"bytes,4,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainTaskResult) Reset()         { *m = TrainTaskResult{} }
func (m *TrainTaskResult) String() string { return proto.CompactTextString(m) }
func (*TrainTaskResult) ProtoMessage()    {}
func (*TrainTaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *TrainTaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainTaskResult.Unmarshal(m, b)
}
func (m *TrainTaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainTaskResult.Marshal(b, m, deterministic)
}
func (m *TrainTaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainTaskResult.Merge(m, src)
}
func (m *TrainTaskResult) XXX_Size() int {
	return xxx_messageInfo_TrainTaskResult.Size(m)
}
func (m *TrainTaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainTaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_TrainTaskResult proto.InternalMessageInfo

func (m *TrainTaskResult) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *TrainTaskResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *TrainTaskResult) GetModel() []byte {
	if m != nil {
		return m.Model
	}
	return nil
}

func (m *TrainTaskResult) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// PredictTaskResult defines final result of prediction
type PredictTaskResult struct {
	TaskID               string   `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Outcomes             []byte   `protobuf:"bytes,3,opt,name=outcomes,proto3" json:"outcomes,omitempty"`
	ErrMsg               string   `protobuf:"bytes,4,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PredictTaskResult) Reset()         { *m = PredictTaskResult{} }
func (m *PredictTaskResult) String() string { return proto.CompactTextString(m) }
func (*PredictTaskResult) ProtoMessage()    {}
func (*PredictTaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *PredictTaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictTaskResult.Unmarshal(m, b)
}
func (m *PredictTaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictTaskResult.Marshal(b, m, deterministic)
}
func (m *PredictTaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictTaskResult.Merge(m, src)
}
func (m *PredictTaskResult) XXX_Size() int {
	return xxx_messageInfo_PredictTaskResult.Size(m)
}
func (m *PredictTaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictTaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_PredictTaskResult proto.InternalMessageInfo

func (m *PredictTaskResult) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *PredictTaskResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PredictTaskResult) GetOutcomes() []byte {
	if m != nil {
		return m.Outcomes
	}
	return nil
}

func (m *PredictTaskResult) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// StartTaskRequest is message sent to a cluster member to start a training task or predicting task.
type StartTaskRequest struct {
	TaskID               string      `protobuf:"bytes,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	File                 []byte      `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Hosts                []string    `protobuf:"bytes,4,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Params               *TaskParams `protobuf:"bytes,5,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StartTaskRequest) Reset()         { *m = StartTaskRequest{} }
func (m *StartTaskRequest) String() string { return proto.CompactTextString(m) }
func (*StartTaskRequest) ProtoMessage()    {}
func (*StartTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{5}
}

func (m *StartTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartTaskRequest.Unmarshal(m, b)
}
func (m *StartTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartTaskRequest.Marshal(b, m, deterministic)
}
func (m *StartTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartTaskRequest.Merge(m, src)
}
func (m *StartTaskRequest) XXX_Size() int {
	return xxx_messageInfo_StartTaskRequest.Size(m)
}
func (m *StartTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartTaskRequest proto.InternalMessageInfo

func (m *StartTaskRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *StartTaskRequest) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *StartTaskRequest) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *StartTaskRequest) GetParams() *TaskParams {
	if m != nil {
		return m.Params
	}
	return nil
}

// StopTaskRequest is message sent to a cluster member to stop a training task or predicting task.
type StopTaskRequest struct {
	TaskID               string      `protobuf:"bytes,2,opt,name=taskID,proto3" json:"taskID,omitempty"`
	Params               *TaskParams `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StopTaskRequest) Reset()         { *m = StopTaskRequest{} }
func (m *StopTaskRequest) String() string { return proto.CompactTextString(m) }
func (*StopTaskRequest) ProtoMessage()    {}
func (*StopTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{6}
}

func (m *StopTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopTaskRequest.Unmarshal(m, b)
}
func (m *StopTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopTaskRequest.Marshal(b, m, deterministic)
}
func (m *StopTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopTaskRequest.Merge(m, src)
}
func (m *StopTaskRequest) XXX_Size() int {
	return xxx_messageInfo_StopTaskRequest.Size(m)
}
func (m *StopTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopTaskRequest proto.InternalMessageInfo

func (m *StopTaskRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *StopTaskRequest) GetParams() *TaskParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.Algorithm", Algorithm_name, Algorithm_value)
	proto.RegisterEnum("common.TaskType", TaskType_name, TaskType_value)
	proto.RegisterEnum("common.RegMode", RegMode_name, RegMode_value)
	proto.RegisterType((*TrainParams)(nil), "common.TrainParams")
	proto.RegisterType((*TrainModels)(nil), "common.TrainModels")
	proto.RegisterMapType((map[string]float64)(nil), "common.TrainModels.SigmasEntry")
	proto.RegisterMapType((map[string]float64)(nil), "common.TrainModels.ThetasEntry")
	proto.RegisterMapType((map[string]float64)(nil), "common.TrainModels.XbarsEntry")
	proto.RegisterType((*TaskParams)(nil), "common.TaskParams")
	proto.RegisterType((*TrainTaskResult)(nil), "common.TrainTaskResult")
	proto.RegisterType((*PredictTaskResult)(nil), "common.PredictTaskResult")
	proto.RegisterType((*StartTaskRequest)(nil), "common.StartTaskRequest")
	proto.RegisterType((*StopTaskRequest)(nil), "common.StopTaskRequest")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdb, 0x6a, 0xe3, 0x46,
	0x18, 0x8e, 0x7c, 0x94, 0x7f, 0xa5, 0x89, 0x32, 0x09, 0xad, 0x30, 0xa5, 0x15, 0x86, 0x82, 0x6b,
	0x4a, 0x0c, 0x4e, 0x43, 0x93, 0x5e, 0x14, 0x72, 0x30, 0xc1, 0xe0, 0x38, 0x66, 0xec, 0x96, 0xd0,
	0x9b, 0x30, 0x96, 0x66, 0x65, 0x11, 0xc9, 0x72, 0x34, 0xa3, 0x65, 0xbd, 0x57, 0xfb, 0x40, 0x7b,
	0xb3, 0x8f, 0xb6, 0x6f, 0xb0, 0xcc, 0xc1, 0x96, 0x12, 0x12, 0x76, 0xc3, 0xde, 0x24, 0xf3, 0xfd,
	0x87, 0xef, 0x9b, 0xff, 0x30, 0x16, 0xec, 0x7b, 0x49, 0x1c, 0x27, 0x8b, 0xae, 0xfa, 0x77, 0xb8,
	0x4c, 0x13, 0x9e, 0xa0, 0x9a, 0x42, 0xad, 0x8f, 0x25, 0xb0, 0xa6, 0x29, 0x09, 0x17, 0x63, 0x92,
	0x92, 0x98, 0xa1, 0x03, 0xa8, 0x46, 0x64, 0x46, 0x23, 0xc7, 0x70, 0x8d, 0x76, 0x03, 0x2b, 0x80,
	0x7e, 0x86, 0x86, 0x3c, 0x8c, 0x48, 0x4c, 0x9d, 0x92, 0xf4, 0xe4, 0x06, 0xf4, 0x3b, 0xd4, 0x53,
	0x1a, 0x5c, 0x27, 0x3e, 0x75, 0xca, 0xae, 0xd1, 0xde, 0xe9, 0xed, 0x1e, 0x6a, 0x2d, 0xac, 0xcc,
	0x78, 0xed, 0x47, 0x4d, 0x30, 0x53, 0x1a, 0x48, 0x2d, 0xa7, 0xe2, 0x1a, 0x6d, 0x03, 0x6f, 0xb0,
	0x90, 0x26, 0xd1, 0x72, 0x4e, 0x9c, 0xaa, 0x74, 0x28, 0x20, 0xa4, 0x49, 0xbc, 0x8c, 0x42, 0x9e,
	0xf9, 0xd4, 0xa9, 0x49, 0x4f, 0x6e, 0x10, 0x7c, 0xc4, 0xf3, 0xb2, 0x94, 0x78, 0x2b, 0xa7, 0xee,
	0x1a, 0xed, 0x32, 0xde, 0x60, 0x91, 0x19, 0xb2, 0x29, 0x11, 0xec, 0xdc, 0x31, 0x5d, 0xa3, 0x6d,
	0xe2, 0xdc, 0x80, 0x7e, 0x84, 0x5a, 0xe8, 0xcb, 0x7a, 0x1a, 0xb2, 0x1e, 0x8d, 0x44, 0xd6, 0x39,
	0xe1, 0xde, 0x7c, 0x12, 0xbe, 0xa7, 0x0e, 0x48, 0xca, 0xdc, 0xd0, 0xfa, 0x54, 0xd6, 0xed, 0x12,
	0xd5, 0x44, 0x0c, 0xfd, 0x05, 0x35, 0x3e, 0xa7, 0x9c, 0x30, 0xc7, 0x70, 0xcb, 0x6d, 0xab, 0xf7,
	0xeb, 0xba, 0xf2, 0x42, 0xd0, 0xe1, 0x54, 0x46, 0xf4, 0x17, 0x3c, 0x5d, 0x61, 0x1d, 0x8e, 0xfe,
	0x84, 0xea, 0xbb, 0x19, 0x49, 0x99, 0x53, 0x92, 0x79, 0xbf, 0x3c, 0x97, 0x77, 0x2b, 0x02, 0x54,
	0x9a, 0x0a, 0x16, 0x72, 0x2c, 0x0c, 0x62, 0xc2, 0x9c, 0xf2, 0xcb, 0x72, 0x13, 0x19, 0xa1, 0xe5,
	0x54, 0x78, 0x3e, 0xd6, 0xca, 0x93, 0xb1, 0xe6, 0x1d, 0xaa, 0xbe, 0xdc, 0xa1, 0x5a, 0xb1, 0x43,
	0xcd, 0x53, 0xb0, 0x0a, 0x15, 0x21, 0x1b, 0xca, 0xf7, 0x74, 0xa5, 0xf7, 0x45, 0x1c, 0x85, 0xd8,
	0x5b, 0x12, 0x65, 0x6a, 0x53, 0x0c, 0xac, 0xc0, 0xdf, 0xa5, 0x13, 0xa3, 0x79, 0x02, 0x90, 0x17,
	0xf5, 0xaa, 0xcc, 0x53, 0xb0, 0x0a, 0x75, 0xbd, 0x26, 0xb5, 0xf5, 0xd9, 0x00, 0x98, 0x12, 0x76,
	0xaf, 0x37, 0xfc, 0x37, 0xa8, 0x90, 0x28, 0x48, 0x64, 0xee, 0x4e, 0x6f, 0x6f, 0xdd, 0xc1, 0xb3,
	0x28, 0x48, 0xd2, 0x90, 0xcf, 0x63, 0x2c, 0xdd, 0xe8, 0x0f, 0x30, 0x39, 0x61, 0xf7, 0xd3, 0xd5,
	0x52, 0x51, 0xee, 0xf4, 0xec, 0x4d, 0xb3, 0xb5, 0x1d, 0x6f, 0x22, 0xd0, 0x31, 0x58, 0x3c, 0x7f,
	0x45, 0xf2, 0x19, 0x58, 0xbd, 0xfd, 0x47, 0xd3, 0x51, 0x2e, 0x5c, 0x8c, 0x43, 0x2e, 0x58, 0xb1,
	0x18, 0x9a, 0x60, 0x1c, 0x5c, 0xea, 0xe1, 0x14, 0x4d, 0x82, 0x58, 0x42, 0x4d, 0x5c, 0x7d, 0x86,
	0x58, 0x8d, 0x1d, 0x17, 0xe3, 0x5a, 0x0f, 0xb0, 0x2b, 0x7d, 0x82, 0x05, 0x53, 0x96, 0x45, 0x72,
	0x9c, 0x5c, 0xc9, 0xa8, 0xae, 0x69, 0x84, 0x1c, 0xa8, 0xb3, 0xcc, 0xf3, 0x28, 0x63, 0xb2, 0x4e,
	0x13, 0xaf, 0xa1, 0x68, 0xa9, 0xe4, 0x94, 0xe5, 0x6c, 0x63, 0x05, 0x04, 0x0f, 0x4d, 0xd3, 0x6b,
	0x16, 0xe8, 0xeb, 0x6a, 0xd4, 0x5a, 0xc1, 0xde, 0x38, 0xa5, 0x7e, 0xe8, 0xf1, 0xef, 0x12, 0x6d,
	0x82, 0x99, 0x64, 0xdc, 0x4b, 0x62, 0xca, 0xb4, 0xee, 0x06, 0xbf, 0x28, 0xfd, 0xc1, 0x00, 0x7b,
	0xc2, 0x49, 0xaa, 0x95, 0x1f, 0x32, 0xca, 0x8a, 0xd2, 0xa5, 0x47, 0xd2, 0x08, 0x2a, 0x6f, 0xc2,
	0x88, 0x6a, 0x72, 0x79, 0x16, 0x95, 0xce, 0x13, 0xc6, 0x99, 0x53, 0x71, 0xcb, 0xe2, 0x79, 0x48,
	0x80, 0x3a, 0x50, 0x5b, 0x16, 0xdb, 0x8e, 0x8a, 0x0b, 0xa0, 0xc7, 0xa9, 0x23, 0x5a, 0xff, 0xc2,
	0xee, 0x84, 0x27, 0xcb, 0x6f, 0xb9, 0x40, 0x4e, 0x5b, 0xf9, 0x1a, 0x6d, 0xe7, 0x1f, 0x68, 0x6c,
	0x16, 0x13, 0x39, 0x70, 0x30, 0x1c, 0x8c, 0xfa, 0x67, 0xf8, 0x0e, 0xf7, 0xaf, 0x70, 0x7f, 0x32,
	0x19, 0xdc, 0x8c, 0xee, 0xfe, 0x1b, 0xda, 0x5b, 0xe8, 0x27, 0xd8, 0x1f, 0xde, 0x5c, 0x0d, 0x2e,
	0x9e, 0x38, 0x8c, 0x4e, 0x0b, 0xcc, 0xf5, 0xb6, 0xa2, 0x06, 0x54, 0x87, 0xfd, 0x33, 0x3c, 0xb2,
	0xb7, 0x90, 0x05, 0xf5, 0x31, 0xee, 0x5f, 0x0e, 0x2e, 0xa6, 0xb6, 0xd1, 0x39, 0x86, 0xba, 0xfe,
	0x9d, 0x46, 0xdb, 0x60, 0x62, 0x1a, 0xdc, 0x8d, 0x92, 0x05, 0xb5, 0xb7, 0xd0, 0x0f, 0xd0, 0x10,
	0x68, 0x48, 0x18, 0x4b, 0x6c, 0x63, 0x0d, 0x71, 0xe8, 0x07, 0xd4, 0x2e, 0x9d, 0x1f, 0xff, 0x7f,
	0x14, 0x84, 0x7c, 0x9e, 0xcd, 0xc4, 0xf5, 0xbb, 0x63, 0xe2, 0xfb, 0x11, 0x55, 0x7f, 0x35, 0xb8,
	0x9c, 0xde, 0x76, 0x7d, 0x12, 0x76, 0xe5, 0x07, 0x87, 0xe9, 0xcf, 0xcf, 0xac, 0x26, 0xe1, 0xd1,
	0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xb8, 0x4b, 0x20, 0x96, 0x06, 0x00, 0x00,
}
