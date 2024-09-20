// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.8.0
//  VPP:              23.06-release
// source: core/sr_pt.api.json

// Package sr_pt contains generated bindings for API file sr_pt.api.
//
// Contents:
// -  6 messages
package sr_pt

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
	APIFile    = "sr_pt"
	APIVersion = "1.0.0"
	VersionCrc = 0x1fddedad
)

// SR PT iface add request
//   - sw_if_index - index of the interface to add to SR PT
//   - id - SR PT interface id
//   - ingress_load - incoming interface load
//   - egress_load - outgoing interface load
//   - tts_template - truncated timestamp template to use
//
// SrPtIfaceAdd defines message 'sr_pt_iface_add'.
type SrPtIfaceAdd struct {
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	ID          uint16                         `binapi:"u16,name=id" json:"id,omitempty"`
	IngressLoad uint8                          `binapi:"u8,name=ingress_load" json:"ingress_load,omitempty"`
	EgressLoad  uint8                          `binapi:"u8,name=egress_load" json:"egress_load,omitempty"`
	TtsTemplate uint8                          `binapi:"u8,name=tts_template" json:"tts_template,omitempty"`
}

func (m *SrPtIfaceAdd) Reset()               { *m = SrPtIfaceAdd{} }
func (*SrPtIfaceAdd) GetMessageName() string { return "sr_pt_iface_add" }
func (*SrPtIfaceAdd) GetCrcString() string   { return "852c0cda" }
func (*SrPtIfaceAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPtIfaceAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 2 // m.ID
	size += 1 // m.IngressLoad
	size += 1 // m.EgressLoad
	size += 1 // m.TtsTemplate
	return size
}
func (m *SrPtIfaceAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint16(m.ID)
	buf.EncodeUint8(m.IngressLoad)
	buf.EncodeUint8(m.EgressLoad)
	buf.EncodeUint8(m.TtsTemplate)
	return buf.Bytes(), nil
}
func (m *SrPtIfaceAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.ID = buf.DecodeUint16()
	m.IngressLoad = buf.DecodeUint8()
	m.EgressLoad = buf.DecodeUint8()
	m.TtsTemplate = buf.DecodeUint8()
	return nil
}

// SrPtIfaceAddReply defines message 'sr_pt_iface_add_reply'.
type SrPtIfaceAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPtIfaceAddReply) Reset()               { *m = SrPtIfaceAddReply{} }
func (*SrPtIfaceAddReply) GetMessageName() string { return "sr_pt_iface_add_reply" }
func (*SrPtIfaceAddReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPtIfaceAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPtIfaceAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPtIfaceAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPtIfaceAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SR PT iface del request
//   - sw_if_index - index of the interface to delete from SR PT
//
// SrPtIfaceDel defines message 'sr_pt_iface_del'.
type SrPtIfaceDel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *SrPtIfaceDel) Reset()               { *m = SrPtIfaceDel{} }
func (*SrPtIfaceDel) GetMessageName() string { return "sr_pt_iface_del" }
func (*SrPtIfaceDel) GetCrcString() string   { return "f9e6675e" }
func (*SrPtIfaceDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPtIfaceDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *SrPtIfaceDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SrPtIfaceDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// SrPtIfaceDelReply defines message 'sr_pt_iface_del_reply'.
type SrPtIfaceDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrPtIfaceDelReply) Reset()               { *m = SrPtIfaceDelReply{} }
func (*SrPtIfaceDelReply) GetMessageName() string { return "sr_pt_iface_del_reply" }
func (*SrPtIfaceDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrPtIfaceDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPtIfaceDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrPtIfaceDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrPtIfaceDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SrPtIfaceDetails defines message 'sr_pt_iface_details'.
type SrPtIfaceDetails struct {
	SwIfIndex   interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	ID          uint16                         `binapi:"u16,name=id" json:"id,omitempty"`
	IngressLoad uint8                          `binapi:"u8,name=ingress_load" json:"ingress_load,omitempty"`
	EgressLoad  uint8                          `binapi:"u8,name=egress_load" json:"egress_load,omitempty"`
	TtsTemplate uint8                          `binapi:"u8,name=tts_template" json:"tts_template,omitempty"`
}

func (m *SrPtIfaceDetails) Reset()               { *m = SrPtIfaceDetails{} }
func (*SrPtIfaceDetails) GetMessageName() string { return "sr_pt_iface_details" }
func (*SrPtIfaceDetails) GetCrcString() string   { return "1f472f85" }
func (*SrPtIfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrPtIfaceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 2 // m.ID
	size += 1 // m.IngressLoad
	size += 1 // m.EgressLoad
	size += 1 // m.TtsTemplate
	return size
}
func (m *SrPtIfaceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint16(m.ID)
	buf.EncodeUint8(m.IngressLoad)
	buf.EncodeUint8(m.EgressLoad)
	buf.EncodeUint8(m.TtsTemplate)
	return buf.Bytes(), nil
}
func (m *SrPtIfaceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.ID = buf.DecodeUint16()
	m.IngressLoad = buf.DecodeUint8()
	m.EgressLoad = buf.DecodeUint8()
	m.TtsTemplate = buf.DecodeUint8()
	return nil
}

// SR PT iface dump request
// SrPtIfaceDump defines message 'sr_pt_iface_dump'.
type SrPtIfaceDump struct{}

func (m *SrPtIfaceDump) Reset()               { *m = SrPtIfaceDump{} }
func (*SrPtIfaceDump) GetMessageName() string { return "sr_pt_iface_dump" }
func (*SrPtIfaceDump) GetCrcString() string   { return "51077d14" }
func (*SrPtIfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrPtIfaceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *SrPtIfaceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *SrPtIfaceDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_sr_pt_binapi_init() }
func file_sr_pt_binapi_init() {
	api.RegisterMessage((*SrPtIfaceAdd)(nil), "sr_pt_iface_add_852c0cda")
	api.RegisterMessage((*SrPtIfaceAddReply)(nil), "sr_pt_iface_add_reply_e8d4e804")
	api.RegisterMessage((*SrPtIfaceDel)(nil), "sr_pt_iface_del_f9e6675e")
	api.RegisterMessage((*SrPtIfaceDelReply)(nil), "sr_pt_iface_del_reply_e8d4e804")
	api.RegisterMessage((*SrPtIfaceDetails)(nil), "sr_pt_iface_details_1f472f85")
	api.RegisterMessage((*SrPtIfaceDump)(nil), "sr_pt_iface_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrPtIfaceAdd)(nil),
		(*SrPtIfaceAddReply)(nil),
		(*SrPtIfaceDel)(nil),
		(*SrPtIfaceDelReply)(nil),
		(*SrPtIfaceDetails)(nil),
		(*SrPtIfaceDump)(nil),
	}
}
