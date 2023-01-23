// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package wireguard contains generated bindings for API file wireguard.api.
//
// Contents:
// -  1 enum
// -  2 structs
// - 17 messages
package wireguard

import (
	"strconv"

	interface_types "github.com/projectcalico/vpp-dataplane/vpplink/generated/bindings/interface_types"
	ip_types "github.com/projectcalico/vpp-dataplane/vpplink/generated/bindings/ip_types"
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "wireguard"
	APIVersion = "0.3.0"
	VersionCrc = 0x5d8f9252
)

// WireguardPeerFlags defines enum 'wireguard_peer_flags'.
type WireguardPeerFlags uint8

const (
	WIREGUARD_PEER_STATUS_DEAD WireguardPeerFlags = 1
	WIREGUARD_PEER_ESTABLISHED WireguardPeerFlags = 2
)

var (
	WireguardPeerFlags_name = map[uint8]string{
		1: "WIREGUARD_PEER_STATUS_DEAD",
		2: "WIREGUARD_PEER_ESTABLISHED",
	}
	WireguardPeerFlags_value = map[string]uint8{
		"WIREGUARD_PEER_STATUS_DEAD": 1,
		"WIREGUARD_PEER_ESTABLISHED": 2,
	}
)

func (x WireguardPeerFlags) String() string {
	s, ok := WireguardPeerFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := WireguardPeerFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "WireguardPeerFlags(" + strconv.Itoa(int(n)) + ")"
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

// WireguardInterface defines type 'wireguard_interface'.
type WireguardInterface struct {
	UserInstance uint32                         `binapi:"u32,name=user_instance,default=4294967295" json:"user_instance,omitempty"`
	SwIfIndex    interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	PrivateKey   []byte                         `binapi:"u8[32],name=private_key" json:"private_key,omitempty"`
	PublicKey    []byte                         `binapi:"u8[32],name=public_key" json:"public_key,omitempty"`
	Port         uint16                         `binapi:"u16,name=port" json:"port,omitempty"`
	SrcIP        ip_types.Address               `binapi:"address,name=src_ip" json:"src_ip,omitempty"`
}

// WireguardPeer defines type 'wireguard_peer'.
type WireguardPeer struct {
	PeerIndex           uint32                         `binapi:"u32,name=peer_index" json:"peer_index,omitempty"`
	PublicKey           []byte                         `binapi:"u8[32],name=public_key" json:"public_key,omitempty"`
	Port                uint16                         `binapi:"u16,name=port" json:"port,omitempty"`
	PersistentKeepalive uint16                         `binapi:"u16,name=persistent_keepalive" json:"persistent_keepalive,omitempty"`
	TableID             uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	Endpoint            ip_types.Address               `binapi:"address,name=endpoint" json:"endpoint,omitempty"`
	SwIfIndex           interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Flags               WireguardPeerFlags             `binapi:"wireguard_peer_flags,name=flags" json:"flags,omitempty"`
	NAllowedIps         uint8                          `binapi:"u8,name=n_allowed_ips" json:"-"`
	AllowedIps          []ip_types.Prefix              `binapi:"prefix[n_allowed_ips],name=allowed_ips" json:"allowed_ips,omitempty"`
}

// WantWireguardPeerEvents defines message 'want_wireguard_peer_events'.
// InProgress: the message form may change in the future versions
type WantWireguardPeerEvents struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
	PeerIndex     uint32                         `binapi:"u32,name=peer_index,default=4294967295" json:"peer_index,omitempty"`
	EnableDisable uint32                         `binapi:"u32,name=enable_disable" json:"enable_disable,omitempty"`
	PID           uint32                         `binapi:"u32,name=pid" json:"pid,omitempty"`
}

func (m *WantWireguardPeerEvents) Reset()               { *m = WantWireguardPeerEvents{} }
func (*WantWireguardPeerEvents) GetMessageName() string { return "want_wireguard_peer_events" }
func (*WantWireguardPeerEvents) GetCrcString() string   { return "3bc666c8" }
func (*WantWireguardPeerEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantWireguardPeerEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 4 // m.PeerIndex
	size += 4 // m.EnableDisable
	size += 4 // m.PID
	return size
}
func (m *WantWireguardPeerEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(m.PeerIndex)
	buf.EncodeUint32(m.EnableDisable)
	buf.EncodeUint32(m.PID)
	return buf.Bytes(), nil
}
func (m *WantWireguardPeerEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.PeerIndex = buf.DecodeUint32()
	m.EnableDisable = buf.DecodeUint32()
	m.PID = buf.DecodeUint32()
	return nil
}

// WantWireguardPeerEventsReply defines message 'want_wireguard_peer_events_reply'.
// InProgress: the message form may change in the future versions
type WantWireguardPeerEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantWireguardPeerEventsReply) Reset() { *m = WantWireguardPeerEventsReply{} }
func (*WantWireguardPeerEventsReply) GetMessageName() string {
	return "want_wireguard_peer_events_reply"
}
func (*WantWireguardPeerEventsReply) GetCrcString() string { return "e8d4e804" }
func (*WantWireguardPeerEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantWireguardPeerEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantWireguardPeerEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantWireguardPeerEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WgSetAsyncMode defines message 'wg_set_async_mode'.
// InProgress: the message form may change in the future versions
type WgSetAsyncMode struct {
	AsyncEnable bool `binapi:"bool,name=async_enable,default=false" json:"async_enable,omitempty"`
}

func (m *WgSetAsyncMode) Reset()               { *m = WgSetAsyncMode{} }
func (*WgSetAsyncMode) GetMessageName() string { return "wg_set_async_mode" }
func (*WgSetAsyncMode) GetCrcString() string   { return "a6465f7c" }
func (*WgSetAsyncMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WgSetAsyncMode) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.AsyncEnable
	return size
}
func (m *WgSetAsyncMode) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.AsyncEnable)
	return buf.Bytes(), nil
}
func (m *WgSetAsyncMode) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.AsyncEnable = buf.DecodeBool()
	return nil
}

// WgSetAsyncModeReply defines message 'wg_set_async_mode_reply'.
// InProgress: the message form may change in the future versions
type WgSetAsyncModeReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WgSetAsyncModeReply) Reset()               { *m = WgSetAsyncModeReply{} }
func (*WgSetAsyncModeReply) GetMessageName() string { return "wg_set_async_mode_reply" }
func (*WgSetAsyncModeReply) GetCrcString() string   { return "e8d4e804" }
func (*WgSetAsyncModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WgSetAsyncModeReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WgSetAsyncModeReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WgSetAsyncModeReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WireguardInterfaceCreate defines message 'wireguard_interface_create'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceCreate struct {
	Interface   WireguardInterface `binapi:"wireguard_interface,name=interface" json:"interface,omitempty"`
	GenerateKey bool               `binapi:"bool,name=generate_key" json:"generate_key,omitempty"`
}

func (m *WireguardInterfaceCreate) Reset()               { *m = WireguardInterfaceCreate{} }
func (*WireguardInterfaceCreate) GetMessageName() string { return "wireguard_interface_create" }
func (*WireguardInterfaceCreate) GetCrcString() string   { return "a530137e" }
func (*WireguardInterfaceCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardInterfaceCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Interface.UserInstance
	size += 4      // m.Interface.SwIfIndex
	size += 1 * 32 // m.Interface.PrivateKey
	size += 1 * 32 // m.Interface.PublicKey
	size += 2      // m.Interface.Port
	size += 1      // m.Interface.SrcIP.Af
	size += 1 * 16 // m.Interface.SrcIP.Un
	size += 1      // m.GenerateKey
	return size
}
func (m *WireguardInterfaceCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Interface.UserInstance)
	buf.EncodeUint32(uint32(m.Interface.SwIfIndex))
	buf.EncodeBytes(m.Interface.PrivateKey, 32)
	buf.EncodeBytes(m.Interface.PublicKey, 32)
	buf.EncodeUint16(m.Interface.Port)
	buf.EncodeUint8(uint8(m.Interface.SrcIP.Af))
	buf.EncodeBytes(m.Interface.SrcIP.Un.XXX_UnionData[:], 16)
	buf.EncodeBool(m.GenerateKey)
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Interface.UserInstance = buf.DecodeUint32()
	m.Interface.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Interface.PrivateKey = make([]byte, 32)
	copy(m.Interface.PrivateKey, buf.DecodeBytes(len(m.Interface.PrivateKey)))
	m.Interface.PublicKey = make([]byte, 32)
	copy(m.Interface.PublicKey, buf.DecodeBytes(len(m.Interface.PublicKey)))
	m.Interface.Port = buf.DecodeUint16()
	m.Interface.SrcIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.GenerateKey = buf.DecodeBool()
	return nil
}

// WireguardInterfaceCreateReply defines message 'wireguard_interface_create_reply'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *WireguardInterfaceCreateReply) Reset() { *m = WireguardInterfaceCreateReply{} }
func (*WireguardInterfaceCreateReply) GetMessageName() string {
	return "wireguard_interface_create_reply"
}
func (*WireguardInterfaceCreateReply) GetCrcString() string { return "5383d31f" }
func (*WireguardInterfaceCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardInterfaceCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *WireguardInterfaceCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// WireguardInterfaceDelete defines message 'wireguard_interface_delete'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *WireguardInterfaceDelete) Reset()               { *m = WireguardInterfaceDelete{} }
func (*WireguardInterfaceDelete) GetMessageName() string { return "wireguard_interface_delete" }
func (*WireguardInterfaceDelete) GetCrcString() string   { return "f9e6675e" }
func (*WireguardInterfaceDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardInterfaceDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *WireguardInterfaceDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// WireguardInterfaceDeleteReply defines message 'wireguard_interface_delete_reply'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WireguardInterfaceDeleteReply) Reset() { *m = WireguardInterfaceDeleteReply{} }
func (*WireguardInterfaceDeleteReply) GetMessageName() string {
	return "wireguard_interface_delete_reply"
}
func (*WireguardInterfaceDeleteReply) GetCrcString() string { return "e8d4e804" }
func (*WireguardInterfaceDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardInterfaceDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WireguardInterfaceDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WireguardInterfaceDetails defines message 'wireguard_interface_details'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceDetails struct {
	Interface WireguardInterface `binapi:"wireguard_interface,name=interface" json:"interface,omitempty"`
}

func (m *WireguardInterfaceDetails) Reset()               { *m = WireguardInterfaceDetails{} }
func (*WireguardInterfaceDetails) GetMessageName() string { return "wireguard_interface_details" }
func (*WireguardInterfaceDetails) GetCrcString() string   { return "0dd4865d" }
func (*WireguardInterfaceDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardInterfaceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Interface.UserInstance
	size += 4      // m.Interface.SwIfIndex
	size += 1 * 32 // m.Interface.PrivateKey
	size += 1 * 32 // m.Interface.PublicKey
	size += 2      // m.Interface.Port
	size += 1      // m.Interface.SrcIP.Af
	size += 1 * 16 // m.Interface.SrcIP.Un
	return size
}
func (m *WireguardInterfaceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Interface.UserInstance)
	buf.EncodeUint32(uint32(m.Interface.SwIfIndex))
	buf.EncodeBytes(m.Interface.PrivateKey, 32)
	buf.EncodeBytes(m.Interface.PublicKey, 32)
	buf.EncodeUint16(m.Interface.Port)
	buf.EncodeUint8(uint8(m.Interface.SrcIP.Af))
	buf.EncodeBytes(m.Interface.SrcIP.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Interface.UserInstance = buf.DecodeUint32()
	m.Interface.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Interface.PrivateKey = make([]byte, 32)
	copy(m.Interface.PrivateKey, buf.DecodeBytes(len(m.Interface.PrivateKey)))
	m.Interface.PublicKey = make([]byte, 32)
	copy(m.Interface.PublicKey, buf.DecodeBytes(len(m.Interface.PublicKey)))
	m.Interface.Port = buf.DecodeUint16()
	m.Interface.SrcIP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Interface.SrcIP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// WireguardInterfaceDump defines message 'wireguard_interface_dump'.
// InProgress: the message form may change in the future versions
type WireguardInterfaceDump struct {
	ShowPrivateKey bool                           `binapi:"bool,name=show_private_key" json:"show_private_key,omitempty"`
	SwIfIndex      interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *WireguardInterfaceDump) Reset()               { *m = WireguardInterfaceDump{} }
func (*WireguardInterfaceDump) GetMessageName() string { return "wireguard_interface_dump" }
func (*WireguardInterfaceDump) GetCrcString() string   { return "2c954158" }
func (*WireguardInterfaceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardInterfaceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.ShowPrivateKey
	size += 4 // m.SwIfIndex
	return size
}
func (m *WireguardInterfaceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.ShowPrivateKey)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *WireguardInterfaceDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ShowPrivateKey = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// WireguardPeerAdd defines message 'wireguard_peer_add'.
// InProgress: the message form may change in the future versions
type WireguardPeerAdd struct {
	Peer WireguardPeer `binapi:"wireguard_peer,name=peer" json:"peer,omitempty"`
}

func (m *WireguardPeerAdd) Reset()               { *m = WireguardPeerAdd{} }
func (*WireguardPeerAdd) GetMessageName() string { return "wireguard_peer_add" }
func (*WireguardPeerAdd) GetCrcString() string   { return "9b8aad61" }
func (*WireguardPeerAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardPeerAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Peer.PeerIndex
	size += 1 * 32 // m.Peer.PublicKey
	size += 2      // m.Peer.Port
	size += 2      // m.Peer.PersistentKeepalive
	size += 4      // m.Peer.TableID
	size += 1      // m.Peer.Endpoint.Af
	size += 1 * 16 // m.Peer.Endpoint.Un
	size += 4      // m.Peer.SwIfIndex
	size += 1      // m.Peer.Flags
	size += 1      // m.Peer.NAllowedIps
	for j2 := 0; j2 < len(m.Peer.AllowedIps); j2++ {
		var s2 ip_types.Prefix
		_ = s2
		if j2 < len(m.Peer.AllowedIps) {
			s2 = m.Peer.AllowedIps[j2]
		}
		size += 1      // s2.Address.Af
		size += 1 * 16 // s2.Address.Un
		size += 1      // s2.Len
	}
	return size
}
func (m *WireguardPeerAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Peer.PeerIndex)
	buf.EncodeBytes(m.Peer.PublicKey, 32)
	buf.EncodeUint16(m.Peer.Port)
	buf.EncodeUint16(m.Peer.PersistentKeepalive)
	buf.EncodeUint32(m.Peer.TableID)
	buf.EncodeUint8(uint8(m.Peer.Endpoint.Af))
	buf.EncodeBytes(m.Peer.Endpoint.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Peer.SwIfIndex))
	buf.EncodeUint8(uint8(m.Peer.Flags))
	buf.EncodeUint8(uint8(len(m.Peer.AllowedIps)))
	for j1 := 0; j1 < len(m.Peer.AllowedIps); j1++ {
		var v1 ip_types.Prefix // AllowedIps
		if j1 < len(m.Peer.AllowedIps) {
			v1 = m.Peer.AllowedIps[j1]
		}
		buf.EncodeUint8(uint8(v1.Address.Af))
		buf.EncodeBytes(v1.Address.Un.XXX_UnionData[:], 16)
		buf.EncodeUint8(v1.Len)
	}
	return buf.Bytes(), nil
}
func (m *WireguardPeerAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Peer.PeerIndex = buf.DecodeUint32()
	m.Peer.PublicKey = make([]byte, 32)
	copy(m.Peer.PublicKey, buf.DecodeBytes(len(m.Peer.PublicKey)))
	m.Peer.Port = buf.DecodeUint16()
	m.Peer.PersistentKeepalive = buf.DecodeUint16()
	m.Peer.TableID = buf.DecodeUint32()
	m.Peer.Endpoint.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Peer.Endpoint.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Peer.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Peer.Flags = WireguardPeerFlags(buf.DecodeUint8())
	m.Peer.NAllowedIps = buf.DecodeUint8()
	m.Peer.AllowedIps = make([]ip_types.Prefix, m.Peer.NAllowedIps)
	for j1 := 0; j1 < len(m.Peer.AllowedIps); j1++ {
		m.Peer.AllowedIps[j1].Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Peer.AllowedIps[j1].Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Peer.AllowedIps[j1].Len = buf.DecodeUint8()
	}
	return nil
}

// WireguardPeerAddReply defines message 'wireguard_peer_add_reply'.
// InProgress: the message form may change in the future versions
type WireguardPeerAddReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	PeerIndex uint32 `binapi:"u32,name=peer_index" json:"peer_index,omitempty"`
}

func (m *WireguardPeerAddReply) Reset()               { *m = WireguardPeerAddReply{} }
func (*WireguardPeerAddReply) GetMessageName() string { return "wireguard_peer_add_reply" }
func (*WireguardPeerAddReply) GetCrcString() string   { return "084a0cd3" }
func (*WireguardPeerAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardPeerAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.PeerIndex
	return size
}
func (m *WireguardPeerAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.PeerIndex)
	return buf.Bytes(), nil
}
func (m *WireguardPeerAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.PeerIndex = buf.DecodeUint32()
	return nil
}

// WireguardPeerEvent defines message 'wireguard_peer_event'.
// InProgress: the message form may change in the future versions
type WireguardPeerEvent struct {
	PID       uint32             `binapi:"u32,name=pid" json:"pid,omitempty"`
	PeerIndex uint32             `binapi:"u32,name=peer_index" json:"peer_index,omitempty"`
	Flags     WireguardPeerFlags `binapi:"wireguard_peer_flags,name=flags" json:"flags,omitempty"`
}

func (m *WireguardPeerEvent) Reset()               { *m = WireguardPeerEvent{} }
func (*WireguardPeerEvent) GetMessageName() string { return "wireguard_peer_event" }
func (*WireguardPeerEvent) GetCrcString() string   { return "4e1b5d67" }
func (*WireguardPeerEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *WireguardPeerEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PID
	size += 4 // m.PeerIndex
	size += 1 // m.Flags
	return size
}
func (m *WireguardPeerEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint32(m.PeerIndex)
	buf.EncodeUint8(uint8(m.Flags))
	return buf.Bytes(), nil
}
func (m *WireguardPeerEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PID = buf.DecodeUint32()
	m.PeerIndex = buf.DecodeUint32()
	m.Flags = WireguardPeerFlags(buf.DecodeUint8())
	return nil
}

// WireguardPeerRemove defines message 'wireguard_peer_remove'.
// InProgress: the message form may change in the future versions
type WireguardPeerRemove struct {
	PeerIndex uint32 `binapi:"u32,name=peer_index" json:"peer_index,omitempty"`
}

func (m *WireguardPeerRemove) Reset()               { *m = WireguardPeerRemove{} }
func (*WireguardPeerRemove) GetMessageName() string { return "wireguard_peer_remove" }
func (*WireguardPeerRemove) GetCrcString() string   { return "3b74607a" }
func (*WireguardPeerRemove) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardPeerRemove) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PeerIndex
	return size
}
func (m *WireguardPeerRemove) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PeerIndex)
	return buf.Bytes(), nil
}
func (m *WireguardPeerRemove) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PeerIndex = buf.DecodeUint32()
	return nil
}

// WireguardPeerRemoveReply defines message 'wireguard_peer_remove_reply'.
// InProgress: the message form may change in the future versions
type WireguardPeerRemoveReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WireguardPeerRemoveReply) Reset()               { *m = WireguardPeerRemoveReply{} }
func (*WireguardPeerRemoveReply) GetMessageName() string { return "wireguard_peer_remove_reply" }
func (*WireguardPeerRemoveReply) GetCrcString() string   { return "e8d4e804" }
func (*WireguardPeerRemoveReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardPeerRemoveReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WireguardPeerRemoveReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WireguardPeerRemoveReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WireguardPeersDetails defines message 'wireguard_peers_details'.
// InProgress: the message form may change in the future versions
type WireguardPeersDetails struct {
	Peer WireguardPeer `binapi:"wireguard_peer,name=peer" json:"peer,omitempty"`
}

func (m *WireguardPeersDetails) Reset()               { *m = WireguardPeersDetails{} }
func (*WireguardPeersDetails) GetMessageName() string { return "wireguard_peers_details" }
func (*WireguardPeersDetails) GetCrcString() string   { return "6a9f6bc3" }
func (*WireguardPeersDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WireguardPeersDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Peer.PeerIndex
	size += 1 * 32 // m.Peer.PublicKey
	size += 2      // m.Peer.Port
	size += 2      // m.Peer.PersistentKeepalive
	size += 4      // m.Peer.TableID
	size += 1      // m.Peer.Endpoint.Af
	size += 1 * 16 // m.Peer.Endpoint.Un
	size += 4      // m.Peer.SwIfIndex
	size += 1      // m.Peer.Flags
	size += 1      // m.Peer.NAllowedIps
	for j2 := 0; j2 < len(m.Peer.AllowedIps); j2++ {
		var s2 ip_types.Prefix
		_ = s2
		if j2 < len(m.Peer.AllowedIps) {
			s2 = m.Peer.AllowedIps[j2]
		}
		size += 1      // s2.Address.Af
		size += 1 * 16 // s2.Address.Un
		size += 1      // s2.Len
	}
	return size
}
func (m *WireguardPeersDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Peer.PeerIndex)
	buf.EncodeBytes(m.Peer.PublicKey, 32)
	buf.EncodeUint16(m.Peer.Port)
	buf.EncodeUint16(m.Peer.PersistentKeepalive)
	buf.EncodeUint32(m.Peer.TableID)
	buf.EncodeUint8(uint8(m.Peer.Endpoint.Af))
	buf.EncodeBytes(m.Peer.Endpoint.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.Peer.SwIfIndex))
	buf.EncodeUint8(uint8(m.Peer.Flags))
	buf.EncodeUint8(uint8(len(m.Peer.AllowedIps)))
	for j1 := 0; j1 < len(m.Peer.AllowedIps); j1++ {
		var v1 ip_types.Prefix // AllowedIps
		if j1 < len(m.Peer.AllowedIps) {
			v1 = m.Peer.AllowedIps[j1]
		}
		buf.EncodeUint8(uint8(v1.Address.Af))
		buf.EncodeBytes(v1.Address.Un.XXX_UnionData[:], 16)
		buf.EncodeUint8(v1.Len)
	}
	return buf.Bytes(), nil
}
func (m *WireguardPeersDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Peer.PeerIndex = buf.DecodeUint32()
	m.Peer.PublicKey = make([]byte, 32)
	copy(m.Peer.PublicKey, buf.DecodeBytes(len(m.Peer.PublicKey)))
	m.Peer.Port = buf.DecodeUint16()
	m.Peer.PersistentKeepalive = buf.DecodeUint16()
	m.Peer.TableID = buf.DecodeUint32()
	m.Peer.Endpoint.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Peer.Endpoint.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Peer.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Peer.Flags = WireguardPeerFlags(buf.DecodeUint8())
	m.Peer.NAllowedIps = buf.DecodeUint8()
	m.Peer.AllowedIps = make([]ip_types.Prefix, m.Peer.NAllowedIps)
	for j1 := 0; j1 < len(m.Peer.AllowedIps); j1++ {
		m.Peer.AllowedIps[j1].Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Peer.AllowedIps[j1].Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
		m.Peer.AllowedIps[j1].Len = buf.DecodeUint8()
	}
	return nil
}

// WireguardPeersDump defines message 'wireguard_peers_dump'.
// InProgress: the message form may change in the future versions
type WireguardPeersDump struct {
	PeerIndex uint32 `binapi:"u32,name=peer_index,default=4294967295" json:"peer_index,omitempty"`
}

func (m *WireguardPeersDump) Reset()               { *m = WireguardPeersDump{} }
func (*WireguardPeersDump) GetMessageName() string { return "wireguard_peers_dump" }
func (*WireguardPeersDump) GetCrcString() string   { return "3b74607a" }
func (*WireguardPeersDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WireguardPeersDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.PeerIndex
	return size
}
func (m *WireguardPeersDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PeerIndex)
	return buf.Bytes(), nil
}
func (m *WireguardPeersDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PeerIndex = buf.DecodeUint32()
	return nil
}

func init() { file_wireguard_binapi_init() }
func file_wireguard_binapi_init() {
	api.RegisterMessage((*WantWireguardPeerEvents)(nil), "want_wireguard_peer_events_3bc666c8")
	api.RegisterMessage((*WantWireguardPeerEventsReply)(nil), "want_wireguard_peer_events_reply_e8d4e804")
	api.RegisterMessage((*WgSetAsyncMode)(nil), "wg_set_async_mode_a6465f7c")
	api.RegisterMessage((*WgSetAsyncModeReply)(nil), "wg_set_async_mode_reply_e8d4e804")
	api.RegisterMessage((*WireguardInterfaceCreate)(nil), "wireguard_interface_create_a530137e")
	api.RegisterMessage((*WireguardInterfaceCreateReply)(nil), "wireguard_interface_create_reply_5383d31f")
	api.RegisterMessage((*WireguardInterfaceDelete)(nil), "wireguard_interface_delete_f9e6675e")
	api.RegisterMessage((*WireguardInterfaceDeleteReply)(nil), "wireguard_interface_delete_reply_e8d4e804")
	api.RegisterMessage((*WireguardInterfaceDetails)(nil), "wireguard_interface_details_0dd4865d")
	api.RegisterMessage((*WireguardInterfaceDump)(nil), "wireguard_interface_dump_2c954158")
	api.RegisterMessage((*WireguardPeerAdd)(nil), "wireguard_peer_add_9b8aad61")
	api.RegisterMessage((*WireguardPeerAddReply)(nil), "wireguard_peer_add_reply_084a0cd3")
	api.RegisterMessage((*WireguardPeerEvent)(nil), "wireguard_peer_event_4e1b5d67")
	api.RegisterMessage((*WireguardPeerRemove)(nil), "wireguard_peer_remove_3b74607a")
	api.RegisterMessage((*WireguardPeerRemoveReply)(nil), "wireguard_peer_remove_reply_e8d4e804")
	api.RegisterMessage((*WireguardPeersDetails)(nil), "wireguard_peers_details_6a9f6bc3")
	api.RegisterMessage((*WireguardPeersDump)(nil), "wireguard_peers_dump_3b74607a")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*WantWireguardPeerEvents)(nil),
		(*WantWireguardPeerEventsReply)(nil),
		(*WgSetAsyncMode)(nil),
		(*WgSetAsyncModeReply)(nil),
		(*WireguardInterfaceCreate)(nil),
		(*WireguardInterfaceCreateReply)(nil),
		(*WireguardInterfaceDelete)(nil),
		(*WireguardInterfaceDeleteReply)(nil),
		(*WireguardInterfaceDetails)(nil),
		(*WireguardInterfaceDump)(nil),
		(*WireguardPeerAdd)(nil),
		(*WireguardPeerAddReply)(nil),
		(*WireguardPeerEvent)(nil),
		(*WireguardPeerRemove)(nil),
		(*WireguardPeerRemoveReply)(nil),
		(*WireguardPeersDetails)(nil),
		(*WireguardPeersDump)(nil),
	}
}
