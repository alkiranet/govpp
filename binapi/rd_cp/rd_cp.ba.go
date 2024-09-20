// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.06-release
// source: core/rd_cp.api.json

// Package rd_cp contains generated bindings for API file rd_cp.api.
//
// Contents:
// -  2 messages
package rd_cp

import (
	api "github.com/alkiranet/govpp/api"
	interface_types "github.com/alkiranet/govpp/binapi/interface_types"
	codec "github.com/alkiranet/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "rd_cp"
	APIVersion = "1.0.1"
	VersionCrc = 0x871c3bee
)

// Enable/disable IPv6 ND address autoconfiguration
//
//	       and setting up default routes
//	- sw_if_index - interface to enable the autoconfigutation on
//	- enable - 1 to enable address autoconfiguration, 0 to disable
//	- install_default_routes - 1 to enable installing defaut routes,
//	                                0 to disable installing defaut routes,
//	                                the value is ignored (and taken as 0)
//	                                when enable param is set to 0
//
// IP6NdAddressAutoconfig defines message 'ip6_nd_address_autoconfig'.
type IP6NdAddressAutoconfig struct {
	SwIfIndex            interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable               bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
	InstallDefaultRoutes bool                           `binapi:"bool,name=install_default_routes" json:"install_default_routes,omitempty"`
}

func (m *IP6NdAddressAutoconfig) Reset()               { *m = IP6NdAddressAutoconfig{} }
func (*IP6NdAddressAutoconfig) GetMessageName() string { return "ip6_nd_address_autoconfig" }
func (*IP6NdAddressAutoconfig) GetCrcString() string   { return "9e14a4a7" }
func (*IP6NdAddressAutoconfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IP6NdAddressAutoconfig) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	size += 1 // m.InstallDefaultRoutes
	return size
}
func (m *IP6NdAddressAutoconfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.Enable)
	buf.EncodeBool(m.InstallDefaultRoutes)
	return buf.Bytes(), nil
}
func (m *IP6NdAddressAutoconfig) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Enable = buf.DecodeBool()
	m.InstallDefaultRoutes = buf.DecodeBool()
	return nil
}

// IP6NdAddressAutoconfigReply defines message 'ip6_nd_address_autoconfig_reply'.
type IP6NdAddressAutoconfigReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IP6NdAddressAutoconfigReply) Reset()               { *m = IP6NdAddressAutoconfigReply{} }
func (*IP6NdAddressAutoconfigReply) GetMessageName() string { return "ip6_nd_address_autoconfig_reply" }
func (*IP6NdAddressAutoconfigReply) GetCrcString() string   { return "e8d4e804" }
func (*IP6NdAddressAutoconfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IP6NdAddressAutoconfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IP6NdAddressAutoconfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IP6NdAddressAutoconfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_rd_cp_binapi_init() }
func file_rd_cp_binapi_init() {
	api.RegisterMessage((*IP6NdAddressAutoconfig)(nil), "ip6_nd_address_autoconfig_9e14a4a7")
	api.RegisterMessage((*IP6NdAddressAutoconfigReply)(nil), "ip6_nd_address_autoconfig_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IP6NdAddressAutoconfig)(nil),
		(*IP6NdAddressAutoconfigReply)(nil),
	}
}
