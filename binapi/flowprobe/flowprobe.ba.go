// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.9.0
//  VPP:              23.10-release
// source: plugins/flowprobe.api.json

// Package flowprobe contains generated bindings for API file flowprobe.api.
//
// Contents:
// -  4 enums
// - 12 messages
package flowprobe

import (
	"strconv"

	api "go.fd.io/govpp/api"
	interface_types "go.fd.io/govpp/binapi/interface_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "flowprobe"
	APIVersion = "2.1.0"
	VersionCrc = 0x72c9c142
)

// FlowprobeDirection defines enum 'flowprobe_direction'.
type FlowprobeDirection uint8

const (
	FLOWPROBE_DIRECTION_RX   FlowprobeDirection = 0
	FLOWPROBE_DIRECTION_TX   FlowprobeDirection = 1
	FLOWPROBE_DIRECTION_BOTH FlowprobeDirection = 2
)

var (
	FlowprobeDirection_name = map[uint8]string{
		0: "FLOWPROBE_DIRECTION_RX",
		1: "FLOWPROBE_DIRECTION_TX",
		2: "FLOWPROBE_DIRECTION_BOTH",
	}
	FlowprobeDirection_value = map[string]uint8{
		"FLOWPROBE_DIRECTION_RX":   0,
		"FLOWPROBE_DIRECTION_TX":   1,
		"FLOWPROBE_DIRECTION_BOTH": 2,
	}
)

func (x FlowprobeDirection) String() string {
	s, ok := FlowprobeDirection_name[uint8(x)]
	if ok {
		return s
	}
	return "FlowprobeDirection(" + strconv.Itoa(int(x)) + ")"
}

// FlowprobeRecordFlags defines enum 'flowprobe_record_flags'.
type FlowprobeRecordFlags uint8

const (
	FLOWPROBE_RECORD_FLAG_L2 FlowprobeRecordFlags = 1
	FLOWPROBE_RECORD_FLAG_L3 FlowprobeRecordFlags = 2
	FLOWPROBE_RECORD_FLAG_L4 FlowprobeRecordFlags = 4
)

var (
	FlowprobeRecordFlags_name = map[uint8]string{
		1: "FLOWPROBE_RECORD_FLAG_L2",
		2: "FLOWPROBE_RECORD_FLAG_L3",
		4: "FLOWPROBE_RECORD_FLAG_L4",
	}
	FlowprobeRecordFlags_value = map[string]uint8{
		"FLOWPROBE_RECORD_FLAG_L2": 1,
		"FLOWPROBE_RECORD_FLAG_L3": 2,
		"FLOWPROBE_RECORD_FLAG_L4": 4,
	}
)

func (x FlowprobeRecordFlags) String() string {
	s, ok := FlowprobeRecordFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeRecordFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeRecordFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// FlowprobeWhich defines enum 'flowprobe_which'.
type FlowprobeWhich uint8

const (
	FLOWPROBE_WHICH_IP4 FlowprobeWhich = 0
	FLOWPROBE_WHICH_IP6 FlowprobeWhich = 1
	FLOWPROBE_WHICH_L2  FlowprobeWhich = 2
)

var (
	FlowprobeWhich_name = map[uint8]string{
		0: "FLOWPROBE_WHICH_IP4",
		1: "FLOWPROBE_WHICH_IP6",
		2: "FLOWPROBE_WHICH_L2",
	}
	FlowprobeWhich_value = map[string]uint8{
		"FLOWPROBE_WHICH_IP4": 0,
		"FLOWPROBE_WHICH_IP6": 1,
		"FLOWPROBE_WHICH_L2":  2,
	}
)

func (x FlowprobeWhich) String() string {
	s, ok := FlowprobeWhich_name[uint8(x)]
	if ok {
		return s
	}
	return "FlowprobeWhich(" + strconv.Itoa(int(x)) + ")"
}

// FlowprobeWhichFlags defines enum 'flowprobe_which_flags'.
type FlowprobeWhichFlags uint8

const (
	FLOWPROBE_WHICH_FLAG_IP4 FlowprobeWhichFlags = 1
	FLOWPROBE_WHICH_FLAG_L2  FlowprobeWhichFlags = 2
	FLOWPROBE_WHICH_FLAG_IP6 FlowprobeWhichFlags = 4
)

var (
	FlowprobeWhichFlags_name = map[uint8]string{
		1: "FLOWPROBE_WHICH_FLAG_IP4",
		2: "FLOWPROBE_WHICH_FLAG_L2",
		4: "FLOWPROBE_WHICH_FLAG_IP6",
	}
	FlowprobeWhichFlags_value = map[string]uint8{
		"FLOWPROBE_WHICH_FLAG_IP4": 1,
		"FLOWPROBE_WHICH_FLAG_L2":  2,
		"FLOWPROBE_WHICH_FLAG_IP6": 4,
	}
)

func (x FlowprobeWhichFlags) String() string {
	s, ok := FlowprobeWhichFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := FlowprobeWhichFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "FlowprobeWhichFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// Get IPFIX flow record generation parameters
// FlowprobeGetParams defines message 'flowprobe_get_params'.
// InProgress: the message form may change in the future versions
type FlowprobeGetParams struct{}

func (m *FlowprobeGetParams) Reset()               { *m = FlowprobeGetParams{} }
func (*FlowprobeGetParams) GetMessageName() string { return "flowprobe_get_params" }
func (*FlowprobeGetParams) GetCrcString() string   { return "51077d14" }
func (*FlowprobeGetParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeGetParams) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *FlowprobeGetParams) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *FlowprobeGetParams) Unmarshal(b []byte) error {
	return nil
}

// Reply to get IPFIX flow record generation parameters
//   - retval - error (0 is "no error")
//   - record_flags - flags indicating what data to record
//   - active_timer - time in seconds after which active flow records are
//     to be exported (0 is "off")
//   - passive_timer - time in seconds after which passive flow records are
//     to be deleted (0 is "off")
//
// FlowprobeGetParamsReply defines message 'flowprobe_get_params_reply'.
// InProgress: the message form may change in the future versions
type FlowprobeGetParamsReply struct {
	Retval       int32                `binapi:"i32,name=retval" json:"retval,omitempty"`
	RecordFlags  FlowprobeRecordFlags `binapi:"flowprobe_record_flags,name=record_flags" json:"record_flags,omitempty"`
	ActiveTimer  uint32               `binapi:"u32,name=active_timer" json:"active_timer,omitempty"`
	PassiveTimer uint32               `binapi:"u32,name=passive_timer" json:"passive_timer,omitempty"`
}

func (m *FlowprobeGetParamsReply) Reset()               { *m = FlowprobeGetParamsReply{} }
func (*FlowprobeGetParamsReply) GetMessageName() string { return "flowprobe_get_params_reply" }
func (*FlowprobeGetParamsReply) GetCrcString() string   { return "f350d621" }
func (*FlowprobeGetParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeGetParamsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 1 // m.RecordFlags
	size += 4 // m.ActiveTimer
	size += 4 // m.PassiveTimer
	return size
}
func (m *FlowprobeGetParamsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint8(uint8(m.RecordFlags))
	buf.EncodeUint32(m.ActiveTimer)
	buf.EncodeUint32(m.PassiveTimer)
	return buf.Bytes(), nil
}
func (m *FlowprobeGetParamsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.RecordFlags = FlowprobeRecordFlags(buf.DecodeUint8())
	m.ActiveTimer = buf.DecodeUint32()
	m.PassiveTimer = buf.DecodeUint32()
	return nil
}

// Enable or disable IPFIX flow record generation on an interface
//   - is_add - add interface if non-zero, else delete
//   - which - datapath on which to record flows
//   - direction - direction of recorded flows
//   - sw_if_index - index of the interface
//
// FlowprobeInterfaceAddDel defines message 'flowprobe_interface_add_del'.
// InProgress: the message form may change in the future versions
type FlowprobeInterfaceAddDel struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Which     FlowprobeWhich                 `binapi:"flowprobe_which,name=which" json:"which,omitempty"`
	Direction FlowprobeDirection             `binapi:"flowprobe_direction,name=direction" json:"direction,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeInterfaceAddDel) Reset()               { *m = FlowprobeInterfaceAddDel{} }
func (*FlowprobeInterfaceAddDel) GetMessageName() string { return "flowprobe_interface_add_del" }
func (*FlowprobeInterfaceAddDel) GetCrcString() string   { return "3420739c" }
func (*FlowprobeInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeInterfaceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.Which
	size += 1 // m.Direction
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeInterfaceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Which))
	buf.EncodeUint8(uint8(m.Direction))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeInterfaceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Which = FlowprobeWhich(buf.DecodeUint8())
	m.Direction = FlowprobeDirection(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// FlowprobeInterfaceAddDelReply defines message 'flowprobe_interface_add_del_reply'.
// InProgress: the message form may change in the future versions
type FlowprobeInterfaceAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeInterfaceAddDelReply) Reset() { *m = FlowprobeInterfaceAddDelReply{} }
func (*FlowprobeInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_interface_add_del_reply"
}
func (*FlowprobeInterfaceAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*FlowprobeInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeInterfaceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeInterfaceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeInterfaceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Details about IPFIX flow record generation enabled on interface
//   - which - datapath on which to record flows
//   - direction - direction of recorded flows
//   - sw_if_index - index of the interface
//
// FlowprobeInterfaceDetails defines message 'flowprobe_interface_details'.
// InProgress: the message form may change in the future versions
type FlowprobeInterfaceDetails struct {
	Which     FlowprobeWhich                 `binapi:"flowprobe_which,name=which" json:"which,omitempty"`
	Direction FlowprobeDirection             `binapi:"flowprobe_direction,name=direction" json:"direction,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeInterfaceDetails) Reset()               { *m = FlowprobeInterfaceDetails{} }
func (*FlowprobeInterfaceDetails) GetMessageName() string { return "flowprobe_interface_details" }
func (*FlowprobeInterfaceDetails) GetCrcString() string   { return "427d77e0" }
func (*FlowprobeInterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeInterfaceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Which
	size += 1 // m.Direction
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeInterfaceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Which))
	buf.EncodeUint8(uint8(m.Direction))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeInterfaceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Which = FlowprobeWhich(buf.DecodeUint8())
	m.Direction = FlowprobeDirection(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Dump interfaces for which IPFIX flow record generation is enabled
//   - sw_if_index - interface index to use as filter (0xffffffff is "all")
//
// FlowprobeInterfaceDump defines message 'flowprobe_interface_dump'.
// InProgress: the message form may change in the future versions
type FlowprobeInterfaceDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeInterfaceDump) Reset()               { *m = FlowprobeInterfaceDump{} }
func (*FlowprobeInterfaceDump) GetMessageName() string { return "flowprobe_interface_dump" }
func (*FlowprobeInterfaceDump) GetCrcString() string   { return "f9e6675e" }
func (*FlowprobeInterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeInterfaceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeInterfaceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeInterfaceDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// FlowprobeParams defines message 'flowprobe_params'.
type FlowprobeParams struct {
	RecordFlags  FlowprobeRecordFlags `binapi:"flowprobe_record_flags,name=record_flags" json:"record_flags,omitempty"`
	ActiveTimer  uint32               `binapi:"u32,name=active_timer" json:"active_timer,omitempty"`
	PassiveTimer uint32               `binapi:"u32,name=passive_timer" json:"passive_timer,omitempty"`
}

func (m *FlowprobeParams) Reset()               { *m = FlowprobeParams{} }
func (*FlowprobeParams) GetMessageName() string { return "flowprobe_params" }
func (*FlowprobeParams) GetCrcString() string   { return "baa46c09" }
func (*FlowprobeParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeParams) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.RecordFlags
	size += 4 // m.ActiveTimer
	size += 4 // m.PassiveTimer
	return size
}
func (m *FlowprobeParams) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.RecordFlags))
	buf.EncodeUint32(m.ActiveTimer)
	buf.EncodeUint32(m.PassiveTimer)
	return buf.Bytes(), nil
}
func (m *FlowprobeParams) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RecordFlags = FlowprobeRecordFlags(buf.DecodeUint8())
	m.ActiveTimer = buf.DecodeUint32()
	m.PassiveTimer = buf.DecodeUint32()
	return nil
}

// FlowprobeParamsReply defines message 'flowprobe_params_reply'.
type FlowprobeParamsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeParamsReply) Reset()               { *m = FlowprobeParamsReply{} }
func (*FlowprobeParamsReply) GetMessageName() string { return "flowprobe_params_reply" }
func (*FlowprobeParamsReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowprobeParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeParamsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeParamsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeParamsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Set IPFIX flow record generation parameters
//   - record_flags - flags indicating what data to record
//   - active_timer - time in seconds after which active flow records are
//     to be exported (0 is "off", 0xffffffff is "use default value")
//   - passive_timer - time in seconds after which passive flow records are
//     to be deleted (0 is "off", 0xffffffff is "use default value")
//
// FlowprobeSetParams defines message 'flowprobe_set_params'.
// InProgress: the message form may change in the future versions
type FlowprobeSetParams struct {
	RecordFlags  FlowprobeRecordFlags `binapi:"flowprobe_record_flags,name=record_flags" json:"record_flags,omitempty"`
	ActiveTimer  uint32               `binapi:"u32,name=active_timer,default=4294967295" json:"active_timer,omitempty"`
	PassiveTimer uint32               `binapi:"u32,name=passive_timer,default=4294967295" json:"passive_timer,omitempty"`
}

func (m *FlowprobeSetParams) Reset()               { *m = FlowprobeSetParams{} }
func (*FlowprobeSetParams) GetMessageName() string { return "flowprobe_set_params" }
func (*FlowprobeSetParams) GetCrcString() string   { return "baa46c09" }
func (*FlowprobeSetParams) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeSetParams) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.RecordFlags
	size += 4 // m.ActiveTimer
	size += 4 // m.PassiveTimer
	return size
}
func (m *FlowprobeSetParams) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.RecordFlags))
	buf.EncodeUint32(m.ActiveTimer)
	buf.EncodeUint32(m.PassiveTimer)
	return buf.Bytes(), nil
}
func (m *FlowprobeSetParams) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RecordFlags = FlowprobeRecordFlags(buf.DecodeUint8())
	m.ActiveTimer = buf.DecodeUint32()
	m.PassiveTimer = buf.DecodeUint32()
	return nil
}

// FlowprobeSetParamsReply defines message 'flowprobe_set_params_reply'.
// InProgress: the message form may change in the future versions
type FlowprobeSetParamsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeSetParamsReply) Reset()               { *m = FlowprobeSetParamsReply{} }
func (*FlowprobeSetParamsReply) GetMessageName() string { return "flowprobe_set_params_reply" }
func (*FlowprobeSetParamsReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowprobeSetParamsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeSetParamsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeSetParamsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeSetParamsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Enable / disable per-packet IPFIX recording on an interface
//   - is_add - add address if non-zero, else delete
//   - which - flags indicating forwarding path
//   - sw_if_index - index of the interface
//
// FlowprobeTxInterfaceAddDel defines message 'flowprobe_tx_interface_add_del'.
type FlowprobeTxInterfaceAddDel struct {
	IsAdd     bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Which     FlowprobeWhichFlags            `binapi:"flowprobe_which_flags,name=which" json:"which,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDel) Reset()               { *m = FlowprobeTxInterfaceAddDel{} }
func (*FlowprobeTxInterfaceAddDel) GetMessageName() string { return "flowprobe_tx_interface_add_del" }
func (*FlowprobeTxInterfaceAddDel) GetCrcString() string   { return "b782c976" }
func (*FlowprobeTxInterfaceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowprobeTxInterfaceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.Which
	size += 4 // m.SwIfIndex
	return size
}
func (m *FlowprobeTxInterfaceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.Which))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Which = FlowprobeWhichFlags(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// FlowprobeTxInterfaceAddDelReply defines message 'flowprobe_tx_interface_add_del_reply'.
type FlowprobeTxInterfaceAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowprobeTxInterfaceAddDelReply) Reset() { *m = FlowprobeTxInterfaceAddDelReply{} }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageName() string {
	return "flowprobe_tx_interface_add_del_reply"
}
func (*FlowprobeTxInterfaceAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*FlowprobeTxInterfaceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowprobeTxInterfaceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowprobeTxInterfaceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowprobeTxInterfaceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_flowprobe_binapi_init() }
func file_flowprobe_binapi_init() {
	api.RegisterMessage((*FlowprobeGetParams)(nil), "flowprobe_get_params_51077d14")
	api.RegisterMessage((*FlowprobeGetParamsReply)(nil), "flowprobe_get_params_reply_f350d621")
	api.RegisterMessage((*FlowprobeInterfaceAddDel)(nil), "flowprobe_interface_add_del_3420739c")
	api.RegisterMessage((*FlowprobeInterfaceAddDelReply)(nil), "flowprobe_interface_add_del_reply_e8d4e804")
	api.RegisterMessage((*FlowprobeInterfaceDetails)(nil), "flowprobe_interface_details_427d77e0")
	api.RegisterMessage((*FlowprobeInterfaceDump)(nil), "flowprobe_interface_dump_f9e6675e")
	api.RegisterMessage((*FlowprobeParams)(nil), "flowprobe_params_baa46c09")
	api.RegisterMessage((*FlowprobeParamsReply)(nil), "flowprobe_params_reply_e8d4e804")
	api.RegisterMessage((*FlowprobeSetParams)(nil), "flowprobe_set_params_baa46c09")
	api.RegisterMessage((*FlowprobeSetParamsReply)(nil), "flowprobe_set_params_reply_e8d4e804")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDel)(nil), "flowprobe_tx_interface_add_del_b782c976")
	api.RegisterMessage((*FlowprobeTxInterfaceAddDelReply)(nil), "flowprobe_tx_interface_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowprobeGetParams)(nil),
		(*FlowprobeGetParamsReply)(nil),
		(*FlowprobeInterfaceAddDel)(nil),
		(*FlowprobeInterfaceAddDelReply)(nil),
		(*FlowprobeInterfaceDetails)(nil),
		(*FlowprobeInterfaceDump)(nil),
		(*FlowprobeParams)(nil),
		(*FlowprobeParamsReply)(nil),
		(*FlowprobeSetParams)(nil),
		(*FlowprobeSetParamsReply)(nil),
		(*FlowprobeTxInterfaceAddDel)(nil),
		(*FlowprobeTxInterfaceAddDelReply)(nil),
	}
}
