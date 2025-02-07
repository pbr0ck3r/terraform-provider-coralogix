// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: com/coralogixapis/dashboards/v1/ast/widget.proto

package golang

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *UUID                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Definition  *Widget_Definition      `protobuf:"bytes,4,opt,name=definition,proto3" json:"definition,omitempty"`
	Appearance  *Widget_Appearance      `protobuf:"bytes,5,opt,name=appearance,proto3" json:"appearance,omitempty"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescGZIP(), []int{0}
}

func (x *Widget) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Widget) GetTitle() *wrapperspb.StringValue {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Widget) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Widget) GetDefinition() *Widget_Definition {
	if x != nil {
		return x.Definition
	}
	return nil
}

func (x *Widget) GetAppearance() *Widget_Appearance {
	if x != nil {
		return x.Appearance
	}
	return nil
}

type Widget_Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Widget_Definition_LineChart
	//	*Widget_Definition_DataTable
	//	*Widget_Definition_Gauge
	//	*Widget_Definition_PieChart
	//	*Widget_Definition_BarChart
	Value isWidget_Definition_Value `protobuf_oneof:"value"`
}

func (x *Widget_Definition) Reset() {
	*x = Widget_Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget_Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget_Definition) ProtoMessage() {}

func (x *Widget_Definition) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget_Definition.ProtoReflect.Descriptor instead.
func (*Widget_Definition) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Widget_Definition) GetValue() isWidget_Definition_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Widget_Definition) GetLineChart() *LineChart {
	if x, ok := x.GetValue().(*Widget_Definition_LineChart); ok {
		return x.LineChart
	}
	return nil
}

func (x *Widget_Definition) GetDataTable() *DataTable {
	if x, ok := x.GetValue().(*Widget_Definition_DataTable); ok {
		return x.DataTable
	}
	return nil
}

func (x *Widget_Definition) GetGauge() *Gauge {
	if x, ok := x.GetValue().(*Widget_Definition_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (x *Widget_Definition) GetPieChart() *PieChart {
	if x, ok := x.GetValue().(*Widget_Definition_PieChart); ok {
		return x.PieChart
	}
	return nil
}

func (x *Widget_Definition) GetBarChart() *BarChart {
	if x, ok := x.GetValue().(*Widget_Definition_BarChart); ok {
		return x.BarChart
	}
	return nil
}

type isWidget_Definition_Value interface {
	isWidget_Definition_Value()
}

type Widget_Definition_LineChart struct {
	LineChart *LineChart `protobuf:"bytes,1,opt,name=line_chart,json=lineChart,proto3,oneof"`
}

type Widget_Definition_DataTable struct {
	DataTable *DataTable `protobuf:"bytes,2,opt,name=data_table,json=dataTable,proto3,oneof"`
}

type Widget_Definition_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,3,opt,name=gauge,proto3,oneof"`
}

type Widget_Definition_PieChart struct {
	PieChart *PieChart `protobuf:"bytes,4,opt,name=pie_chart,json=pieChart,proto3,oneof"`
}

type Widget_Definition_BarChart struct {
	BarChart *BarChart `protobuf:"bytes,5,opt,name=bar_chart,json=barChart,proto3,oneof"`
}

func (*Widget_Definition_LineChart) isWidget_Definition_Value() {}

func (*Widget_Definition_DataTable) isWidget_Definition_Value() {}

func (*Widget_Definition_Gauge) isWidget_Definition_Value() {}

func (*Widget_Definition_PieChart) isWidget_Definition_Value() {}

func (*Widget_Definition_BarChart) isWidget_Definition_Value() {}

type Widget_Appearance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width *wrapperspb.Int32Value `protobuf:"bytes,1,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *Widget_Appearance) Reset() {
	*x = Widget_Appearance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget_Appearance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget_Appearance) ProtoMessage() {}

func (x *Widget_Appearance) ProtoReflect() protoreflect.Message {
	mi := &file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget_Appearance.ProtoReflect.Descriptor instead.
func (*Widget_Appearance) Descriptor() ([]byte, []int) {
	return file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Widget_Appearance) GetWidth() *wrapperspb.Int32Value {
	if x != nil {
		return x.Width
	}
	return nil
}

var File_com_coralogixapis_dashboards_v1_ast_widget_proto protoreflect.FileDescriptor

var file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72,
	0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x2f, 0x62, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f,
	0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74,
	0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69,
	0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f,
	0x67, 0x61, 0x75, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x74, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x70, 0x69, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x06, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x35,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x0a, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x56, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0xbf, 0x03, 0x0a, 0x0a, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x57, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61,
	0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x67, 0x61,
	0x75, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x47, 0x61, 0x75, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x05, 0x67, 0x61, 0x75, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x70, 0x69, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e,
	0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x50, 0x69, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x69, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x54, 0x0a, 0x09,
	0x62, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x61, 0x6c, 0x6f, 0x67, 0x69, 0x78, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x73, 0x74, 0x2e, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x42, 0x61,
	0x72, 0x43, 0x68, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x08, 0x62, 0x61, 0x72, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x3f, 0x0a, 0x0a, 0x41,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x42, 0x09, 0x5a, 0x07,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescOnce sync.Once
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescData = file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDesc
)

func file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescGZIP() []byte {
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescOnce.Do(func() {
		file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescData)
	})
	return file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDescData
}

var file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_com_coralogixapis_dashboards_v1_ast_widget_proto_goTypes = []interface{}{
	(*Widget)(nil),                 // 0: com.coralogixapis.dashboards.v1.ast.Widget
	(*Widget_Definition)(nil),      // 1: com.coralogixapis.dashboards.v1.ast.Widget.Definition
	(*Widget_Appearance)(nil),      // 2: com.coralogixapis.dashboards.v1.ast.Widget.Appearance
	(*UUID)(nil),                   // 3: com.coralogixapis.dashboards.v1.UUID
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*LineChart)(nil),              // 5: com.coralogixapis.dashboards.v1.ast.widgets.LineChart
	(*DataTable)(nil),              // 6: com.coralogixapis.dashboards.v1.ast.widgets.DataTable
	(*Gauge)(nil),                  // 7: com.coralogixapis.dashboards.v1.ast.widgets.Gauge
	(*PieChart)(nil),               // 8: com.coralogixapis.dashboards.v1.ast.widgets.PieChart
	(*BarChart)(nil),               // 9: com.coralogixapis.dashboards.v1.ast.widgets.BarChart
	(*wrapperspb.Int32Value)(nil),  // 10: google.protobuf.Int32Value
}
var file_com_coralogixapis_dashboards_v1_ast_widget_proto_depIdxs = []int32{
	3,  // 0: com.coralogixapis.dashboards.v1.ast.Widget.id:type_name -> com.coralogixapis.dashboards.v1.UUID
	4,  // 1: com.coralogixapis.dashboards.v1.ast.Widget.title:type_name -> google.protobuf.StringValue
	4,  // 2: com.coralogixapis.dashboards.v1.ast.Widget.description:type_name -> google.protobuf.StringValue
	1,  // 3: com.coralogixapis.dashboards.v1.ast.Widget.definition:type_name -> com.coralogixapis.dashboards.v1.ast.Widget.Definition
	2,  // 4: com.coralogixapis.dashboards.v1.ast.Widget.appearance:type_name -> com.coralogixapis.dashboards.v1.ast.Widget.Appearance
	5,  // 5: com.coralogixapis.dashboards.v1.ast.Widget.Definition.line_chart:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.LineChart
	6,  // 6: com.coralogixapis.dashboards.v1.ast.Widget.Definition.data_table:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.DataTable
	7,  // 7: com.coralogixapis.dashboards.v1.ast.Widget.Definition.gauge:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.Gauge
	8,  // 8: com.coralogixapis.dashboards.v1.ast.Widget.Definition.pie_chart:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.PieChart
	9,  // 9: com.coralogixapis.dashboards.v1.ast.Widget.Definition.bar_chart:type_name -> com.coralogixapis.dashboards.v1.ast.widgets.BarChart
	10, // 10: com.coralogixapis.dashboards.v1.ast.Widget.Appearance.width:type_name -> google.protobuf.Int32Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_com_coralogixapis_dashboards_v1_ast_widget_proto_init() }
func file_com_coralogixapis_dashboards_v1_ast_widget_proto_init() {
	if File_com_coralogixapis_dashboards_v1_ast_widget_proto != nil {
		return
	}
	file_com_coralogixapis_dashboards_v1_ast_widgets_bar_chart_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_data_table_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_gauge_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_line_chart_proto_init()
	file_com_coralogixapis_dashboards_v1_ast_widgets_pie_chart_proto_init()
	file_com_coralogixapis_dashboards_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget); i {
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
		file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget_Definition); i {
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
		file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget_Appearance); i {
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
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Widget_Definition_LineChart)(nil),
		(*Widget_Definition_DataTable)(nil),
		(*Widget_Definition_Gauge)(nil),
		(*Widget_Definition_PieChart)(nil),
		(*Widget_Definition_BarChart)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_coralogixapis_dashboards_v1_ast_widget_proto_goTypes,
		DependencyIndexes: file_com_coralogixapis_dashboards_v1_ast_widget_proto_depIdxs,
		MessageInfos:      file_com_coralogixapis_dashboards_v1_ast_widget_proto_msgTypes,
	}.Build()
	File_com_coralogixapis_dashboards_v1_ast_widget_proto = out.File
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_rawDesc = nil
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_goTypes = nil
	file_com_coralogixapis_dashboards_v1_ast_widget_proto_depIdxs = nil
}
