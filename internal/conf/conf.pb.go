// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.28.3
// source: conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AppMetadata_Environment int32

const (
	AppMetadata_NONE  AppMetadata_Environment = 0
	AppMetadata_DEV   AppMetadata_Environment = 1
	AppMetadata_STAGE AppMetadata_Environment = 2
	AppMetadata_PROD  AppMetadata_Environment = 3
)

// Enum value maps for AppMetadata_Environment.
var (
	AppMetadata_Environment_name = map[int32]string{
		0: "NONE",
		1: "DEV",
		2: "STAGE",
		3: "PROD",
	}
	AppMetadata_Environment_value = map[string]int32{
		"NONE":  0,
		"DEV":   1,
		"STAGE": 2,
		"PROD":  3,
	}
)

func (x AppMetadata_Environment) Enum() *AppMetadata_Environment {
	p := new(AppMetadata_Environment)
	*p = x
	return p
}

func (x AppMetadata_Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppMetadata_Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_conf_conf_proto_enumTypes[0].Descriptor()
}

func (AppMetadata_Environment) Type() protoreflect.EnumType {
	return &file_conf_conf_proto_enumTypes[0]
}

func (x AppMetadata_Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppMetadata_Environment.Descriptor instead.
func (AppMetadata_Environment) EnumDescriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

type Log_Logger int32

const (
	Log_ZAP    Log_Logger = 0
	Log_LOGRUS Log_Logger = 1
)

// Enum value maps for Log_Logger.
var (
	Log_Logger_name = map[int32]string{
		0: "ZAP",
		1: "LOGRUS",
	}
	Log_Logger_value = map[string]int32{
		"ZAP":    0,
		"LOGRUS": 1,
	}
)

func (x Log_Logger) Enum() *Log_Logger {
	p := new(Log_Logger)
	*p = x
	return p
}

func (x Log_Logger) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Log_Logger) Descriptor() protoreflect.EnumDescriptor {
	return file_conf_conf_proto_enumTypes[1].Descriptor()
}

func (Log_Logger) Type() protoreflect.EnumType {
	return &file_conf_conf_proto_enumTypes[1]
}

func (x Log_Logger) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Log_Logger.Descriptor instead.
func (Log_Logger) EnumDescriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3, 0}
}

type Bootstrap struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Server        *Server                `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data          *Data                  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Metadata      *AppMetadata           `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Otel          *Otel                  `protobuf:"bytes,4,opt,name=otel,proto3" json:"otel,omitempty"`
	Log           *Log                   `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_conf_conf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetMetadata() *AppMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Bootstrap) GetOtel() *Otel {
	if x != nil {
		return x.Otel
	}
	return nil
}

func (x *Bootstrap) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

type AppMetadata struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Name          string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Env           AppMetadata_Environment `protobuf:"varint,2,opt,name=env,proto3,enum=kratos.api.AppMetadata_Environment" json:"env,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppMetadata) Reset() {
	*x = AppMetadata{}
	mi := &file_conf_conf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMetadata) ProtoMessage() {}

func (x *AppMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMetadata.ProtoReflect.Descriptor instead.
func (*AppMetadata) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *AppMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppMetadata) GetEnv() AppMetadata_Environment {
	if x != nil {
		return x.Env
	}
	return AppMetadata_NONE
}

type Otel struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Trace         *Otel_Trace            `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	Metrics       *Otel_Metrics          `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Otel) Reset() {
	*x = Otel{}
	mi := &file_conf_conf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Otel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Otel) ProtoMessage() {}

func (x *Otel) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Otel.ProtoReflect.Descriptor instead.
func (*Otel) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Otel) GetTrace() *Otel_Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Otel) GetMetrics() *Otel_Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         string                 `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	Filepath      string                 `protobuf:"bytes,2,opt,name=filepath,proto3" json:"filepath,omitempty"`
	Logger        string                 `protobuf:"bytes,3,opt,name=logger,proto3" json:"logger,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Log) Reset() {
	*x = Log{}
	mi := &file_conf_conf_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Log) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *Log) GetLogger() string {
	if x != nil {
		return x.Logger
	}
	return ""
}

type Server struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server) Reset() {
	*x = Server{}
	mi := &file_conf_conf_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis         *Data_Redis            `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	Mongo         *Data_Mongo            `protobuf:"bytes,3,opt,name=mongo,proto3" json:"mongo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_conf_conf_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data) GetMongo() *Data_Mongo {
	if x != nil {
		return x.Mongo
	}
	return nil
}

type Otel_Trace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Endpoint      string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Insecure      bool                   `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Otel_Trace) Reset() {
	*x = Otel_Trace{}
	mi := &file_conf_conf_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Otel_Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Otel_Trace) ProtoMessage() {}

func (x *Otel_Trace) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Otel_Trace.ProtoReflect.Descriptor instead.
func (*Otel_Trace) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Otel_Trace) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Otel_Trace) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

type Otel_Metrics struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	EnableExemplar bool                   `protobuf:"varint,1,opt,name=enable_exemplar,json=enableExemplar,proto3" json:"enable_exemplar,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Otel_Metrics) Reset() {
	*x = Otel_Metrics{}
	mi := &file_conf_conf_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Otel_Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Otel_Metrics) ProtoMessage() {}

func (x *Otel_Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Otel_Metrics.ProtoReflect.Descriptor instead.
func (*Otel_Metrics) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Otel_Metrics) GetEnableExemplar() bool {
	if x != nil {
		return x.EnableExemplar
	}
	return false
}

type Server_HTTP struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	mi := &file_conf_conf_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout       *durationpb.Duration   `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	mi := &file_conf_conf_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{4, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source        string                 `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	mi := &file_conf_conf_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Network       string                 `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout   *durationpb.Duration   `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout  *durationpb.Duration   `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	mi := &file_conf_conf_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Data_Mongo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uri           string                 `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Database      string                 `protobuf:"bytes,4,opt,name=database,proto3" json:"database,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data_Mongo) Reset() {
	*x = Data_Mongo{}
	mi := &file_conf_conf_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data_Mongo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Mongo) ProtoMessage() {}

func (x *Data_Mongo) ProtoReflect() protoreflect.Message {
	mi := &file_conf_conf_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Mongo.ProtoReflect.Descriptor instead.
func (*Data_Mongo) Descriptor() ([]byte, []int) {
	return file_conf_conf_proto_rawDescGZIP(), []int{5, 2}
}

func (x *Data_Mongo) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Data_Mongo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Data_Mongo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Mongo) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

var File_conf_conf_proto protoreflect.FileDescriptor

var file_conf_conf_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x6f, 0x74, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x74,
	0x65, 0x6c, 0x52, 0x04, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x8f, 0x01, 0x0a, 0x0b,
	0x41, 0x70, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22, 0x35, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x44, 0x45, 0x56, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x47,
	0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x4f, 0x44, 0x10, 0x03, 0x22, 0xdd, 0x01,
	0x0a, 0x04, 0x4f, 0x74, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4f, 0x74, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4f, 0x74, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x3f, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x1a, 0x32, 0x0a, 0x07, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65,
	0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x72, 0x22, 0x6e, 0x0a,
	0x03, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x22, 0x1d,
	0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x5a, 0x41, 0x50, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x52, 0x55, 0x53, 0x10, 0x01, 0x22, 0xb8, 0x02,
	0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x69, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x69, 0x0a,
	0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xfa, 0x03, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52,
	0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x67, 0x6f, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x1a, 0xb3, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x6d, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_conf_conf_proto_rawDescOnce sync.Once
	file_conf_conf_proto_rawDescData []byte
)

func file_conf_conf_proto_rawDescGZIP() []byte {
	file_conf_conf_proto_rawDescOnce.Do(func() {
		file_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)))
	})
	return file_conf_conf_proto_rawDescData
}

var file_conf_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_conf_conf_proto_goTypes = []any{
	(AppMetadata_Environment)(0), // 0: kratos.api.AppMetadata.Environment
	(Log_Logger)(0),              // 1: kratos.api.Log.Logger
	(*Bootstrap)(nil),            // 2: kratos.api.Bootstrap
	(*AppMetadata)(nil),          // 3: kratos.api.AppMetadata
	(*Otel)(nil),                 // 4: kratos.api.Otel
	(*Log)(nil),                  // 5: kratos.api.Log
	(*Server)(nil),               // 6: kratos.api.Server
	(*Data)(nil),                 // 7: kratos.api.Data
	(*Otel_Trace)(nil),           // 8: kratos.api.Otel.Trace
	(*Otel_Metrics)(nil),         // 9: kratos.api.Otel.Metrics
	(*Server_HTTP)(nil),          // 10: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),          // 11: kratos.api.Server.GRPC
	(*Data_Database)(nil),        // 12: kratos.api.Data.Database
	(*Data_Redis)(nil),           // 13: kratos.api.Data.Redis
	(*Data_Mongo)(nil),           // 14: kratos.api.Data.Mongo
	(*durationpb.Duration)(nil),  // 15: google.protobuf.Duration
}
var file_conf_conf_proto_depIdxs = []int32{
	6,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	7,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	3,  // 2: kratos.api.Bootstrap.metadata:type_name -> kratos.api.AppMetadata
	4,  // 3: kratos.api.Bootstrap.otel:type_name -> kratos.api.Otel
	5,  // 4: kratos.api.Bootstrap.log:type_name -> kratos.api.Log
	0,  // 5: kratos.api.AppMetadata.env:type_name -> kratos.api.AppMetadata.Environment
	8,  // 6: kratos.api.Otel.trace:type_name -> kratos.api.Otel.Trace
	9,  // 7: kratos.api.Otel.metrics:type_name -> kratos.api.Otel.Metrics
	10, // 8: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	11, // 9: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	12, // 10: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	13, // 11: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	14, // 12: kratos.api.Data.mongo:type_name -> kratos.api.Data.Mongo
	15, // 13: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	15, // 14: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	15, // 15: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	15, // 16: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_conf_conf_proto_init() }
func file_conf_conf_proto_init() {
	if File_conf_conf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_conf_proto_goTypes,
		DependencyIndexes: file_conf_conf_proto_depIdxs,
		EnumInfos:         file_conf_conf_proto_enumTypes,
		MessageInfos:      file_conf_conf_proto_msgTypes,
	}.Build()
	File_conf_conf_proto = out.File
	file_conf_conf_proto_goTypes = nil
	file_conf_conf_proto_depIdxs = nil
}
