// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BehaviorOnMxFailure string

// Enum values for BehaviorOnMxFailure
const (
	BehaviorOnMxFailureUseDefaultValue BehaviorOnMxFailure = "USE_DEFAULT_VALUE"
	BehaviorOnMxFailureRejectMessage   BehaviorOnMxFailure = "REJECT_MESSAGE"
)

// Values returns all known values for BehaviorOnMxFailure. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BehaviorOnMxFailure) Values() []BehaviorOnMxFailure {
	return []BehaviorOnMxFailure{
		"USE_DEFAULT_VALUE",
		"REJECT_MESSAGE",
	}
}

type BulkEmailStatus string

// Enum values for BulkEmailStatus
const (
	BulkEmailStatusSuccess                       BulkEmailStatus = "SUCCESS"
	BulkEmailStatusMessageRejected               BulkEmailStatus = "MESSAGE_REJECTED"
	BulkEmailStatusMailFromDomainNotVerified     BulkEmailStatus = "MAIL_FROM_DOMAIN_NOT_VERIFIED"
	BulkEmailStatusConfigurationSetNotFound      BulkEmailStatus = "CONFIGURATION_SET_NOT_FOUND"
	BulkEmailStatusTemplateNotFound              BulkEmailStatus = "TEMPLATE_NOT_FOUND"
	BulkEmailStatusAccountSuspended              BulkEmailStatus = "ACCOUNT_SUSPENDED"
	BulkEmailStatusAccountThrottled              BulkEmailStatus = "ACCOUNT_THROTTLED"
	BulkEmailStatusAccountDailyQuotaExceeded     BulkEmailStatus = "ACCOUNT_DAILY_QUOTA_EXCEEDED"
	BulkEmailStatusInvalidSendingPoolName        BulkEmailStatus = "INVALID_SENDING_POOL_NAME"
	BulkEmailStatusAccountSendingPaused          BulkEmailStatus = "ACCOUNT_SENDING_PAUSED"
	BulkEmailStatusConfigurationSetSendingPaused BulkEmailStatus = "CONFIGURATION_SET_SENDING_PAUSED"
	BulkEmailStatusInvalidParameter              BulkEmailStatus = "INVALID_PARAMETER"
	BulkEmailStatusTransientFailure              BulkEmailStatus = "TRANSIENT_FAILURE"
	BulkEmailStatusFailed                        BulkEmailStatus = "FAILED"
)

// Values returns all known values for BulkEmailStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BulkEmailStatus) Values() []BulkEmailStatus {
	return []BulkEmailStatus{
		"SUCCESS",
		"MESSAGE_REJECTED",
		"MAIL_FROM_DOMAIN_NOT_VERIFIED",
		"CONFIGURATION_SET_NOT_FOUND",
		"TEMPLATE_NOT_FOUND",
		"ACCOUNT_SUSPENDED",
		"ACCOUNT_THROTTLED",
		"ACCOUNT_DAILY_QUOTA_EXCEEDED",
		"INVALID_SENDING_POOL_NAME",
		"ACCOUNT_SENDING_PAUSED",
		"CONFIGURATION_SET_SENDING_PAUSED",
		"INVALID_PARAMETER",
		"TRANSIENT_FAILURE",
		"FAILED",
	}
}

type ContactLanguage string

// Enum values for ContactLanguage
const (
	ContactLanguageEn ContactLanguage = "EN"
	ContactLanguageJa ContactLanguage = "JA"
)

// Values returns all known values for ContactLanguage. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContactLanguage) Values() []ContactLanguage {
	return []ContactLanguage{
		"EN",
		"JA",
	}
}

type ContactListImportAction string

// Enum values for ContactListImportAction
const (
	ContactListImportActionDelete ContactListImportAction = "DELETE"
	ContactListImportActionPut    ContactListImportAction = "PUT"
)

// Values returns all known values for ContactListImportAction. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContactListImportAction) Values() []ContactListImportAction {
	return []ContactListImportAction{
		"DELETE",
		"PUT",
	}
}

type DataFormat string

// Enum values for DataFormat
const (
	DataFormatCsv  DataFormat = "CSV"
	DataFormatJson DataFormat = "JSON"
)

// Values returns all known values for DataFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DataFormat) Values() []DataFormat {
	return []DataFormat{
		"CSV",
		"JSON",
	}
}

type DeliverabilityDashboardAccountStatus string

// Enum values for DeliverabilityDashboardAccountStatus
const (
	DeliverabilityDashboardAccountStatusActive            DeliverabilityDashboardAccountStatus = "ACTIVE"
	DeliverabilityDashboardAccountStatusPendingExpiration DeliverabilityDashboardAccountStatus = "PENDING_EXPIRATION"
	DeliverabilityDashboardAccountStatusDisabled          DeliverabilityDashboardAccountStatus = "DISABLED"
)

// Values returns all known values for DeliverabilityDashboardAccountStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DeliverabilityDashboardAccountStatus) Values() []DeliverabilityDashboardAccountStatus {
	return []DeliverabilityDashboardAccountStatus{
		"ACTIVE",
		"PENDING_EXPIRATION",
		"DISABLED",
	}
}

type DeliverabilityTestStatus string

// Enum values for DeliverabilityTestStatus
const (
	DeliverabilityTestStatusInProgress DeliverabilityTestStatus = "IN_PROGRESS"
	DeliverabilityTestStatusCompleted  DeliverabilityTestStatus = "COMPLETED"
)

// Values returns all known values for DeliverabilityTestStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliverabilityTestStatus) Values() []DeliverabilityTestStatus {
	return []DeliverabilityTestStatus{
		"IN_PROGRESS",
		"COMPLETED",
	}
}

type DimensionValueSource string

// Enum values for DimensionValueSource
const (
	DimensionValueSourceMessageTag  DimensionValueSource = "MESSAGE_TAG"
	DimensionValueSourceEmailHeader DimensionValueSource = "EMAIL_HEADER"
	DimensionValueSourceLinkTag     DimensionValueSource = "LINK_TAG"
)

// Values returns all known values for DimensionValueSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DimensionValueSource) Values() []DimensionValueSource {
	return []DimensionValueSource{
		"MESSAGE_TAG",
		"EMAIL_HEADER",
		"LINK_TAG",
	}
}

type DkimSigningAttributesOrigin string

// Enum values for DkimSigningAttributesOrigin
const (
	DkimSigningAttributesOriginAwsSes   DkimSigningAttributesOrigin = "AWS_SES"
	DkimSigningAttributesOriginExternal DkimSigningAttributesOrigin = "EXTERNAL"
)

// Values returns all known values for DkimSigningAttributesOrigin. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DkimSigningAttributesOrigin) Values() []DkimSigningAttributesOrigin {
	return []DkimSigningAttributesOrigin{
		"AWS_SES",
		"EXTERNAL",
	}
}

type DkimSigningKeyLength string

// Enum values for DkimSigningKeyLength
const (
	DkimSigningKeyLengthRsa1024Bit DkimSigningKeyLength = "RSA_1024_BIT"
	DkimSigningKeyLengthRsa2048Bit DkimSigningKeyLength = "RSA_2048_BIT"
)

// Values returns all known values for DkimSigningKeyLength. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DkimSigningKeyLength) Values() []DkimSigningKeyLength {
	return []DkimSigningKeyLength{
		"RSA_1024_BIT",
		"RSA_2048_BIT",
	}
}

type DkimStatus string

// Enum values for DkimStatus
const (
	DkimStatusPending          DkimStatus = "PENDING"
	DkimStatusSuccess          DkimStatus = "SUCCESS"
	DkimStatusFailed           DkimStatus = "FAILED"
	DkimStatusTemporaryFailure DkimStatus = "TEMPORARY_FAILURE"
	DkimStatusNotStarted       DkimStatus = "NOT_STARTED"
)

// Values returns all known values for DkimStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DkimStatus) Values() []DkimStatus {
	return []DkimStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
		"TEMPORARY_FAILURE",
		"NOT_STARTED",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeSend             EventType = "SEND"
	EventTypeReject           EventType = "REJECT"
	EventTypeBounce           EventType = "BOUNCE"
	EventTypeComplaint        EventType = "COMPLAINT"
	EventTypeDelivery         EventType = "DELIVERY"
	EventTypeOpen             EventType = "OPEN"
	EventTypeClick            EventType = "CLICK"
	EventTypeRenderingFailure EventType = "RENDERING_FAILURE"
	EventTypeDeliveryDelay    EventType = "DELIVERY_DELAY"
	EventTypeSubscription     EventType = "SUBSCRIPTION"
)

// Values returns all known values for EventType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"SEND",
		"REJECT",
		"BOUNCE",
		"COMPLAINT",
		"DELIVERY",
		"OPEN",
		"CLICK",
		"RENDERING_FAILURE",
		"DELIVERY_DELAY",
		"SUBSCRIPTION",
	}
}

type FeatureStatus string

// Enum values for FeatureStatus
const (
	FeatureStatusEnabled  FeatureStatus = "ENABLED"
	FeatureStatusDisabled FeatureStatus = "DISABLED"
)

// Values returns all known values for FeatureStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FeatureStatus) Values() []FeatureStatus {
	return []FeatureStatus{
		"ENABLED",
		"DISABLED",
	}
}

type IdentityType string

// Enum values for IdentityType
const (
	IdentityTypeEmailAddress  IdentityType = "EMAIL_ADDRESS"
	IdentityTypeDomain        IdentityType = "DOMAIN"
	IdentityTypeManagedDomain IdentityType = "MANAGED_DOMAIN"
)

// Values returns all known values for IdentityType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IdentityType) Values() []IdentityType {
	return []IdentityType{
		"EMAIL_ADDRESS",
		"DOMAIN",
		"MANAGED_DOMAIN",
	}
}

type ImportDestinationType string

// Enum values for ImportDestinationType
const (
	ImportDestinationTypeSuppressionList ImportDestinationType = "SUPPRESSION_LIST"
	ImportDestinationTypeContactList     ImportDestinationType = "CONTACT_LIST"
)

// Values returns all known values for ImportDestinationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImportDestinationType) Values() []ImportDestinationType {
	return []ImportDestinationType{
		"SUPPRESSION_LIST",
		"CONTACT_LIST",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated    JobStatus = "CREATED"
	JobStatusProcessing JobStatus = "PROCESSING"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"CREATED",
		"PROCESSING",
		"COMPLETED",
		"FAILED",
	}
}

type ListRecommendationsFilterKey string

// Enum values for ListRecommendationsFilterKey
const (
	ListRecommendationsFilterKeyType        ListRecommendationsFilterKey = "TYPE"
	ListRecommendationsFilterKeyImpact      ListRecommendationsFilterKey = "IMPACT"
	ListRecommendationsFilterKeyStatus      ListRecommendationsFilterKey = "STATUS"
	ListRecommendationsFilterKeyResourceArn ListRecommendationsFilterKey = "RESOURCE_ARN"
)

// Values returns all known values for ListRecommendationsFilterKey. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListRecommendationsFilterKey) Values() []ListRecommendationsFilterKey {
	return []ListRecommendationsFilterKey{
		"TYPE",
		"IMPACT",
		"STATUS",
		"RESOURCE_ARN",
	}
}

type MailFromDomainStatus string

// Enum values for MailFromDomainStatus
const (
	MailFromDomainStatusPending          MailFromDomainStatus = "PENDING"
	MailFromDomainStatusSuccess          MailFromDomainStatus = "SUCCESS"
	MailFromDomainStatusFailed           MailFromDomainStatus = "FAILED"
	MailFromDomainStatusTemporaryFailure MailFromDomainStatus = "TEMPORARY_FAILURE"
)

// Values returns all known values for MailFromDomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MailFromDomainStatus) Values() []MailFromDomainStatus {
	return []MailFromDomainStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
		"TEMPORARY_FAILURE",
	}
}

type MailType string

// Enum values for MailType
const (
	MailTypeMarketing     MailType = "MARKETING"
	MailTypeTransactional MailType = "TRANSACTIONAL"
)

// Values returns all known values for MailType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (MailType) Values() []MailType {
	return []MailType{
		"MARKETING",
		"TRANSACTIONAL",
	}
}

type Metric string

// Enum values for Metric
const (
	MetricSend              Metric = "SEND"
	MetricComplaint         Metric = "COMPLAINT"
	MetricPermanentBounce   Metric = "PERMANENT_BOUNCE"
	MetricTransientBounce   Metric = "TRANSIENT_BOUNCE"
	MetricOpen              Metric = "OPEN"
	MetricClick             Metric = "CLICK"
	MetricDelivery          Metric = "DELIVERY"
	MetricDeliveryOpen      Metric = "DELIVERY_OPEN"
	MetricDeliveryClick     Metric = "DELIVERY_CLICK"
	MetricDeliveryComplaint Metric = "DELIVERY_COMPLAINT"
)

// Values returns all known values for Metric. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Metric) Values() []Metric {
	return []Metric{
		"SEND",
		"COMPLAINT",
		"PERMANENT_BOUNCE",
		"TRANSIENT_BOUNCE",
		"OPEN",
		"CLICK",
		"DELIVERY",
		"DELIVERY_OPEN",
		"DELIVERY_CLICK",
		"DELIVERY_COMPLAINT",
	}
}

type MetricDimensionName string

// Enum values for MetricDimensionName
const (
	MetricDimensionNameEmailIdentity    MetricDimensionName = "EMAIL_IDENTITY"
	MetricDimensionNameConfigurationSet MetricDimensionName = "CONFIGURATION_SET"
	MetricDimensionNameIsp              MetricDimensionName = "ISP"
)

// Values returns all known values for MetricDimensionName. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricDimensionName) Values() []MetricDimensionName {
	return []MetricDimensionName{
		"EMAIL_IDENTITY",
		"CONFIGURATION_SET",
		"ISP",
	}
}

type MetricNamespace string

// Enum values for MetricNamespace
const (
	MetricNamespaceVdm MetricNamespace = "VDM"
)

// Values returns all known values for MetricNamespace. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricNamespace) Values() []MetricNamespace {
	return []MetricNamespace{
		"VDM",
	}
}

type QueryErrorCode string

// Enum values for QueryErrorCode
const (
	QueryErrorCodeInternalFailure QueryErrorCode = "INTERNAL_FAILURE"
	QueryErrorCodeAccessDenied    QueryErrorCode = "ACCESS_DENIED"
)

// Values returns all known values for QueryErrorCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QueryErrorCode) Values() []QueryErrorCode {
	return []QueryErrorCode{
		"INTERNAL_FAILURE",
		"ACCESS_DENIED",
	}
}

type RecommendationImpact string

// Enum values for RecommendationImpact
const (
	RecommendationImpactLow  RecommendationImpact = "LOW"
	RecommendationImpactHigh RecommendationImpact = "HIGH"
)

// Values returns all known values for RecommendationImpact. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationImpact) Values() []RecommendationImpact {
	return []RecommendationImpact{
		"LOW",
		"HIGH",
	}
}

type RecommendationStatus string

// Enum values for RecommendationStatus
const (
	RecommendationStatusOpen  RecommendationStatus = "OPEN"
	RecommendationStatusFixed RecommendationStatus = "FIXED"
)

// Values returns all known values for RecommendationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationStatus) Values() []RecommendationStatus {
	return []RecommendationStatus{
		"OPEN",
		"FIXED",
	}
}

type RecommendationType string

// Enum values for RecommendationType
const (
	RecommendationTypeDkim  RecommendationType = "DKIM"
	RecommendationTypeDmarc RecommendationType = "DMARC"
	RecommendationTypeSpf   RecommendationType = "SPF"
	RecommendationTypeBimi  RecommendationType = "BIMI"
)

// Values returns all known values for RecommendationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecommendationType) Values() []RecommendationType {
	return []RecommendationType{
		"DKIM",
		"DMARC",
		"SPF",
		"BIMI",
	}
}

type ReviewStatus string

// Enum values for ReviewStatus
const (
	ReviewStatusPending ReviewStatus = "PENDING"
	ReviewStatusFailed  ReviewStatus = "FAILED"
	ReviewStatusGranted ReviewStatus = "GRANTED"
	ReviewStatusDenied  ReviewStatus = "DENIED"
)

// Values returns all known values for ReviewStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReviewStatus) Values() []ReviewStatus {
	return []ReviewStatus{
		"PENDING",
		"FAILED",
		"GRANTED",
		"DENIED",
	}
}

type ScalingMode string

// Enum values for ScalingMode
const (
	ScalingModeStandard ScalingMode = "STANDARD"
	ScalingModeManaged  ScalingMode = "MANAGED"
)

// Values returns all known values for ScalingMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScalingMode) Values() []ScalingMode {
	return []ScalingMode{
		"STANDARD",
		"MANAGED",
	}
}

type SubscriptionStatus string

// Enum values for SubscriptionStatus
const (
	SubscriptionStatusOptIn  SubscriptionStatus = "OPT_IN"
	SubscriptionStatusOptOut SubscriptionStatus = "OPT_OUT"
)

// Values returns all known values for SubscriptionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionStatus) Values() []SubscriptionStatus {
	return []SubscriptionStatus{
		"OPT_IN",
		"OPT_OUT",
	}
}

type SuppressionListImportAction string

// Enum values for SuppressionListImportAction
const (
	SuppressionListImportActionDelete SuppressionListImportAction = "DELETE"
	SuppressionListImportActionPut    SuppressionListImportAction = "PUT"
)

// Values returns all known values for SuppressionListImportAction. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionListImportAction) Values() []SuppressionListImportAction {
	return []SuppressionListImportAction{
		"DELETE",
		"PUT",
	}
}

type SuppressionListReason string

// Enum values for SuppressionListReason
const (
	SuppressionListReasonBounce    SuppressionListReason = "BOUNCE"
	SuppressionListReasonComplaint SuppressionListReason = "COMPLAINT"
)

// Values returns all known values for SuppressionListReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuppressionListReason) Values() []SuppressionListReason {
	return []SuppressionListReason{
		"BOUNCE",
		"COMPLAINT",
	}
}

type TlsPolicy string

// Enum values for TlsPolicy
const (
	TlsPolicyRequire  TlsPolicy = "REQUIRE"
	TlsPolicyOptional TlsPolicy = "OPTIONAL"
)

// Values returns all known values for TlsPolicy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TlsPolicy) Values() []TlsPolicy {
	return []TlsPolicy{
		"REQUIRE",
		"OPTIONAL",
	}
}

type VerificationStatus string

// Enum values for VerificationStatus
const (
	VerificationStatusPending          VerificationStatus = "PENDING"
	VerificationStatusSuccess          VerificationStatus = "SUCCESS"
	VerificationStatusFailed           VerificationStatus = "FAILED"
	VerificationStatusTemporaryFailure VerificationStatus = "TEMPORARY_FAILURE"
	VerificationStatusNotStarted       VerificationStatus = "NOT_STARTED"
)

// Values returns all known values for VerificationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VerificationStatus) Values() []VerificationStatus {
	return []VerificationStatus{
		"PENDING",
		"SUCCESS",
		"FAILED",
		"TEMPORARY_FAILURE",
		"NOT_STARTED",
	}
}

type WarmupStatus string

// Enum values for WarmupStatus
const (
	WarmupStatusInProgress WarmupStatus = "IN_PROGRESS"
	WarmupStatusDone       WarmupStatus = "DONE"
)

// Values returns all known values for WarmupStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (WarmupStatus) Values() []WarmupStatus {
	return []WarmupStatus{
		"IN_PROGRESS",
		"DONE",
	}
}
