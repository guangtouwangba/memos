// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1/auth_service.proto

package apiv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type GetCurrentSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCurrentSessionRequest) Reset() {
	*x = GetCurrentSessionRequest{}
	mi := &file_api_v1_auth_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSessionRequest) ProtoMessage() {}

func (x *GetCurrentSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSessionRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{0}
}

type GetCurrentSessionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	User  *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Last time the session was accessed.
	// Used for sliding expiration calculation (last_accessed_time + 2 weeks).
	LastAccessedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetCurrentSessionResponse) Reset() {
	*x = GetCurrentSessionResponse{}
	mi := &file_api_v1_auth_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCurrentSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSessionResponse) ProtoMessage() {}

func (x *GetCurrentSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSessionResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentSessionResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetCurrentSessionResponse) GetLastAccessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAccessedAt
	}
	return nil
}

type CreateSessionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Provide one authentication method (username/password or SSO).
	// Required field to specify the authentication method.
	//
	// Types that are valid to be assigned to Credentials:
	//
	//	*CreateSessionRequest_PasswordCredentials_
	//	*CreateSessionRequest_SsoCredentials
	Credentials   isCreateSessionRequest_Credentials `protobuf_oneof:"credentials"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	mi := &file_api_v1_auth_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSessionRequest) GetCredentials() isCreateSessionRequest_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *CreateSessionRequest) GetPasswordCredentials() *CreateSessionRequest_PasswordCredentials {
	if x != nil {
		if x, ok := x.Credentials.(*CreateSessionRequest_PasswordCredentials_); ok {
			return x.PasswordCredentials
		}
	}
	return nil
}

func (x *CreateSessionRequest) GetSsoCredentials() *CreateSessionRequest_SSOCredentials {
	if x != nil {
		if x, ok := x.Credentials.(*CreateSessionRequest_SsoCredentials); ok {
			return x.SsoCredentials
		}
	}
	return nil
}

type isCreateSessionRequest_Credentials interface {
	isCreateSessionRequest_Credentials()
}

type CreateSessionRequest_PasswordCredentials_ struct {
	// Username and password authentication method.
	PasswordCredentials *CreateSessionRequest_PasswordCredentials `protobuf:"bytes,1,opt,name=password_credentials,json=passwordCredentials,proto3,oneof"`
}

type CreateSessionRequest_SsoCredentials struct {
	// SSO provider authentication method.
	SsoCredentials *CreateSessionRequest_SSOCredentials `protobuf:"bytes,2,opt,name=sso_credentials,json=ssoCredentials,proto3,oneof"`
}

func (*CreateSessionRequest_PasswordCredentials_) isCreateSessionRequest_Credentials() {}

func (*CreateSessionRequest_SsoCredentials) isCreateSessionRequest_Credentials() {}

type CreateSessionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The authenticated user information.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Last time the session was accessed.
	// Used for sliding expiration calculation (last_accessed_time + 2 weeks).
	LastAccessedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	mi := &file_api_v1_auth_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSessionResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateSessionResponse) GetLastAccessedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAccessedAt
	}
	return nil
}

type DeleteSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSessionRequest) Reset() {
	*x = DeleteSessionRequest{}
	mi := &file_api_v1_auth_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSessionRequest) ProtoMessage() {}

func (x *DeleteSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSessionRequest.ProtoReflect.Descriptor instead.
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{4}
}

// Nested message for password-based authentication credentials.
type CreateSessionRequest_PasswordCredentials struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The username to sign in with.
	// Required field for password-based authentication.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// The password to sign in with.
	// Required field for password-based authentication.
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionRequest_PasswordCredentials) Reset() {
	*x = CreateSessionRequest_PasswordCredentials{}
	mi := &file_api_v1_auth_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionRequest_PasswordCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest_PasswordCredentials) ProtoMessage() {}

func (x *CreateSessionRequest_PasswordCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest_PasswordCredentials.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest_PasswordCredentials) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CreateSessionRequest_PasswordCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateSessionRequest_PasswordCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Nested message for SSO authentication credentials.
type CreateSessionRequest_SSOCredentials struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the SSO provider.
	// Required field to identify the SSO provider.
	IdpId int32 `protobuf:"varint,1,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	// The authorization code from the SSO provider.
	// Required field for completing the SSO flow.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// The redirect URI used in the SSO flow.
	// Required field for security validation.
	RedirectUri   string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSessionRequest_SSOCredentials) Reset() {
	*x = CreateSessionRequest_SSOCredentials{}
	mi := &file_api_v1_auth_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSessionRequest_SSOCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest_SSOCredentials) ProtoMessage() {}

func (x *CreateSessionRequest_SSOCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_auth_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest_SSOCredentials.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest_SSOCredentials) Descriptor() ([]byte, []int) {
	return file_api_v1_auth_service_proto_rawDescGZIP(), []int{2, 1}
}

func (x *CreateSessionRequest_SSOCredentials) GetIdpId() int32 {
	if x != nil {
		return x.IdpId
	}
	return 0
}

func (x *CreateSessionRequest_SSOCredentials) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateSessionRequest_SSOCredentials) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

var File_api_v1_auth_service_proto protoreflect.FileDescriptor

const file_api_v1_auth_service_proto_rawDesc = "" +
	"\n" +
	"\x19api/v1/auth_service.proto\x12\fmemos.api.v1\x1a\x19api/v1/user_service.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x1a\n" +
	"\x18GetCurrentSessionRequest\"\x89\x01\n" +
	"\x19GetCurrentSessionResponse\x12&\n" +
	"\x04user\x18\x01 \x01(\v2\x12.memos.api.v1.UserR\x04user\x12D\n" +
	"\x10last_accessed_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x0elastAccessedAt\"\xb8\x03\n" +
	"\x14CreateSessionRequest\x12k\n" +
	"\x14password_credentials\x18\x01 \x01(\v26.memos.api.v1.CreateSessionRequest.PasswordCredentialsH\x00R\x13passwordCredentials\x12\\\n" +
	"\x0fsso_credentials\x18\x02 \x01(\v21.memos.api.v1.CreateSessionRequest.SSOCredentialsH\x00R\x0essoCredentials\x1aW\n" +
	"\x13PasswordCredentials\x12\x1f\n" +
	"\busername\x18\x01 \x01(\tB\x03\xe0A\x02R\busername\x12\x1f\n" +
	"\bpassword\x18\x02 \x01(\tB\x03\xe0A\x02R\bpassword\x1am\n" +
	"\x0eSSOCredentials\x12\x1a\n" +
	"\x06idp_id\x18\x01 \x01(\x05B\x03\xe0A\x02R\x05idpId\x12\x17\n" +
	"\x04code\x18\x02 \x01(\tB\x03\xe0A\x02R\x04code\x12&\n" +
	"\fredirect_uri\x18\x03 \x01(\tB\x03\xe0A\x02R\vredirectUriB\r\n" +
	"\vcredentials\"\x85\x01\n" +
	"\x15CreateSessionResponse\x12&\n" +
	"\x04user\x18\x01 \x01(\v2\x12.memos.api.v1.UserR\x04user\x12D\n" +
	"\x10last_accessed_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x0elastAccessedAt\"\x16\n" +
	"\x14DeleteSessionRequest2\x8b\x03\n" +
	"\vAuthService\x12\x8b\x01\n" +
	"\x11GetCurrentSession\x12&.memos.api.v1.GetCurrentSessionRequest\x1a'.memos.api.v1.GetCurrentSessionResponse\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/api/v1/auth/sessions/current\x12z\n" +
	"\rCreateSession\x12\".memos.api.v1.CreateSessionRequest\x1a#.memos.api.v1.CreateSessionResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/api/v1/auth/sessions\x12r\n" +
	"\rDeleteSession\x12\".memos.api.v1.DeleteSessionRequest\x1a\x16.google.protobuf.Empty\"%\x82\xd3\xe4\x93\x02\x1f*\x1d/api/v1/auth/sessions/currentB\xa8\x01\n" +
	"\x10com.memos.api.v1B\x10AuthServiceProtoP\x01Z0github.com/usememos/memos/proto/gen/api/v1;apiv1\xa2\x02\x03MAX\xaa\x02\fMemos.Api.V1\xca\x02\fMemos\\Api\\V1\xe2\x02\x18Memos\\Api\\V1\\GPBMetadata\xea\x02\x0eMemos::Api::V1b\x06proto3"

var (
	file_api_v1_auth_service_proto_rawDescOnce sync.Once
	file_api_v1_auth_service_proto_rawDescData []byte
)

func file_api_v1_auth_service_proto_rawDescGZIP() []byte {
	file_api_v1_auth_service_proto_rawDescOnce.Do(func() {
		file_api_v1_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_auth_service_proto_rawDesc), len(file_api_v1_auth_service_proto_rawDesc)))
	})
	return file_api_v1_auth_service_proto_rawDescData
}

var file_api_v1_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_auth_service_proto_goTypes = []any{
	(*GetCurrentSessionRequest)(nil),                 // 0: memos.api.v1.GetCurrentSessionRequest
	(*GetCurrentSessionResponse)(nil),                // 1: memos.api.v1.GetCurrentSessionResponse
	(*CreateSessionRequest)(nil),                     // 2: memos.api.v1.CreateSessionRequest
	(*CreateSessionResponse)(nil),                    // 3: memos.api.v1.CreateSessionResponse
	(*DeleteSessionRequest)(nil),                     // 4: memos.api.v1.DeleteSessionRequest
	(*CreateSessionRequest_PasswordCredentials)(nil), // 5: memos.api.v1.CreateSessionRequest.PasswordCredentials
	(*CreateSessionRequest_SSOCredentials)(nil),      // 6: memos.api.v1.CreateSessionRequest.SSOCredentials
	(*User)(nil),                  // 7: memos.api.v1.User
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_api_v1_auth_service_proto_depIdxs = []int32{
	7, // 0: memos.api.v1.GetCurrentSessionResponse.user:type_name -> memos.api.v1.User
	8, // 1: memos.api.v1.GetCurrentSessionResponse.last_accessed_at:type_name -> google.protobuf.Timestamp
	5, // 2: memos.api.v1.CreateSessionRequest.password_credentials:type_name -> memos.api.v1.CreateSessionRequest.PasswordCredentials
	6, // 3: memos.api.v1.CreateSessionRequest.sso_credentials:type_name -> memos.api.v1.CreateSessionRequest.SSOCredentials
	7, // 4: memos.api.v1.CreateSessionResponse.user:type_name -> memos.api.v1.User
	8, // 5: memos.api.v1.CreateSessionResponse.last_accessed_at:type_name -> google.protobuf.Timestamp
	0, // 6: memos.api.v1.AuthService.GetCurrentSession:input_type -> memos.api.v1.GetCurrentSessionRequest
	2, // 7: memos.api.v1.AuthService.CreateSession:input_type -> memos.api.v1.CreateSessionRequest
	4, // 8: memos.api.v1.AuthService.DeleteSession:input_type -> memos.api.v1.DeleteSessionRequest
	1, // 9: memos.api.v1.AuthService.GetCurrentSession:output_type -> memos.api.v1.GetCurrentSessionResponse
	3, // 10: memos.api.v1.AuthService.CreateSession:output_type -> memos.api.v1.CreateSessionResponse
	9, // 11: memos.api.v1.AuthService.DeleteSession:output_type -> google.protobuf.Empty
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_auth_service_proto_init() }
func file_api_v1_auth_service_proto_init() {
	if File_api_v1_auth_service_proto != nil {
		return
	}
	file_api_v1_user_service_proto_init()
	file_api_v1_auth_service_proto_msgTypes[2].OneofWrappers = []any{
		(*CreateSessionRequest_PasswordCredentials_)(nil),
		(*CreateSessionRequest_SsoCredentials)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_auth_service_proto_rawDesc), len(file_api_v1_auth_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_auth_service_proto_goTypes,
		DependencyIndexes: file_api_v1_auth_service_proto_depIdxs,
		MessageInfos:      file_api_v1_auth_service_proto_msgTypes,
	}.Build()
	File_api_v1_auth_service_proto = out.File
	file_api_v1_auth_service_proto_goTypes = nil
	file_api_v1_auth_service_proto_depIdxs = nil
}
