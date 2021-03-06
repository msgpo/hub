package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterNode)(nil)
	_ sdk.Msg = (*MsgUpdateNode)(nil)
	_ sdk.Msg = (*MsgSetNodeStatus)(nil)
)

// MsgRegisterNode is for registering a VPN node.
type MsgRegisterNode struct {
	From          sdk.AccAddress  `json:"from"`
	Provider      hub.ProvAddress `json:"provider,omitempty"`
	Price         sdk.Coins       `json:"price,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      NodeCategory    `json:"category"`
}

func NewMsgRegisterNode(from sdk.AccAddress, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category NodeCategory) MsgRegisterNode {
	return MsgRegisterNode{
		From:          from,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
	}
}

func (m MsgRegisterNode) Route() string {
	return RouterKey
}

func (m MsgRegisterNode) Type() string {
	return "register_node"
}

func (m MsgRegisterNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Either provider or price should be nil
	if (m.Provider != nil && m.Price != nil) ||
		(m.Provider == nil && m.Price == nil) {
		return ErrorInvalidField("provider and price")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return ErrorInvalidField("provider")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidField("price")
	}

	// InternetSpeed shouldn't be negative and zero
	if !m.InternetSpeed.IsValid() {
		return ErrorInvalidField("internet_speed")
	}

	// RemoteURL can't be empty and length should be (0, 64]
	if len(m.RemoteURL) == 0 || len(m.RemoteURL) > 64 {
		return ErrorInvalidField("remote_url")
	}

	// Version can't be empty and length should be (0, 64]
	if len(m.Version) == 0 || len(m.Version) > 64 {
		return ErrorInvalidField("version")
	}

	// Category should be valid
	if !m.Category.IsValid() {
		return ErrorInvalidField("category")
	}

	return nil
}

func (m MsgRegisterNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// MsgUpdateNode is for updating the information of a VPN node.
type MsgUpdateNode struct {
	From          hub.NodeAddress `json:"from"`
	Provider      hub.ProvAddress `json:"provider,omitempty"`
	Price         sdk.Coins       `json:"price,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed,omitempty"`
	RemoteURL     string          `json:"remote_url,omitempty"`
	Version       string          `json:"version,omitempty"`
	Category      NodeCategory    `json:"category,omitempty"`
}

func NewMsgUpdateNode(from hub.NodeAddress, provider hub.ProvAddress, price sdk.Coins,
	speed hub.Bandwidth, remoteURL, version string, category NodeCategory) MsgUpdateNode {
	return MsgUpdateNode{
		From:          from,
		Provider:      provider,
		Price:         price,
		InternetSpeed: speed,
		RemoteURL:     remoteURL,
		Version:       version,
		Category:      category,
	}
}

func (m MsgUpdateNode) Route() string {
	return RouterKey
}

func (m MsgUpdateNode) Type() string {
	return "update_node"
}

func (m MsgUpdateNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Provider and Price both shouldn't nil at the same time
	if m.Provider != nil && m.Price != nil {
		return ErrorInvalidField("provider and price")
	}

	// Provider can be nil. If not, it shouldn't be empty
	if m.Provider != nil && m.Provider.Empty() {
		return ErrorInvalidField("provider")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidField("price")
	}

	// InternetSpeed can be zero. If not, it shouldn't be negative and zero
	if !m.InternetSpeed.IsAllZero() && !m.InternetSpeed.IsValid() {
		return ErrorInvalidField("internet_speed")
	}

	// RemoteURL length should be [0, 64]
	if len(m.RemoteURL) > 64 {
		return ErrorInvalidField("remote_url")
	}

	// Version length should be [0, 64]
	if len(m.Version) > 64 {
		return ErrorInvalidField("version")
	}

	// Category can be Unknown. If not, should be valid
	if m.Category != CategoryUnknown && !m.Category.IsValid() {
		return ErrorInvalidField("category")
	}

	return nil
}

func (m MsgUpdateNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}

// MsgSetNodeStatus is for updating the status of a VPN node.
type MsgSetNodeStatus struct {
	From   hub.NodeAddress `json:"from"`
	Status hub.Status      `json:"status"`
}

func NewMsgSetNodeStatus(from hub.NodeAddress, status hub.Status) MsgSetNodeStatus {
	return MsgSetNodeStatus{
		From:   from,
		Status: status,
	}
}

func (m MsgSetNodeStatus) Route() string {
	return RouterKey
}

func (m MsgSetNodeStatus) Type() string {
	return "set_node_status"
}

func (m MsgSetNodeStatus) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}

	// Status should be valid
	if !m.Status.IsValid() {
		return ErrorInvalidField("status")
	}

	return nil
}

func (m MsgSetNodeStatus) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSetNodeStatus) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From.Bytes()}
}
