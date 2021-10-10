// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/alert.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Operators for combining conditions.
type AlertPolicy_ConditionCombinerType int32

const (
	// An unspecified combiner.
	AlertPolicy_COMBINE_UNSPECIFIED AlertPolicy_ConditionCombinerType = 0
	// Combine conditions using the logical `AND` operator. An
	// incident is created only if all conditions are met
	// simultaneously. This combiner is satisfied if all conditions are
	// met, even if they are met on completely different resources.
	AlertPolicy_AND AlertPolicy_ConditionCombinerType = 1
	// Combine conditions using the logical `OR` operator. An incident
	// is created if any of the listed conditions is met.
	AlertPolicy_OR AlertPolicy_ConditionCombinerType = 2
	// Combine conditions using logical `AND` operator, but unlike the regular
	// `AND` option, an incident is created only if all conditions are met
	// simultaneously on at least one resource.
	AlertPolicy_AND_WITH_MATCHING_RESOURCE AlertPolicy_ConditionCombinerType = 3
)

var AlertPolicy_ConditionCombinerType_name = map[int32]string{
	0: "COMBINE_UNSPECIFIED",
	1: "AND",
	2: "OR",
	3: "AND_WITH_MATCHING_RESOURCE",
}

var AlertPolicy_ConditionCombinerType_value = map[string]int32{
	"COMBINE_UNSPECIFIED":        0,
	"AND":                        1,
	"OR":                         2,
	"AND_WITH_MATCHING_RESOURCE": 3,
}

func (x AlertPolicy_ConditionCombinerType) String() string {
	return proto.EnumName(AlertPolicy_ConditionCombinerType_name, int32(x))
}

func (AlertPolicy_ConditionCombinerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 0}
}

// A description of the conditions under which some aspect of your system is
// considered to be "unhealthy" and the ways to notify people or services about
// this state. For an overview of alert policies, see
// [Introduction to Alerting](/monitoring/alerts/).
type AlertPolicy struct {
	// Required if the policy exists. The resource name for this policy. The
	// syntax is:
	//
	//     projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	//
	// `[ALERT_POLICY_ID]` is assigned by Stackdriver Monitoring when the policy
	// is created.  When calling the
	// [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	// method, do not include the `name` field in the alerting policy passed as
	// part of the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A short name or phrase used to identify the policy in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple policies in the same project. The name is
	// limited to 512 Unicode characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Documentation that is included with notifications and incidents related to
	// this policy. Best practice is for the documentation to include information
	// to help responders understand, mitigate, escalate, and correct the
	// underlying problems detected by the alerting policy. Notification channels
	// that have limited capacity might not show this documentation.
	Documentation *AlertPolicy_Documentation `protobuf:"bytes,13,opt,name=documentation,proto3" json:"documentation,omitempty"`
	// User-supplied key/value data to be used for organizing and
	// identifying the `AlertPolicy` objects.
	//
	// The field can contain up to 64 entries. Each key and value is limited to
	// 63 Unicode characters or 128 bytes, whichever is smaller. Labels and
	// values can contain only lowercase letters, numerals, underscores, and
	// dashes. Keys must begin with a letter.
	UserLabels map[string]string `protobuf:"bytes,16,rep,name=user_labels,json=userLabels,proto3" json:"user_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of conditions for the policy. The conditions are combined by AND or
	// OR according to the `combiner` field. If the combined conditions evaluate
	// to true, then an incident is created. A policy can have from one to six
	// conditions.
	Conditions []*AlertPolicy_Condition `protobuf:"bytes,12,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// How to combine the results of multiple conditions to determine if an
	// incident should be opened.
	Combiner AlertPolicy_ConditionCombinerType `protobuf:"varint,6,opt,name=combiner,proto3,enum=google.monitoring.v3.AlertPolicy_ConditionCombinerType" json:"combiner,omitempty"`
	// Whether or not the policy is enabled. On write, the default interpretation
	// if unset is that the policy is enabled. On read, clients should not make
	// any assumption about the state if it has not been populated. The
	// field should always be populated on List and Get operations, unless
	// a field projection has been specified that strips it out.
	Enabled *wrappers.BoolValue `protobuf:"bytes,17,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Read-only description of how the alert policy is invalid. OK if the alert
	// policy is valid. If not OK, the alert policy will not generate incidents.
	Validity *status.Status `protobuf:"bytes,18,opt,name=validity,proto3" json:"validity,omitempty"`
	// Identifies the notification channels to which notifications should be sent
	// when incidents are opened or closed or when new violations occur on
	// an already opened incident. Each element of this array corresponds to
	// the `name` field in each of the
	// [`NotificationChannel`][google.monitoring.v3.NotificationChannel]
	// objects that are returned from the [`ListNotificationChannels`]
	// [google.monitoring.v3.NotificationChannelService.ListNotificationChannels]
	// method. The syntax of the entries in this field is:
	//
	//     projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	NotificationChannels []string `protobuf:"bytes,14,rep,name=notification_channels,json=notificationChannels,proto3" json:"notification_channels,omitempty"`
	// A read-only record of the creation of the alerting policy. If provided
	// in a call to create or update, this field will be ignored.
	CreationRecord *MutationRecord `protobuf:"bytes,10,opt,name=creation_record,json=creationRecord,proto3" json:"creation_record,omitempty"`
	// A read-only record of the most recent change to the alerting policy. If
	// provided in a call to create or update, this field will be ignored.
	MutationRecord       *MutationRecord `protobuf:"bytes,11,opt,name=mutation_record,json=mutationRecord,proto3" json:"mutation_record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AlertPolicy) Reset()         { *m = AlertPolicy{} }
func (m *AlertPolicy) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy) ProtoMessage()    {}
func (*AlertPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0}
}

func (m *AlertPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy.Unmarshal(m, b)
}
func (m *AlertPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy.Marshal(b, m, deterministic)
}
func (m *AlertPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy.Merge(m, src)
}
func (m *AlertPolicy) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy.Size(m)
}
func (m *AlertPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy proto.InternalMessageInfo

func (m *AlertPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertPolicy) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AlertPolicy) GetDocumentation() *AlertPolicy_Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *AlertPolicy) GetUserLabels() map[string]string {
	if m != nil {
		return m.UserLabels
	}
	return nil
}

func (m *AlertPolicy) GetConditions() []*AlertPolicy_Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *AlertPolicy) GetCombiner() AlertPolicy_ConditionCombinerType {
	if m != nil {
		return m.Combiner
	}
	return AlertPolicy_COMBINE_UNSPECIFIED
}

func (m *AlertPolicy) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *AlertPolicy) GetValidity() *status.Status {
	if m != nil {
		return m.Validity
	}
	return nil
}

func (m *AlertPolicy) GetNotificationChannels() []string {
	if m != nil {
		return m.NotificationChannels
	}
	return nil
}

func (m *AlertPolicy) GetCreationRecord() *MutationRecord {
	if m != nil {
		return m.CreationRecord
	}
	return nil
}

func (m *AlertPolicy) GetMutationRecord() *MutationRecord {
	if m != nil {
		return m.MutationRecord
	}
	return nil
}

// A content string and a MIME type that describes the content string's
// format.
type AlertPolicy_Documentation struct {
	// The text of the documentation, interpreted according to `mime_type`.
	// The content may not exceed 8,192 Unicode characters and may not exceed
	// more than 10,240 bytes when encoded in UTF-8 format, whichever is
	// smaller.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The format of the `content` field. Presently, only the value
	// `"text/markdown"` is supported. See
	// [Markdown](https://en.wikipedia.org/wiki/Markdown) for more information.
	MimeType             string   `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertPolicy_Documentation) Reset()         { *m = AlertPolicy_Documentation{} }
func (m *AlertPolicy_Documentation) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Documentation) ProtoMessage()    {}
func (*AlertPolicy_Documentation) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 0}
}

func (m *AlertPolicy_Documentation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy_Documentation.Unmarshal(m, b)
}
func (m *AlertPolicy_Documentation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy_Documentation.Marshal(b, m, deterministic)
}
func (m *AlertPolicy_Documentation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy_Documentation.Merge(m, src)
}
func (m *AlertPolicy_Documentation) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy_Documentation.Size(m)
}
func (m *AlertPolicy_Documentation) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy_Documentation.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy_Documentation proto.InternalMessageInfo

func (m *AlertPolicy_Documentation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *AlertPolicy_Documentation) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// A condition is a true/false test that determines when an alerting policy
// should open an incident. If a condition evaluates to true, it signifies
// that something is wrong.
type AlertPolicy_Condition struct {
	// Required if the condition exists. The unique resource name for this
	// condition. Its syntax is:
	//
	//     projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	//
	// `[CONDITION_ID]` is assigned by Stackdriver Monitoring when the
	// condition is created as part of a new or updated alerting policy.
	//
	// When calling the
	// [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	// method, do not include the `name` field in the conditions of the
	// requested alerting policy. Stackdriver Monitoring creates the
	// condition identifiers and includes them in the new policy.
	//
	// When calling the
	// [alertPolicies.update][google.monitoring.v3.AlertPolicyService.UpdateAlertPolicy]
	// method to update a policy, including a condition `name` causes the
	// existing condition to be updated. Conditions without names are added to
	// the updated policy. Existing conditions are deleted if they are not
	// updated.
	//
	// Best practice is to preserve `[CONDITION_ID]` if you make only small
	// changes, such as those to condition thresholds, durations, or trigger
	// values.  Otherwise, treat the change as a new condition and let the
	// existing condition be deleted.
	Name string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	// A short name or phrase used to identify the condition in dashboards,
	// notifications, and incidents. To avoid confusion, don't use the same
	// display name for multiple conditions in the same policy.
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Only one of the following condition types will be specified.
	//
	// Types that are valid to be assigned to Condition:
	//	*AlertPolicy_Condition_ConditionThreshold
	//	*AlertPolicy_Condition_ConditionAbsent
	Condition            isAlertPolicy_Condition_Condition `protobuf_oneof:"condition"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AlertPolicy_Condition) Reset()         { *m = AlertPolicy_Condition{} }
func (m *AlertPolicy_Condition) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition) ProtoMessage()    {}
func (*AlertPolicy_Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 1}
}

func (m *AlertPolicy_Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy_Condition.Unmarshal(m, b)
}
func (m *AlertPolicy_Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy_Condition.Marshal(b, m, deterministic)
}
func (m *AlertPolicy_Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy_Condition.Merge(m, src)
}
func (m *AlertPolicy_Condition) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy_Condition.Size(m)
}
func (m *AlertPolicy_Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy_Condition proto.InternalMessageInfo

func (m *AlertPolicy_Condition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlertPolicy_Condition) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

type isAlertPolicy_Condition_Condition interface {
	isAlertPolicy_Condition_Condition()
}

type AlertPolicy_Condition_ConditionThreshold struct {
	ConditionThreshold *AlertPolicy_Condition_MetricThreshold `protobuf:"bytes,1,opt,name=condition_threshold,json=conditionThreshold,proto3,oneof"`
}

type AlertPolicy_Condition_ConditionAbsent struct {
	ConditionAbsent *AlertPolicy_Condition_MetricAbsence `protobuf:"bytes,2,opt,name=condition_absent,json=conditionAbsent,proto3,oneof"`
}

func (*AlertPolicy_Condition_ConditionThreshold) isAlertPolicy_Condition_Condition() {}

func (*AlertPolicy_Condition_ConditionAbsent) isAlertPolicy_Condition_Condition() {}

func (m *AlertPolicy_Condition) GetCondition() isAlertPolicy_Condition_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *AlertPolicy_Condition) GetConditionThreshold() *AlertPolicy_Condition_MetricThreshold {
	if x, ok := m.GetCondition().(*AlertPolicy_Condition_ConditionThreshold); ok {
		return x.ConditionThreshold
	}
	return nil
}

func (m *AlertPolicy_Condition) GetConditionAbsent() *AlertPolicy_Condition_MetricAbsence {
	if x, ok := m.GetCondition().(*AlertPolicy_Condition_ConditionAbsent); ok {
		return x.ConditionAbsent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AlertPolicy_Condition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AlertPolicy_Condition_ConditionThreshold)(nil),
		(*AlertPolicy_Condition_ConditionAbsent)(nil),
	}
}

// Specifies how many time series must fail a predicate to trigger a
// condition. If not specified, then a `{count: 1}` trigger is used.
type AlertPolicy_Condition_Trigger struct {
	// A type of trigger.
	//
	// Types that are valid to be assigned to Type:
	//	*AlertPolicy_Condition_Trigger_Count
	//	*AlertPolicy_Condition_Trigger_Percent
	Type                 isAlertPolicy_Condition_Trigger_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AlertPolicy_Condition_Trigger) Reset()         { *m = AlertPolicy_Condition_Trigger{} }
func (m *AlertPolicy_Condition_Trigger) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_Trigger) ProtoMessage()    {}
func (*AlertPolicy_Condition_Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 1, 0}
}

func (m *AlertPolicy_Condition_Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy_Condition_Trigger.Unmarshal(m, b)
}
func (m *AlertPolicy_Condition_Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy_Condition_Trigger.Marshal(b, m, deterministic)
}
func (m *AlertPolicy_Condition_Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy_Condition_Trigger.Merge(m, src)
}
func (m *AlertPolicy_Condition_Trigger) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy_Condition_Trigger.Size(m)
}
func (m *AlertPolicy_Condition_Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy_Condition_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy_Condition_Trigger proto.InternalMessageInfo

type isAlertPolicy_Condition_Trigger_Type interface {
	isAlertPolicy_Condition_Trigger_Type()
}

type AlertPolicy_Condition_Trigger_Count struct {
	Count int32 `protobuf:"varint,1,opt,name=count,proto3,oneof"`
}

type AlertPolicy_Condition_Trigger_Percent struct {
	Percent float64 `protobuf:"fixed64,2,opt,name=percent,proto3,oneof"`
}

func (*AlertPolicy_Condition_Trigger_Count) isAlertPolicy_Condition_Trigger_Type() {}

func (*AlertPolicy_Condition_Trigger_Percent) isAlertPolicy_Condition_Trigger_Type() {}

func (m *AlertPolicy_Condition_Trigger) GetType() isAlertPolicy_Condition_Trigger_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AlertPolicy_Condition_Trigger) GetCount() int32 {
	if x, ok := m.GetType().(*AlertPolicy_Condition_Trigger_Count); ok {
		return x.Count
	}
	return 0
}

func (m *AlertPolicy_Condition_Trigger) GetPercent() float64 {
	if x, ok := m.GetType().(*AlertPolicy_Condition_Trigger_Percent); ok {
		return x.Percent
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AlertPolicy_Condition_Trigger) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AlertPolicy_Condition_Trigger_Count)(nil),
		(*AlertPolicy_Condition_Trigger_Percent)(nil),
	}
}

// A condition type that compares a collection of time series
// against a threshold.
type AlertPolicy_Condition_MetricThreshold struct {
	// A [filter](/monitoring/api/v3/filters) that
	// identifies which time series should be compared with the threshold.
	//
	// The filter is similar to the one that is specified in the
	// [`ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that
	// call is useful to verify the time series that will be retrieved /
	// processed) and must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	Filter string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	// Specifies the alignment of data points in individual time series as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.
	//
	// This field is similar to the one in the [`ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list). It
	// is advisable to use the `ListTimeSeries` method when debugging this
	// field.
	Aggregations []*Aggregation `protobuf:"bytes,8,rep,name=aggregations,proto3" json:"aggregations,omitempty"`
	// A [filter](/monitoring/api/v3/filters) that identifies a time
	// series that should be used as the denominator of a ratio that will be
	// compared with the threshold. If a `denominator_filter` is specified,
	// the time series specified by the `filter` field will be used as the
	// numerator.
	//
	// The filter must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	DenominatorFilter string `protobuf:"bytes,9,opt,name=denominator_filter,json=denominatorFilter,proto3" json:"denominator_filter,omitempty"`
	// Specifies the alignment of data points in individual time series
	// selected by `denominatorFilter` as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resources).
	//
	// When computing ratios, the `aggregations` and
	// `denominator_aggregations` fields must use the same alignment period
	// and produce time series that have the same periodicity and labels.
	DenominatorAggregations []*Aggregation `protobuf:"bytes,10,rep,name=denominator_aggregations,json=denominatorAggregations,proto3" json:"denominator_aggregations,omitempty"`
	// The comparison to apply between the time series (indicated by `filter`
	// and `aggregation`) and the threshold (indicated by `threshold_value`).
	// The comparison is applied on each time series, with the time series
	// on the left-hand side and the threshold on the right-hand side.
	//
	// Only `COMPARISON_LT` and `COMPARISON_GT` are supported currently.
	Comparison ComparisonType `protobuf:"varint,4,opt,name=comparison,proto3,enum=google.monitoring.v3.ComparisonType" json:"comparison,omitempty"`
	// A value against which to compare the time series.
	ThresholdValue float64 `protobuf:"fixed64,5,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	// The amount of time that a time series must violate the
	// threshold to be considered failing. Currently, only values
	// that are a multiple of a minute--e.g., 0, 60, 120, or 300
	// seconds--are supported. If an invalid value is given, an
	// error will be returned. When choosing a duration, it is useful to
	// keep in mind the frequency of the underlying time series data
	// (which may also be affected by any alignments specified in the
	// `aggregations` field); a good duration is long enough so that a single
	// outlier does not generate spurious alerts, but short enough that
	// unhealthy states are detected and alerted on quickly.
	Duration *duration.Duration `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// The number/percent of time series for which the comparison must hold
	// in order for the condition to trigger. If unspecified, then the
	// condition will trigger if the comparison is true for any of the
	// time series that have been identified by `filter` and `aggregations`,
	// or by the ratio, if `denominator_filter` and `denominator_aggregations`
	// are specified.
	Trigger              *AlertPolicy_Condition_Trigger `protobuf:"bytes,7,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AlertPolicy_Condition_MetricThreshold) Reset()         { *m = AlertPolicy_Condition_MetricThreshold{} }
func (m *AlertPolicy_Condition_MetricThreshold) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_MetricThreshold) ProtoMessage()    {}
func (*AlertPolicy_Condition_MetricThreshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 1, 1}
}

func (m *AlertPolicy_Condition_MetricThreshold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy_Condition_MetricThreshold.Unmarshal(m, b)
}
func (m *AlertPolicy_Condition_MetricThreshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy_Condition_MetricThreshold.Marshal(b, m, deterministic)
}
func (m *AlertPolicy_Condition_MetricThreshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy_Condition_MetricThreshold.Merge(m, src)
}
func (m *AlertPolicy_Condition_MetricThreshold) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy_Condition_MetricThreshold.Size(m)
}
func (m *AlertPolicy_Condition_MetricThreshold) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy_Condition_MetricThreshold.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy_Condition_MetricThreshold proto.InternalMessageInfo

func (m *AlertPolicy_Condition_MetricThreshold) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricThreshold) GetAggregations() []*Aggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDenominatorFilter() string {
	if m != nil {
		return m.DenominatorFilter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDenominatorAggregations() []*Aggregation {
	if m != nil {
		return m.DenominatorAggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetComparison() ComparisonType {
	if m != nil {
		return m.Comparison
	}
	return ComparisonType_COMPARISON_UNSPECIFIED
}

func (m *AlertPolicy_Condition_MetricThreshold) GetThresholdValue() float64 {
	if m != nil {
		return m.ThresholdValue
	}
	return 0
}

func (m *AlertPolicy_Condition_MetricThreshold) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricThreshold) GetTrigger() *AlertPolicy_Condition_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

// A condition type that checks that monitored resources
// are reporting data. The configuration defines a metric and
// a set of monitored resources. The predicate is considered in violation
// when a time series for the specified metric of a monitored
// resource does not include any data in the specified `duration`.
type AlertPolicy_Condition_MetricAbsence struct {
	// A [filter](/monitoring/api/v3/filters) that
	// identifies which time series should be compared with the threshold.
	//
	// The filter is similar to the one that is specified in the
	// [`ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that
	// call is useful to verify the time series that will be retrieved /
	// processed) and must specify the metric type and optionally may contain
	// restrictions on resource type, resource labels, and metric labels.
	// This field may not exceed 2048 Unicode characters in length.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// Specifies the alignment of data points in individual time series as
	// well as how to combine the retrieved time series together (such as
	// when aggregating multiple streams on each resource to a single
	// stream for each resource or when aggregating streams across all
	// members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.
	//
	// This field is similar to the one in the [`ListTimeSeries`
	// request](/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list). It
	// is advisable to use the `ListTimeSeries` method when debugging this
	// field.
	Aggregations []*Aggregation `protobuf:"bytes,5,rep,name=aggregations,proto3" json:"aggregations,omitempty"`
	// The amount of time that a time series must fail to report new
	// data to be considered failing. Currently, only values that
	// are a multiple of a minute--e.g.  60, 120, or 300
	// seconds--are supported. If an invalid value is given, an
	// error will be returned. The `Duration.nanos` field is
	// ignored.
	Duration *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// The number/percent of time series for which the comparison must hold
	// in order for the condition to trigger. If unspecified, then the
	// condition will trigger if the comparison is true for any of the
	// time series that have been identified by `filter` and `aggregations`.
	Trigger              *AlertPolicy_Condition_Trigger `protobuf:"bytes,3,opt,name=trigger,proto3" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AlertPolicy_Condition_MetricAbsence) Reset()         { *m = AlertPolicy_Condition_MetricAbsence{} }
func (m *AlertPolicy_Condition_MetricAbsence) String() string { return proto.CompactTextString(m) }
func (*AlertPolicy_Condition_MetricAbsence) ProtoMessage()    {}
func (*AlertPolicy_Condition_MetricAbsence) Descriptor() ([]byte, []int) {
	return fileDescriptor_014ef0e1a0f00a00, []int{0, 1, 2}
}

func (m *AlertPolicy_Condition_MetricAbsence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertPolicy_Condition_MetricAbsence.Unmarshal(m, b)
}
func (m *AlertPolicy_Condition_MetricAbsence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertPolicy_Condition_MetricAbsence.Marshal(b, m, deterministic)
}
func (m *AlertPolicy_Condition_MetricAbsence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertPolicy_Condition_MetricAbsence.Merge(m, src)
}
func (m *AlertPolicy_Condition_MetricAbsence) XXX_Size() int {
	return xxx_messageInfo_AlertPolicy_Condition_MetricAbsence.Size(m)
}
func (m *AlertPolicy_Condition_MetricAbsence) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertPolicy_Condition_MetricAbsence.DiscardUnknown(m)
}

var xxx_messageInfo_AlertPolicy_Condition_MetricAbsence proto.InternalMessageInfo

func (m *AlertPolicy_Condition_MetricAbsence) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *AlertPolicy_Condition_MetricAbsence) GetAggregations() []*Aggregation {
	if m != nil {
		return m.Aggregations
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricAbsence) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *AlertPolicy_Condition_MetricAbsence) GetTrigger() *AlertPolicy_Condition_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.monitoring.v3.AlertPolicy_ConditionCombinerType", AlertPolicy_ConditionCombinerType_name, AlertPolicy_ConditionCombinerType_value)
	proto.RegisterType((*AlertPolicy)(nil), "google.monitoring.v3.AlertPolicy")
	proto.RegisterMapType((map[string]string)(nil), "google.monitoring.v3.AlertPolicy.UserLabelsEntry")
	proto.RegisterType((*AlertPolicy_Documentation)(nil), "google.monitoring.v3.AlertPolicy.Documentation")
	proto.RegisterType((*AlertPolicy_Condition)(nil), "google.monitoring.v3.AlertPolicy.Condition")
	proto.RegisterType((*AlertPolicy_Condition_Trigger)(nil), "google.monitoring.v3.AlertPolicy.Condition.Trigger")
	proto.RegisterType((*AlertPolicy_Condition_MetricThreshold)(nil), "google.monitoring.v3.AlertPolicy.Condition.MetricThreshold")
	proto.RegisterType((*AlertPolicy_Condition_MetricAbsence)(nil), "google.monitoring.v3.AlertPolicy.Condition.MetricAbsence")
}

func init() { proto.RegisterFile("google/monitoring/v3/alert.proto", fileDescriptor_014ef0e1a0f00a00) }

var fileDescriptor_014ef0e1a0f00a00 = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xfd, 0x6e, 0xe3, 0x44,
	0x10, 0xaf, 0x93, 0xe6, 0x6b, 0xd2, 0x36, 0xb9, 0xbd, 0xde, 0xd5, 0x18, 0xe9, 0x94, 0x9e, 0x90,
	0x88, 0x40, 0x38, 0x22, 0x01, 0xf1, 0x71, 0x02, 0x29, 0x5f, 0xbd, 0x44, 0x90, 0xb4, 0xda, 0xa6,
	0x45, 0x42, 0x95, 0x2c, 0xc7, 0xde, 0xba, 0x16, 0xf6, 0xae, 0xb5, 0xb6, 0x8b, 0xf2, 0x18, 0xbc,
	0x02, 0x7f, 0xf2, 0x28, 0x3c, 0x02, 0x8f, 0x81, 0x78, 0x00, 0xe4, 0xf5, 0x47, 0x9c, 0x5e, 0x7a,
	0x47, 0x8e, 0xff, 0x32, 0x3b, 0xbf, 0xf9, 0xcd, 0xce, 0xcc, 0x6f, 0xd6, 0x81, 0x96, 0xc5, 0x98,
	0xe5, 0x90, 0x8e, 0xcb, 0xa8, 0x1d, 0x30, 0x6e, 0x53, 0xab, 0x73, 0xdf, 0xeb, 0xe8, 0x0e, 0xe1,
	0x81, 0xea, 0x71, 0x16, 0x30, 0x74, 0x1c, 0x23, 0xd4, 0x35, 0x42, 0xbd, 0xef, 0x29, 0xa7, 0x5b,
	0xe3, 0x0c, 0xe6, 0xba, 0x8c, 0xc6, 0x81, 0xca, 0x27, 0x5b, 0x21, 0x6e, 0x18, 0xe8, 0x81, 0xcd,
	0xa8, 0xc6, 0x89, 0xc1, 0xb8, 0x99, 0x60, 0x5f, 0x24, 0x58, 0x61, 0x2d, 0xc3, 0xdb, 0x8e, 0x19,
	0x72, 0x01, 0x7b, 0xcc, 0xff, 0x2b, 0xd7, 0x3d, 0x8f, 0x70, 0x3f, 0xf1, 0x9f, 0x24, 0x7e, 0xee,
	0x19, 0x1d, 0x3f, 0xd0, 0x83, 0x30, 0x71, 0xbc, 0xfc, 0xad, 0x09, 0xf5, 0x7e, 0x54, 0xcd, 0x05,
	0x73, 0x6c, 0x63, 0x85, 0x10, 0xec, 0x53, 0xdd, 0x25, 0xb2, 0xd4, 0x92, 0xda, 0x35, 0x2c, 0x7e,
	0xa3, 0x53, 0x38, 0x30, 0x6d, 0xdf, 0x73, 0xf4, 0x95, 0x26, 0x7c, 0x05, 0xe1, 0xab, 0x27, 0x67,
	0xf3, 0x08, 0x72, 0x05, 0x87, 0x26, 0x33, 0x42, 0x97, 0xd0, 0xf8, 0xf6, 0xf2, 0x61, 0x4b, 0x6a,
	0xd7, 0xbb, 0x1d, 0x75, 0x5b, 0x73, 0xd4, 0x5c, 0x42, 0x75, 0x94, 0x0f, 0xc3, 0x9b, 0x2c, 0x08,
	0x43, 0x3d, 0xf4, 0x09, 0xd7, 0x1c, 0x7d, 0x49, 0x1c, 0x5f, 0x6e, 0xb6, 0x8a, 0xed, 0x7a, 0xf7,
	0xf3, 0x77, 0x93, 0x5e, 0xf9, 0x84, 0xff, 0x28, 0x62, 0xc6, 0x34, 0xe0, 0x2b, 0x0c, 0x61, 0x76,
	0x80, 0x7e, 0x00, 0x30, 0x18, 0x35, 0xed, 0x28, 0x81, 0x2f, 0x1f, 0x08, 0xca, 0x4f, 0xdf, 0x4d,
	0x39, 0x4c, 0x63, 0x70, 0x2e, 0x1c, 0x5d, 0x42, 0xd5, 0x60, 0xee, 0xd2, 0xa6, 0x84, 0xcb, 0xe5,
	0x96, 0xd4, 0x3e, 0xea, 0x7e, 0xb5, 0x03, 0xd5, 0x30, 0x09, 0x5d, 0xac, 0x3c, 0x82, 0x33, 0x22,
	0xf4, 0x05, 0x54, 0x08, 0xd5, 0x97, 0x0e, 0x31, 0xe5, 0x27, 0xa2, 0x8d, 0x4a, 0xca, 0x99, 0x8e,
	0x57, 0x1d, 0x30, 0xe6, 0x5c, 0xeb, 0x4e, 0x48, 0x70, 0x0a, 0x45, 0x2a, 0x54, 0xef, 0x75, 0xc7,
	0x36, 0xed, 0x60, 0x25, 0x23, 0x11, 0x86, 0xd2, 0x30, 0xee, 0x19, 0xea, 0xa5, 0x98, 0x3a, 0xce,
	0x30, 0xa8, 0x07, 0xcf, 0x28, 0x0b, 0xec, 0x5b, 0xdb, 0x88, 0xf5, 0x66, 0xdc, 0xe9, 0x94, 0x46,
	0x5d, 0x3e, 0x6a, 0x15, 0xdb, 0x35, 0x7c, 0x9c, 0x77, 0x0e, 0x13, 0x1f, 0x9a, 0x41, 0xc3, 0xe0,
	0x24, 0x2f, 0x50, 0x19, 0x44, 0xae, 0x8f, 0xb6, 0x97, 0x3d, 0x4b, 0xd4, 0x8c, 0x05, 0x16, 0x1f,
	0xa5, 0xc1, 0xb1, 0x1d, 0xd1, 0x3d, 0xd0, 0xbb, 0x5c, 0xdf, 0x85, 0xce, 0xdd, 0xb0, 0x95, 0x33,
	0x38, 0xdc, 0x90, 0x13, 0x92, 0xa1, 0x62, 0x30, 0x1a, 0x10, 0x1a, 0x24, 0x82, 0x4e, 0x4d, 0xf4,
	0x21, 0xd4, 0x5c, 0xdb, 0x25, 0x5a, 0xb0, 0xf2, 0x52, 0x41, 0x57, 0xa3, 0x83, 0x68, 0x14, 0xca,
	0x5f, 0x55, 0xa8, 0x65, 0x43, 0xca, 0x56, 0xe2, 0xe0, 0x2d, 0x2b, 0x51, 0x7e, 0x73, 0x25, 0x28,
	0x3c, 0xcd, 0x84, 0xa2, 0x05, 0x77, 0x9c, 0xf8, 0x77, 0xcc, 0x31, 0xc5, 0x3d, 0xea, 0xdd, 0x57,
	0x3b, 0xa8, 0x44, 0x9d, 0x91, 0x80, 0xdb, 0xc6, 0x22, 0xa5, 0x98, 0xec, 0x61, 0x94, 0x31, 0x67,
	0xa7, 0xe8, 0x16, 0x9a, 0xeb, 0x7c, 0xfa, 0xd2, 0x8f, 0x8a, 0x2e, 0x88, 0x64, 0xdf, 0xec, 0x9e,
	0xac, 0x1f, 0xc5, 0x1b, 0x64, 0xb2, 0x87, 0x1b, 0x19, 0xa9, 0x38, 0x0b, 0x94, 0x31, 0x54, 0x16,
	0xdc, 0xb6, 0x2c, 0xc2, 0xd1, 0x73, 0x28, 0x19, 0x2c, 0x4c, 0x9a, 0x5b, 0x9a, 0xec, 0xe1, 0xd8,
	0x44, 0x0a, 0x54, 0x3c, 0xc2, 0x8d, 0xf4, 0x06, 0xd2, 0x64, 0x0f, 0xa7, 0x07, 0x83, 0x32, 0xec,
	0x47, 0x3d, 0x57, 0xfe, 0x2e, 0x42, 0xe3, 0x41, 0x61, 0xe8, 0x39, 0x94, 0x6f, 0x6d, 0x27, 0x20,
	0x3c, 0x99, 0x48, 0x62, 0xa1, 0x31, 0x1c, 0xe8, 0x96, 0xc5, 0x89, 0xa5, 0xc7, 0x4b, 0x5b, 0x15,
	0x4b, 0x7b, 0xfa, 0x48, 0x59, 0x6b, 0x24, 0xde, 0x08, 0x43, 0x9f, 0x01, 0x32, 0x09, 0x65, 0xae,
	0x4d, 0xf5, 0x80, 0x71, 0x2d, 0x49, 0x55, 0x13, 0xa9, 0x9e, 0xe4, 0x3c, 0x67, 0x71, 0xd6, 0x1b,
	0x90, 0xf3, 0xf0, 0x8d, 0x1b, 0xc0, 0x7f, 0xbd, 0xc1, 0x49, 0x8e, 0xa2, 0x9f, 0xbf, 0xcc, 0x28,
	0x7a, 0x86, 0x5c, 0x4f, 0xe7, 0xb6, 0xcf, 0xa8, 0xbc, 0x2f, 0xde, 0x8e, 0x47, 0x54, 0x3f, 0xcc,
	0x70, 0xe2, 0xa1, 0xc8, 0xc5, 0xa1, 0x8f, 0xa1, 0x91, 0x49, 0x4b, 0xbb, 0x8f, 0x1e, 0x04, 0xb9,
	0x14, 0x75, 0x1c, 0x1f, 0x65, 0xc7, 0xe2, 0x99, 0x40, 0x5f, 0x42, 0x35, 0xfd, 0x64, 0x08, 0xb1,
	0xd6, 0xbb, 0x1f, 0xbc, 0xf1, 0xa8, 0x8c, 0x12, 0x00, 0xce, 0xa0, 0x68, 0x06, 0x95, 0x20, 0x1e,
	0xb6, 0x5c, 0x11, 0x51, 0xbd, 0x5d, 0xb4, 0x94, 0xe8, 0x04, 0xa7, 0x1c, 0xca, 0x3f, 0x12, 0x1c,
	0x6e, 0x08, 0x2c, 0x37, 0x72, 0xe9, 0xad, 0x23, 0x2f, 0xbd, 0xdf, 0xc8, 0xf3, 0x65, 0x17, 0xde,
	0xab, 0xec, 0xe2, 0xff, 0x2f, 0x7b, 0x50, 0x87, 0x5a, 0xb6, 0x45, 0xca, 0x77, 0xd0, 0x78, 0xf0,
	0x79, 0x42, 0x4d, 0x28, 0xfe, 0x42, 0x56, 0x49, 0x07, 0xa2, 0x9f, 0xe8, 0x18, 0x4a, 0xf1, 0x34,
	0xe3, 0x45, 0x88, 0x8d, 0x6f, 0x0b, 0x5f, 0x4b, 0x2f, 0x75, 0x78, 0xb6, 0xf5, 0xfb, 0x81, 0x4e,
	0xe0, 0xe9, 0xf0, 0x7c, 0x36, 0x98, 0xce, 0xc7, 0xda, 0xd5, 0xfc, 0xf2, 0x62, 0x3c, 0x9c, 0x9e,
	0x4d, 0xc7, 0xa3, 0xe6, 0x1e, 0xaa, 0x40, 0xb1, 0x3f, 0x1f, 0x35, 0x25, 0x54, 0x86, 0xc2, 0x39,
	0x6e, 0x16, 0xd0, 0x0b, 0x50, 0xfa, 0xf3, 0x91, 0xf6, 0xd3, 0x74, 0x31, 0xd1, 0x66, 0xfd, 0xc5,
	0x70, 0x32, 0x9d, 0xbf, 0xd6, 0xf0, 0xf8, 0xf2, 0xfc, 0x0a, 0x0f, 0xc7, 0xcd, 0xe2, 0xe0, 0x77,
	0x09, 0x64, 0x83, 0xb9, 0x5b, 0x4b, 0x1e, 0x40, 0x5c, 0x73, 0xd4, 0xbc, 0x0b, 0xe9, 0xe7, 0xef,
	0x13, 0x8c, 0xc5, 0x1c, 0x9d, 0x5a, 0x2a, 0xe3, 0x56, 0xc7, 0x22, 0x54, 0xb4, 0xb6, 0x13, 0xbb,
	0x74, 0xcf, 0xf6, 0x37, 0xff, 0xe2, 0xbc, 0x5a, 0x5b, 0x7f, 0x14, 0x94, 0xd7, 0x31, 0xc1, 0xd0,
	0x61, 0xa1, 0xa9, 0xce, 0xd6, 0xa9, 0xae, 0x7b, 0x7f, 0xa6, 0xce, 0x1b, 0xe1, 0xbc, 0x59, 0x3b,
	0x6f, 0xae, 0x7b, 0xcb, 0xb2, 0x48, 0xd2, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x41, 0x91, 0x1a,
	0x51, 0xa1, 0x09, 0x00, 0x00,
}
