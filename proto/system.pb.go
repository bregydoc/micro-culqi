// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/system.proto

package pculqi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
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

type AvailableCurrency int32

const (
	AvailableCurrency_PEN AvailableCurrency = 0
	AvailableCurrency_USD AvailableCurrency = 1
)

var AvailableCurrency_name = map[int32]string{
	0: "PEN",
	1: "USD",
}

var AvailableCurrency_value = map[string]int32{
	"PEN": 0,
	"USD": 1,
}

func (x AvailableCurrency) String() string {
	return proto.EnumName(AvailableCurrency_name, int32(x))
}

func (AvailableCurrency) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{0}
}

type IsOk struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsOk) Reset()         { *m = IsOk{} }
func (m *IsOk) String() string { return proto.CompactTextString(m) }
func (*IsOk) ProtoMessage()    {}
func (*IsOk) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{0}
}

func (m *IsOk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsOk.Unmarshal(m, b)
}
func (m *IsOk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsOk.Marshal(b, m, deterministic)
}
func (m *IsOk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsOk.Merge(m, src)
}
func (m *IsOk) XXX_Size() int {
	return xxx_messageInfo_IsOk.Size(m)
}
func (m *IsOk) XXX_DiscardUnknown() {
	xxx_messageInfo_IsOk.DiscardUnknown(m)
}

var xxx_messageInfo_IsOk proto.InternalMessageInfo

func (m *IsOk) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type TemplateData struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateData) Reset()         { *m = TemplateData{} }
func (m *TemplateData) String() string { return proto.CompactTextString(m) }
func (*TemplateData) ProtoMessage()    {}
func (*TemplateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{1}
}

func (m *TemplateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateData.Unmarshal(m, b)
}
func (m *TemplateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateData.Marshal(b, m, deterministic)
}
func (m *TemplateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateData.Merge(m, src)
}
func (m *TemplateData) XXX_Size() int {
	return xxx_messageInfo_TemplateData.Size(m)
}
func (m *TemplateData) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateData.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateData proto.InternalMessageInfo

func (m *TemplateData) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type InvoiceID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceID) Reset()         { *m = InvoiceID{} }
func (m *InvoiceID) String() string { return proto.CompactTextString(m) }
func (*InvoiceID) ProtoMessage()    {}
func (*InvoiceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{2}
}

func (m *InvoiceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceID.Unmarshal(m, b)
}
func (m *InvoiceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceID.Marshal(b, m, deterministic)
}
func (m *InvoiceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceID.Merge(m, src)
}
func (m *InvoiceID) XXX_Size() int {
	return xxx_messageInfo_InvoiceID.Size(m)
}
func (m *InvoiceID) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceID.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceID proto.InternalMessageInfo

func (m *InvoiceID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MinimalInvoice struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Email                string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Products             []*Product        `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	Currency             AvailableCurrency `protobuf:"varint,4,opt,name=currency,proto3,enum=culqi.AvailableCurrency" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MinimalInvoice) Reset()         { *m = MinimalInvoice{} }
func (m *MinimalInvoice) String() string { return proto.CompactTextString(m) }
func (*MinimalInvoice) ProtoMessage()    {}
func (*MinimalInvoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{3}
}

func (m *MinimalInvoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MinimalInvoice.Unmarshal(m, b)
}
func (m *MinimalInvoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MinimalInvoice.Marshal(b, m, deterministic)
}
func (m *MinimalInvoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinimalInvoice.Merge(m, src)
}
func (m *MinimalInvoice) XXX_Size() int {
	return xxx_messageInfo_MinimalInvoice.Size(m)
}
func (m *MinimalInvoice) XXX_DiscardUnknown() {
	xxx_messageInfo_MinimalInvoice.DiscardUnknown(m)
}

var xxx_messageInfo_MinimalInvoice proto.InternalMessageInfo

func (m *MinimalInvoice) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MinimalInvoice) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *MinimalInvoice) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *MinimalInvoice) GetCurrency() AvailableCurrency {
	if m != nil {
		return m.Currency
	}
	return AvailableCurrency_PEN
}

type Product struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price                float64           `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Currency             AvailableCurrency `protobuf:"varint,4,opt,name=currency,proto3,enum=culqi.AvailableCurrency" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{4}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Product) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetCurrency() AvailableCurrency {
	if m != nil {
		return m.Currency
	}
	return AvailableCurrency_PEN
}

type Currency struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Multiplier           int32    `protobuf:"varint,4,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{5}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Currency) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Currency) GetMultiplier() int32 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type Order struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            string              `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Token                string              `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Products             []*Product          `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	Info                 *PersonInfo         `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Currency             *Currency           `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	Card                 string              `protobuf:"bytes,7,opt,name=card,proto3" json:"card,omitempty"`
	Discount             float64             `protobuf:"fixed64,8,opt,name=discount,proto3" json:"discount,omitempty"`
	Metadata             map[string]*any.Any `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{6}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Order) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Order) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Order) GetInfo() *PersonInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Order) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *Order) GetCard() string {
	if m != nil {
		return m.Card
	}
	return ""
}

func (m *Order) GetDiscount() float64 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *Order) GetMetadata() map[string]*any.Any {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PersonInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	IdNumber             string   `protobuf:"bytes,3,opt,name=idNumber,proto3" json:"idNumber,omitempty"`
	Phone                int64    `protobuf:"varint,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	AddressCity          string   `protobuf:"bytes,6,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	CountryCode          string   `protobuf:"bytes,7,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonInfo) Reset()         { *m = PersonInfo{} }
func (m *PersonInfo) String() string { return proto.CompactTextString(m) }
func (*PersonInfo) ProtoMessage()    {}
func (*PersonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{7}
}

func (m *PersonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonInfo.Unmarshal(m, b)
}
func (m *PersonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonInfo.Marshal(b, m, deterministic)
}
func (m *PersonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonInfo.Merge(m, src)
}
func (m *PersonInfo) XXX_Size() int {
	return xxx_messageInfo_PersonInfo.Size(m)
}
func (m *PersonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PersonInfo proto.InternalMessageInfo

func (m *PersonInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *PersonInfo) GetIdNumber() string {
	if m != nil {
		return m.IdNumber
	}
	return ""
}

func (m *PersonInfo) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *PersonInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PersonInfo) GetAddressCity() string {
	if m != nil {
		return m.AddressCity
	}
	return ""
}

func (m *PersonInfo) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

type Invoice struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Order                *Order       `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	Email                string       `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Charged              bool         `protobuf:"varint,4,opt,name=charged,proto3" json:"charged,omitempty"`
	Company              *CompanyInfo `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	ChargedAt            string       `protobuf:"bytes,6,opt,name=chargedAt,proto3" json:"chargedAt,omitempty"`
	ExternalId           string       `protobuf:"bytes,7,opt,name=externalId,proto3" json:"externalId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{8}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Invoice) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *Invoice) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Invoice) GetCharged() bool {
	if m != nil {
		return m.Charged
	}
	return false
}

func (m *Invoice) GetCompany() *CompanyInfo {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *Invoice) GetChargedAt() string {
	if m != nil {
		return m.ChargedAt
	}
	return ""
}

func (m *Invoice) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

type CompanyInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PhoneSupport         string   `protobuf:"bytes,2,opt,name=phoneSupport,proto3" json:"phoneSupport,omitempty"`
	EmailSupport         string   `protobuf:"bytes,3,opt,name=emailSupport,proto3" json:"emailSupport,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompanyInfo) Reset()         { *m = CompanyInfo{} }
func (m *CompanyInfo) String() string { return proto.CompactTextString(m) }
func (*CompanyInfo) ProtoMessage()    {}
func (*CompanyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b14b94b6dd6149c, []int{9}
}

func (m *CompanyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompanyInfo.Unmarshal(m, b)
}
func (m *CompanyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompanyInfo.Marshal(b, m, deterministic)
}
func (m *CompanyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompanyInfo.Merge(m, src)
}
func (m *CompanyInfo) XXX_Size() int {
	return xxx_messageInfo_CompanyInfo.Size(m)
}
func (m *CompanyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CompanyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CompanyInfo proto.InternalMessageInfo

func (m *CompanyInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CompanyInfo) GetPhoneSupport() string {
	if m != nil {
		return m.PhoneSupport
	}
	return ""
}

func (m *CompanyInfo) GetEmailSupport() string {
	if m != nil {
		return m.EmailSupport
	}
	return ""
}

func init() {
	proto.RegisterEnum("culqi.AvailableCurrency", AvailableCurrency_name, AvailableCurrency_value)
	proto.RegisterType((*IsOk)(nil), "culqi.IsOk")
	proto.RegisterType((*TemplateData)(nil), "culqi.TemplateData")
	proto.RegisterType((*InvoiceID)(nil), "culqi.InvoiceID")
	proto.RegisterType((*MinimalInvoice)(nil), "culqi.MinimalInvoice")
	proto.RegisterType((*Product)(nil), "culqi.Product")
	proto.RegisterType((*Currency)(nil), "culqi.Currency")
	proto.RegisterType((*Order)(nil), "culqi.Order")
	proto.RegisterMapType((map[string]*any.Any)(nil), "culqi.Order.MetadataEntry")
	proto.RegisterType((*PersonInfo)(nil), "culqi.PersonInfo")
	proto.RegisterType((*Invoice)(nil), "culqi.Invoice")
	proto.RegisterType((*CompanyInfo)(nil), "culqi.CompanyInfo")
}

func init() { proto.RegisterFile("proto/system.proto", fileDescriptor_8b14b94b6dd6149c) }

var fileDescriptor_8b14b94b6dd6149c = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x71, 0x7e, 0x9c, 0x93, 0x6e, 0xe8, 0xce, 0x2e, 0x95, 0x09, 0x08, 0x45, 0x96, 0x56,
	0x8a, 0x0a, 0x4a, 0xa5, 0x80, 0x0a, 0xe2, 0xae, 0x9b, 0x54, 0x28, 0x17, 0xbb, 0x5d, 0xbc, 0x54,
	0x48, 0xdc, 0x4d, 0x3c, 0xd3, 0x76, 0x88, 0x3d, 0x63, 0xc6, 0xe3, 0x6a, 0xfd, 0x02, 0xdc, 0xf0,
	0x0e, 0x3c, 0x0e, 0xef, 0xc0, 0x03, 0xf0, 0x1e, 0x68, 0xfe, 0x12, 0xa7, 0xad, 0x84, 0xc4, 0xdd,
	0x9c, 0xef, 0x7c, 0x33, 0x39, 0xdf, 0x77, 0xce, 0x71, 0x00, 0x95, 0x52, 0x28, 0x71, 0x56, 0x35,
	0x95, 0xa2, 0xc5, 0xdc, 0x04, 0xa8, 0x97, 0xd5, 0xf9, 0x6f, 0x6c, 0xf2, 0xe9, 0xad, 0x10, 0xb7,
	0x39, 0x3d, 0x33, 0xe0, 0xa6, 0xbe, 0x39, 0xc3, 0xbc, 0xb1, 0x8c, 0xe4, 0x04, 0xba, 0xeb, 0xea,
	0x6a, 0x8b, 0xc6, 0xd0, 0x11, 0xdb, 0x38, 0x98, 0x06, 0xb3, 0x28, 0xed, 0x88, 0x6d, 0x92, 0xc0,
	0xd1, 0x4f, 0xb4, 0x28, 0x73, 0xac, 0xe8, 0x0a, 0x2b, 0x8c, 0x10, 0x74, 0x09, 0x56, 0xd8, 0x30,
	0x86, 0xa9, 0x39, 0x27, 0x9f, 0xc1, 0x70, 0xcd, 0xef, 0x05, 0xcb, 0xe8, 0x7a, 0xa5, 0x1f, 0x60,
	0xc4, 0xa5, 0x3b, 0x8c, 0x24, 0x7f, 0x06, 0x30, 0x7e, 0xc3, 0x38, 0x2b, 0x70, 0xee, 0x48, 0xe8,
	0x25, 0xf4, 0x94, 0xd8, 0x52, 0xee, 0x58, 0x36, 0xd0, 0x28, 0x2d, 0x30, 0xcb, 0xe3, 0x8e, 0x45,
	0x4d, 0x80, 0x4e, 0x21, 0x2a, 0xa5, 0x20, 0x75, 0xa6, 0xaa, 0x38, 0x9c, 0x86, 0xb3, 0xd1, 0x62,
	0x3c, 0x37, 0x62, 0xe6, 0xef, 0x2c, 0x9c, 0xee, 0xf2, 0xe8, 0x1b, 0x88, 0xb2, 0x5a, 0x4a, 0xca,
	0xb3, 0x26, 0xee, 0x4e, 0x83, 0xd9, 0x78, 0x11, 0x3b, 0xee, 0xc5, 0x3d, 0x66, 0x39, 0xde, 0xe4,
	0x74, 0xe9, 0xf2, 0xe9, 0x8e, 0x99, 0xfc, 0x11, 0xc0, 0xc0, 0xbd, 0xa5, 0xd5, 0x71, 0x5c, 0x50,
	0xaf, 0x4e, 0x9f, 0xd1, 0x14, 0x46, 0x84, 0x56, 0x99, 0x64, 0xa5, 0x62, 0x82, 0xbb, 0xea, 0xda,
	0x90, 0xae, 0xbc, 0x94, 0x2c, 0xa3, 0x71, 0x38, 0x0d, 0x66, 0x41, 0x6a, 0x83, 0xff, 0x59, 0xcd,
	0xaf, 0x10, 0x79, 0xf4, 0xc9, 0x6a, 0x4e, 0xa0, 0x5f, 0x35, 0xc5, 0x46, 0x78, 0x9b, 0x5c, 0xa4,
	0xb9, 0x99, 0x20, 0xb6, 0x84, 0x61, 0x6a, 0xce, 0xe8, 0x0b, 0x80, 0xa2, 0xce, 0x15, 0x2b, 0x73,
	0x46, 0xa5, 0xa9, 0xa1, 0x97, 0xb6, 0x90, 0xe4, 0xf7, 0x10, 0x7a, 0x57, 0x92, 0x50, 0xf9, 0xb0,
	0x69, 0xe8, 0x73, 0x18, 0x66, 0x92, 0x62, 0x45, 0xc9, 0x85, 0x72, 0x3f, 0xb4, 0x07, 0xf6, 0xfd,
	0x0b, 0xdb, 0xfd, 0x6b, 0x77, 0xaa, 0xfb, 0x1f, 0x9d, 0x7a, 0x05, 0x5d, 0xc6, 0x6f, 0x44, 0xdc,
	0x9b, 0x06, 0xb3, 0xd1, 0xe2, 0xb9, 0xe7, 0x51, 0x59, 0x09, 0xbe, 0xe6, 0x37, 0x22, 0x35, 0x69,
	0xf4, 0x65, 0xcb, 0xc2, 0xbe, 0xa1, 0x7e, 0xec, 0xa8, 0x8f, 0x9d, 0x33, 0x0e, 0x60, 0x49, 0xe2,
	0x81, 0x73, 0x00, 0x4b, 0x82, 0x26, 0x10, 0x11, 0x56, 0x65, 0xa2, 0xe6, 0x2a, 0x8e, 0x4c, 0x73,
	0x76, 0x31, 0x3a, 0x87, 0xa8, 0xa0, 0x0a, 0x9b, 0x69, 0x1e, 0x9a, 0x7a, 0x27, 0xee, 0x71, 0xe3,
	0xc9, 0xfc, 0x8d, 0x4b, 0x5e, 0x72, 0x25, 0x9b, 0x74, 0xc7, 0x9d, 0xfc, 0x08, 0xcf, 0x0e, 0x52,
	0xe8, 0x18, 0xc2, 0x2d, 0x6d, 0x9c, 0x7b, 0xfa, 0x88, 0x4e, 0xa1, 0x77, 0x8f, 0xf3, 0x9a, 0x1a,
	0xeb, 0x46, 0x8b, 0x97, 0x73, 0xbb, 0x77, 0x73, 0xbf, 0x77, 0xf3, 0x0b, 0xde, 0xa4, 0x96, 0xf2,
	0x7d, 0xe7, 0xbb, 0x20, 0xf9, 0x2b, 0x00, 0xd8, 0x8b, 0x7f, 0xb2, 0xef, 0x4f, 0x6f, 0xc7, 0x04,
	0x22, 0x46, 0xde, 0xd6, 0xc5, 0x86, 0x4a, 0xd7, 0x8c, 0x5d, 0x6c, 0xa6, 0xf2, 0x4e, 0x70, 0x6a,
	0x1a, 0x1f, 0xa6, 0x36, 0x40, 0x31, 0x0c, 0x30, 0x21, 0x92, 0x56, 0x95, 0x31, 0x7f, 0x98, 0xfa,
	0x50, 0xcf, 0xb9, 0x3b, 0x2e, 0x99, 0xb2, 0x7e, 0x0f, 0xd3, 0x36, 0xa4, 0x19, 0xc6, 0x3a, 0xd9,
	0x2c, 0xf5, 0xa8, 0x59, 0xa3, 0xdb, 0x50, 0xf2, 0x77, 0x00, 0x03, 0xbf, 0xe5, 0x0f, 0x67, 0x2a,
	0x81, 0x9e, 0xd0, 0xc6, 0x3a, 0x53, 0x8e, 0xda, 0x66, 0xa7, 0x36, 0xb5, 0x57, 0x19, 0xb6, 0x55,
	0xc6, 0x30, 0xc8, 0xee, 0xb0, 0xbc, 0xa5, 0xc4, 0x68, 0x89, 0x52, 0x1f, 0xa2, 0xaf, 0x60, 0x90,
	0x89, 0xa2, 0xc4, 0xbc, 0x71, 0xa3, 0x84, 0xfc, 0x7c, 0x58, 0xd4, 0xcc, 0x92, 0xa7, 0x98, 0xa9,
	0xb6, 0x17, 0x2f, 0x94, 0xd3, 0xb7, 0x07, 0xf4, 0xb6, 0xd0, 0x0f, 0x8a, 0x4a, 0x8e, 0xf3, 0xb5,
	0x9f, 0xa2, 0x16, 0x92, 0x30, 0x18, 0xb5, 0x5e, 0x7d, 0xb2, 0x49, 0x09, 0x1c, 0x19, 0x97, 0xdf,
	0xd7, 0x65, 0x29, 0xa4, 0xdf, 0x9c, 0x03, 0x4c, 0x73, 0x8c, 0x2a, 0xcf, 0xb1, 0x4a, 0x0f, 0xb0,
	0xd3, 0x57, 0xf0, 0xfc, 0xd1, 0x37, 0x02, 0x0d, 0x20, 0x7c, 0x77, 0xf9, 0xf6, 0xf8, 0x23, 0x7d,
	0xb8, 0x7e, 0xbf, 0x3a, 0x0e, 0x16, 0xff, 0x04, 0xd0, 0xbf, 0x5e, 0x6a, 0xbd, 0xe8, 0x1c, 0x9e,
	0x5d, 0x7e, 0xa0, 0x59, 0xad, 0xe8, 0xd2, 0x08, 0x42, 0x9f, 0x38, 0x23, 0x0e, 0x3f, 0xbd, 0x13,
	0xbf, 0x92, 0xbe, 0x49, 0xe7, 0x70, 0x72, 0x70, 0xef, 0x67, 0xa6, 0xee, 0xec, 0x27, 0xe1, 0xa0,
	0x3f, 0x8f, 0xee, 0x2d, 0x60, 0xfc, 0x03, 0x55, 0x2e, 0x7a, 0xdd, 0xac, 0x57, 0xe8, 0xf8, 0x90,
	0xb1, 0x5e, 0x3d, 0xba, 0xf3, 0x2d, 0xbc, 0xb8, 0x2e, 0x09, 0x56, 0xf4, 0x52, 0x6b, 0xf5, 0xff,
	0x2a, 0xe8, 0x85, 0xa3, 0xb5, 0xff, 0x66, 0x26, 0x23, 0x7f, 0xb7, 0xba, 0xda, 0xbe, 0x8e, 0x7e,
	0xe9, 0x97, 0x26, 0xdc, 0xf4, 0xcd, 0x06, 0x7d, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0x5c, 0xfa, 0xc9, 0xe4, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UCulqiClient is the client API for UCulqi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UCulqiClient interface {
	ExecuteCharge(ctx context.Context, in *MinimalInvoice, opts ...grpc.CallOption) (*Invoice, error)
	ExecuteChargeWithOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Invoice, error)
	GetInvoiceByID(ctx context.Context, in *InvoiceID, opts ...grpc.CallOption) (*Invoice, error)
	UpdateEmailTemplate(ctx context.Context, in *TemplateData, opts ...grpc.CallOption) (*IsOk, error)
}

type uCulqiClient struct {
	cc *grpc.ClientConn
}

func NewUCulqiClient(cc *grpc.ClientConn) UCulqiClient {
	return &uCulqiClient{cc}
}

func (c *uCulqiClient) ExecuteCharge(ctx context.Context, in *MinimalInvoice, opts ...grpc.CallOption) (*Invoice, error) {
	out := new(Invoice)
	err := c.cc.Invoke(ctx, "/culqi.UCulqi/ExecuteCharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCulqiClient) ExecuteChargeWithOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Invoice, error) {
	out := new(Invoice)
	err := c.cc.Invoke(ctx, "/culqi.UCulqi/ExecuteChargeWithOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCulqiClient) GetInvoiceByID(ctx context.Context, in *InvoiceID, opts ...grpc.CallOption) (*Invoice, error) {
	out := new(Invoice)
	err := c.cc.Invoke(ctx, "/culqi.UCulqi/GetInvoiceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uCulqiClient) UpdateEmailTemplate(ctx context.Context, in *TemplateData, opts ...grpc.CallOption) (*IsOk, error) {
	out := new(IsOk)
	err := c.cc.Invoke(ctx, "/culqi.UCulqi/UpdateEmailTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UCulqiServer is the server API for UCulqi service.
type UCulqiServer interface {
	ExecuteCharge(context.Context, *MinimalInvoice) (*Invoice, error)
	ExecuteChargeWithOrder(context.Context, *Order) (*Invoice, error)
	GetInvoiceByID(context.Context, *InvoiceID) (*Invoice, error)
	UpdateEmailTemplate(context.Context, *TemplateData) (*IsOk, error)
}

func RegisterUCulqiServer(s *grpc.Server, srv UCulqiServer) {
	s.RegisterService(&_UCulqi_serviceDesc, srv)
}

func _UCulqi_ExecuteCharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinimalInvoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCulqiServer).ExecuteCharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/culqi.UCulqi/ExecuteCharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCulqiServer).ExecuteCharge(ctx, req.(*MinimalInvoice))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCulqi_ExecuteChargeWithOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCulqiServer).ExecuteChargeWithOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/culqi.UCulqi/ExecuteChargeWithOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCulqiServer).ExecuteChargeWithOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCulqi_GetInvoiceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCulqiServer).GetInvoiceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/culqi.UCulqi/GetInvoiceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCulqiServer).GetInvoiceByID(ctx, req.(*InvoiceID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UCulqi_UpdateEmailTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UCulqiServer).UpdateEmailTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/culqi.UCulqi/UpdateEmailTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UCulqiServer).UpdateEmailTemplate(ctx, req.(*TemplateData))
	}
	return interceptor(ctx, in, info, handler)
}

var _UCulqi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "culqi.UCulqi",
	HandlerType: (*UCulqiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCharge",
			Handler:    _UCulqi_ExecuteCharge_Handler,
		},
		{
			MethodName: "ExecuteChargeWithOrder",
			Handler:    _UCulqi_ExecuteChargeWithOrder_Handler,
		},
		{
			MethodName: "GetInvoiceByID",
			Handler:    _UCulqi_GetInvoiceByID_Handler,
		},
		{
			MethodName: "UpdateEmailTemplate",
			Handler:    _UCulqi_UpdateEmailTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/system.proto",
}
