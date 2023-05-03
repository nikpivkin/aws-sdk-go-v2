// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Describes an additional authentication provider.
type AdditionalAuthenticationProvider struct {

	// The authentication type: API key, Identity and Access Management (IAM), OpenID
	// Connect (OIDC), Amazon Cognito user pools, or Lambda.
	AuthenticationType AuthenticationType

	// Configuration for Lambda function authorization.
	LambdaAuthorizerConfig *LambdaAuthorizerConfig

	// The OIDC configuration.
	OpenIDConnectConfig *OpenIDConnectConfig

	// The Amazon Cognito user pool configuration.
	UserPoolConfig *CognitoUserPoolConfig

	noSmithyDocumentSerde
}

// Describes an ApiAssociation object.
type ApiAssociation struct {

	// The API ID.
	ApiId *string

	// Identifies the status of an association.
	//   - PROCESSING: The API association is being created. You cannot modify
	//   association requests during processing.
	//   - SUCCESS: The API association was successful. You can modify associations
	//   after success.
	//   - FAILED: The API association has failed. You can modify associations after
	//   failure.
	AssociationStatus AssociationStatus

	// Details about the last deployment status.
	DeploymentDetail *string

	// The domain name.
	DomainName *string

	noSmithyDocumentSerde
}

// The ApiCache object.
type ApiCache struct {

	// Caching behavior.
	//   - FULL_REQUEST_CACHING: All requests are fully cached.
	//   - PER_RESOLVER_CACHING: Individual resolvers that you specify are cached.
	ApiCachingBehavior ApiCachingBehavior

	// At-rest encryption flag for cache. You cannot update this setting after
	// creation.
	AtRestEncryptionEnabled bool

	// The cache instance status.
	//   - AVAILABLE: The instance is available for use.
	//   - CREATING: The instance is currently creating.
	//   - DELETING: The instance is currently deleting.
	//   - MODIFYING: The instance is currently modifying.
	//   - FAILED: The instance has failed creation.
	Status ApiCacheStatus

	// Transit encryption flag when connecting to cache. You cannot update this
	// setting after creation.
	TransitEncryptionEnabled bool

	// TTL in seconds for cache entries. Valid values are 1–3,600 seconds.
	Ttl int64

	// The cache instance type. Valid values are
	//   - SMALL
	//   - MEDIUM
	//   - LARGE
	//   - XLARGE
	//   - LARGE_2X
	//   - LARGE_4X
	//   - LARGE_8X (not available in all regions)
	//   - LARGE_12X
	// Historically, instance types were identified by an EC2-style value. As of July
	// 2020, this is deprecated, and the generic identifiers above should be used. The
	// following legacy instance types are available, but their use is discouraged:
	//   - T2_SMALL: A t2.small instance type.
	//   - T2_MEDIUM: A t2.medium instance type.
	//   - R4_LARGE: A r4.large instance type.
	//   - R4_XLARGE: A r4.xlarge instance type.
	//   - R4_2XLARGE: A r4.2xlarge instance type.
	//   - R4_4XLARGE: A r4.4xlarge instance type.
	//   - R4_8XLARGE: A r4.8xlarge instance type.
	Type ApiCacheType

	noSmithyDocumentSerde
}

// Describes an API key. Customers invoke AppSync GraphQL API operations with API
// keys as an identity mechanism. There are two key versions: da1: We introduced
// this version at launch in November 2017. These keys always expire after 7 days.
// Amazon DynamoDB TTL manages key expiration. These keys ceased to be valid after
// February 21, 2018, and they should no longer be used.
//   - ListApiKeys returns the expiration time in milliseconds.
//   - CreateApiKey returns the expiration time in milliseconds.
//   - UpdateApiKey is not available for this key version.
//   - DeleteApiKey deletes the item from the table.
//   - Expiration is stored in DynamoDB as milliseconds. This results in a bug
//     where keys are not automatically deleted because DynamoDB expects the TTL to be
//     stored in seconds. As a one-time action, we deleted these keys from the table on
//     February 21, 2018.
//
// da2: We introduced this version in February 2018 when AppSync added support to
// extend key expiration.
//   - ListApiKeys returns the expiration time and deletion time in seconds.
//   - CreateApiKey returns the expiration time and deletion time in seconds and
//     accepts a user-provided expiration time in seconds.
//   - UpdateApiKey returns the expiration time and and deletion time in seconds
//     and accepts a user-provided expiration time in seconds. Expired API keys are
//     kept for 60 days after the expiration time. You can update the key expiration
//     time as long as the key isn't deleted.
//   - DeleteApiKey deletes the item from the table.
//   - Expiration is stored in DynamoDB as seconds. After the expiration time,
//     using the key to authenticate will fail. However, you can reinstate the key
//     before deletion.
//   - Deletion is stored in DynamoDB as seconds. The key is deleted after
//     deletion time.
type ApiKey struct {

	// The time after which the API key is deleted. The date is represented as seconds
	// since the epoch, rounded down to the nearest hour.
	Deletes int64

	// A description of the purpose of the API key.
	Description *string

	// The time after which the API key expires. The date is represented as seconds
	// since the epoch, rounded down to the nearest hour.
	Expires int64

	// The API key ID.
	Id *string

	noSmithyDocumentSerde
}

// Describes a runtime used by an Amazon Web Services AppSync pipeline resolver or
// Amazon Web Services AppSync function. Specifies the name and version of the
// runtime to use. Note that if a runtime is specified, code must also be
// specified.
type AppSyncRuntime struct {

	// The name of the runtime to use. Currently, the only allowed value is APPSYNC_JS .
	//
	// This member is required.
	Name RuntimeName

	// The version of the runtime to use. Currently, the only allowed version is 1.0.0 .
	//
	// This member is required.
	RuntimeVersion *string

	noSmithyDocumentSerde
}

// The authorization configuration in case the HTTP endpoint requires
// authorization.
type AuthorizationConfig struct {

	// The authorization type that the HTTP endpoint requires.
	//   - AWS_IAM: The authorization type is Signature Version 4 (SigV4).
	//
	// This member is required.
	AuthorizationType AuthorizationType

	// The Identity and Access Management (IAM) settings.
	AwsIamConfig *AwsIamConfig

	noSmithyDocumentSerde
}

// The Identity and Access Management (IAM) configuration.
type AwsIamConfig struct {

	// The signing Amazon Web Services Region for IAM authorization.
	SigningRegion *string

	// The signing service name for IAM authorization.
	SigningServiceName *string

	noSmithyDocumentSerde
}

// Provides further details for the reason behind the bad request. For reason type
// CODE_ERROR , the detail will contain a list of code errors.
type BadRequestDetail struct {

	// Contains the list of errors in the request.
	CodeErrors []CodeError

	noSmithyDocumentSerde
}

// The caching configuration for a resolver that has caching activated.
type CachingConfig struct {

	// The TTL in seconds for a resolver that has caching activated. Valid values are
	// 1–3,600 seconds.
	//
	// This member is required.
	Ttl int64

	// The caching keys for a resolver that has caching activated. Valid values are
	// entries from the $context.arguments , $context.source , and $context.identity
	// maps.
	CachingKeys []string

	noSmithyDocumentSerde
}

// Describes an AppSync error.
type CodeError struct {

	// The type of code error. Examples include, but aren't limited to: LINT_ERROR ,
	// PARSER_ERROR .
	ErrorType *string

	// The line, column, and span location of the error in the code.
	Location *CodeErrorLocation

	// A user presentable error. Examples include, but aren't limited to: Parsing
	// error: Unterminated string literal .
	Value *string

	noSmithyDocumentSerde
}

// Describes the location of the error in a code sample.
type CodeErrorLocation struct {

	// The column number in the code. Defaults to 0 if unknown.
	Column int32

	// The line number in the code. Defaults to 0 if unknown.
	Line int32

	// The span/length of the error. Defaults to -1 if unknown.
	Span int32

	noSmithyDocumentSerde
}

// Describes an Amazon Cognito user pool configuration.
type CognitoUserPoolConfig struct {

	// The Amazon Web Services Region in which the user pool was created.
	//
	// This member is required.
	AwsRegion *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// A regular expression for validating the incoming Amazon Cognito user pool app
	// client ID. If this value isn't set, no filtering is applied.
	AppIdClientRegex *string

	noSmithyDocumentSerde
}

// Describes a data source.
type DataSource struct {

	// The data source Amazon Resource Name (ARN).
	DataSourceArn *string

	// The description of the data source.
	Description *string

	// DynamoDB settings.
	DynamodbConfig *DynamodbDataSourceConfig

	// Amazon OpenSearch Service settings.
	ElasticsearchConfig *ElasticsearchDataSourceConfig

	// Amazon EventBridge settings.
	EventBridgeConfig *EventBridgeDataSourceConfig

	// HTTP endpoint settings.
	HttpConfig *HttpDataSourceConfig

	// Lambda settings.
	LambdaConfig *LambdaDataSourceConfig

	// The name of the data source.
	Name *string

	// Amazon OpenSearch Service settings.
	OpenSearchServiceConfig *OpenSearchServiceDataSourceConfig

	// Relational database settings.
	RelationalDatabaseConfig *RelationalDatabaseDataSourceConfig

	// The Identity and Access Management (IAM) service role Amazon Resource Name
	// (ARN) for the data source. The system assumes this role when accessing the data
	// source.
	ServiceRoleArn *string

	// The type of the data source.
	//   - AWS_LAMBDA: The data source is an Lambda function.
	//   - AMAZON_DYNAMODB: The data source is an Amazon DynamoDB table.
	//   - AMAZON_ELASTICSEARCH: The data source is an Amazon OpenSearch Service
	//   domain.
	//   - AMAZON_OPENSEARCH_SERVICE: The data source is an Amazon OpenSearch Service
	//   domain.
	//   - AMAZON_EVENTBRIDGE: The data source is an Amazon EventBridge configuration.
	//   - NONE: There is no data source. Use this type when you want to invoke a
	//   GraphQL operation without connecting to a data source, such as when you're
	//   performing data transformation with resolvers or invoking a subscription from a
	//   mutation.
	//   - HTTP: The data source is an HTTP endpoint.
	//   - RELATIONAL_DATABASE: The data source is a relational database.
	Type DataSourceType

	noSmithyDocumentSerde
}

// Describes a Delta Sync configuration.
type DeltaSyncConfig struct {

	// The number of minutes that an Item is stored in the data source.
	BaseTableTTL int64

	// The Delta Sync table name.
	DeltaSyncTableName *string

	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync
	// table.
	DeltaSyncTableTTL int64

	noSmithyDocumentSerde
}

// Describes a configuration for a custom domain.
type DomainNameConfig struct {

	// The domain name that AppSync provides.
	AppsyncDomainName *string

	// The Amazon Resource Name (ARN) of the certificate. This can be an Certificate
	// Manager (ACM) certificate or an Identity and Access Management (IAM) server
	// certificate.
	CertificateArn *string

	// A description of the DomainName configuration.
	Description *string

	// The domain name.
	DomainName *string

	// The ID of your Amazon Route 53 hosted zone.
	HostedZoneId *string

	noSmithyDocumentSerde
}

// Describes an Amazon DynamoDB data source configuration.
type DynamodbDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The table name.
	//
	// This member is required.
	TableName *string

	// The DeltaSyncConfig for a versioned data source.
	DeltaSyncConfig *DeltaSyncConfig

	// Set to TRUE to use Amazon Cognito credentials with this data source.
	UseCallerCredentials bool

	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	Versioned bool

	noSmithyDocumentSerde
}

// Describes an OpenSearch data source configuration. As of September 2021, Amazon
// Elasticsearch service is Amazon OpenSearch Service. This configuration is
// deprecated. For new data sources, use OpenSearchServiceDataSourceConfig to
// specify an OpenSearch data source.
type ElasticsearchDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The endpoint.
	//
	// This member is required.
	Endpoint *string

	noSmithyDocumentSerde
}

// Contains the list of errors generated. When using JavaScript, this will apply
// to the request or response function evaluation.
type ErrorDetail struct {

	// The error payload.
	Message *string

	noSmithyDocumentSerde
}

// Contains the list of errors from a code evaluation response.
type EvaluateCodeErrorDetail struct {

	// Contains the list of CodeError objects.
	CodeErrors []CodeError

	// The error payload.
	Message *string

	noSmithyDocumentSerde
}

// Describes an Amazon EventBridge bus data source configuration.
type EventBridgeDataSourceConfig struct {

	// The ARN of the event bus. For more information about event buses, see Amazon
	// EventBridge event buses (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-event-bus.html)
	// .
	//
	// This member is required.
	EventBusArn *string

	noSmithyDocumentSerde
}

// A function is a reusable entity. You can use multiple functions to compose the
// resolver logic.
type FunctionConfiguration struct {

	// The function code that contains the request and response functions. When code
	// is used, the runtime is required. The runtime value must be APPSYNC_JS .
	Code *string

	// The name of the DataSource .
	DataSourceName *string

	// The Function description.
	Description *string

	// The Amazon Resource Name (ARN) of the Function object.
	FunctionArn *string

	// A unique ID representing the Function object.
	FunctionId *string

	// The version of the request mapping template. Currently, only the 2018-05-29
	// version of the template is supported.
	FunctionVersion *string

	// The maximum batching size for a resolver.
	MaxBatchSize int32

	// The name of the Function object.
	Name *string

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	RequestMappingTemplate *string

	// The Function response mapping template.
	ResponseMappingTemplate *string

	// Describes a runtime used by an Amazon Web Services AppSync pipeline resolver or
	// Amazon Web Services AppSync function. Specifies the name and version of the
	// runtime to use. Note that if a runtime is specified, code must also be
	// specified.
	Runtime *AppSyncRuntime

	// Describes a Sync configuration for a resolver. Specifies which Conflict
	// Detection strategy and Resolution strategy to use when the resolver is invoked.
	SyncConfig *SyncConfig

	noSmithyDocumentSerde
}

// Describes a GraphQL API.
type GraphqlApi struct {

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []AdditionalAuthenticationProvider

	// The API ID.
	ApiId *string

	// The Amazon Resource Name (ARN).
	Arn *string

	// The authentication type.
	AuthenticationType AuthenticationType

	// The DNS records for the API.
	Dns map[string]string

	// Configuration for Lambda function authorization.
	LambdaAuthorizerConfig *LambdaAuthorizerConfig

	// The Amazon CloudWatch Logs configuration.
	LogConfig *LogConfig

	// The API name.
	Name *string

	// The OpenID Connect configuration.
	OpenIDConnectConfig *OpenIDConnectConfig

	// The tags.
	Tags map[string]string

	// The URIs.
	Uris map[string]string

	// The Amazon Cognito user pool configuration.
	UserPoolConfig *UserPoolConfig

	// Sets the value of the GraphQL API to public ( GLOBAL ) or private ( PRIVATE ).
	// If no value is provided, the visibility will be set to GLOBAL by default. This
	// value cannot be changed once the API has been created.
	Visibility GraphQLApiVisibility

	// The ARN of the WAF access control list (ACL) associated with this GraphqlApi ,
	// if one exists.
	WafWebAclArn *string

	// A flag indicating whether to use X-Ray tracing for this GraphqlApi .
	XrayEnabled bool

	noSmithyDocumentSerde
}

// Describes an HTTP data source configuration.
type HttpDataSourceConfig struct {

	// The authorization configuration in case the HTTP endpoint requires
	// authorization.
	AuthorizationConfig *AuthorizationConfig

	// The HTTP URL endpoint. You can specify either the domain name or IP, and port
	// combination, and the URL scheme must be HTTP or HTTPS. If you don't specify the
	// port, AppSync uses the default port 80 for the HTTP endpoint and port 443 for
	// HTTPS endpoints.
	Endpoint *string

	noSmithyDocumentSerde
}

// A LambdaAuthorizerConfig specifies how to authorize AppSync API access when
// using the AWS_LAMBDA authorizer mode. Be aware that an AppSync API can have
// only one Lambda authorizer configured at a time.
type LambdaAuthorizerConfig struct {

	// The Amazon Resource Name (ARN) of the Lambda function to be called for
	// authorization. This can be a standard Lambda ARN, a version ARN ( .../v3 ), or
	// an alias ARN. Note: This Lambda function must have the following resource-based
	// policy assigned to it. When configuring Lambda authorizers in the console, this
	// is done for you. To use the Command Line Interface (CLI), run the following:
	// aws lambda add-permission --function-name
	// "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id
	// "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction
	//
	// This member is required.
	AuthorizerUri *string

	// The number of seconds a response should be cached for. The default is 0
	// seconds, which disables caching. If you don't specify a value for
	// authorizerResultTtlInSeconds , the default value is used. The maximum value is
	// one hour (3600 seconds). The Lambda function can override this by returning a
	// ttlOverride key in its response.
	AuthorizerResultTtlInSeconds int32

	// A regular expression for validation of tokens before the Lambda function is
	// called.
	IdentityValidationExpression *string

	noSmithyDocumentSerde
}

// The LambdaConflictHandlerConfig object when configuring LAMBDA as the Conflict
// Handler.
type LambdaConflictHandlerConfig struct {

	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict
	// Handler.
	LambdaConflictHandlerArn *string

	noSmithyDocumentSerde
}

// Describes an Lambda data source configuration.
type LambdaDataSourceConfig struct {

	// The Amazon Resource Name (ARN) for the Lambda function.
	//
	// This member is required.
	LambdaFunctionArn *string

	noSmithyDocumentSerde
}

// The Amazon CloudWatch Logs configuration.
type LogConfig struct {

	// The service role that AppSync assumes to publish to CloudWatch logs in your
	// account.
	//
	// This member is required.
	CloudWatchLogsRoleArn *string

	// The field logging level. Values can be NONE, ERROR, or ALL.
	//   - NONE: No field-level logs are captured.
	//   - ERROR: Logs the following information only for the fields that are in
	//   error:
	//   - The error section in the server response.
	//   - Field-level errors.
	//   - The generated request/response functions that got resolved for error
	//   fields.
	//   - ALL: The following information is logged for all fields in the query:
	//   - Field-level tracing information.
	//   - The generated request/response functions that got resolved for each field.
	//
	// This member is required.
	FieldLogLevel FieldLogLevel

	// Set to TRUE to exclude sections that contain information such as headers,
	// context, and evaluated mapping templates, regardless of logging level.
	ExcludeVerboseContent bool

	noSmithyDocumentSerde
}

// Describes an OpenID Connect (OIDC) configuration.
type OpenIDConnectConfig struct {

	// The issuer for the OIDC configuration. The issuer returned by discovery must
	// exactly match the value of iss in the ID token.
	//
	// This member is required.
	Issuer *string

	// The number of milliseconds that a token is valid after being authenticated.
	AuthTTL int64

	// The client identifier of the relying party at the OpenID identity provider.
	// This identifier is typically obtained when the relying party is registered with
	// the OpenID identity provider. You can specify a regular expression so that
	// AppSync can validate against multiple client identifiers at a time.
	ClientId *string

	// The number of milliseconds that a token is valid after it's issued to a user.
	IatTTL int64

	noSmithyDocumentSerde
}

// Describes an OpenSearch data source configuration.
type OpenSearchServiceDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The endpoint.
	//
	// This member is required.
	Endpoint *string

	noSmithyDocumentSerde
}

// The pipeline configuration for a resolver of kind PIPELINE .
type PipelineConfig struct {

	// A list of Function objects.
	Functions []string

	noSmithyDocumentSerde
}

// The Amazon Relational Database Service (Amazon RDS) HTTP endpoint configuration.
type RdsHttpEndpointConfig struct {

	// Amazon Web Services Region for Amazon RDS HTTP endpoint.
	AwsRegion *string

	// Amazon Web Services secret store Amazon Resource Name (ARN) for database
	// credentials.
	AwsSecretStoreArn *string

	// Logical database name.
	DatabaseName *string

	// Amazon RDS cluster Amazon Resource Name (ARN).
	DbClusterIdentifier *string

	// Logical schema name.
	Schema *string

	noSmithyDocumentSerde
}

// Describes a relational database data source configuration.
type RelationalDatabaseDataSourceConfig struct {

	// Amazon RDS HTTP endpoint settings.
	RdsHttpEndpointConfig *RdsHttpEndpointConfig

	// Source type for the relational database.
	//   - RDS_HTTP_ENDPOINT: The relational database source type is an Amazon
	//   Relational Database Service (Amazon RDS) HTTP endpoint.
	RelationalDatabaseSourceType RelationalDatabaseSourceType

	noSmithyDocumentSerde
}

// Describes a resolver.
type Resolver struct {

	// The caching configuration for the resolver.
	CachingConfig *CachingConfig

	// The resolver code that contains the request and response functions. When code
	// is used, the runtime is required. The runtime value must be APPSYNC_JS .
	Code *string

	// The resolver data source name.
	DataSourceName *string

	// The resolver field name.
	FieldName *string

	// The resolver type.
	//   - UNIT: A UNIT resolver type. A UNIT resolver is the default resolver type.
	//   You can use a UNIT resolver to run a GraphQL query against a single data source.
	//
	//   - PIPELINE: A PIPELINE resolver type. You can use a PIPELINE resolver to
	//   invoke a series of Function objects in a serial manner. You can use a pipeline
	//   resolver to run a GraphQL query against multiple data sources.
	Kind ResolverKind

	// The maximum batching size for a resolver.
	MaxBatchSize int32

	// The PipelineConfig .
	PipelineConfig *PipelineConfig

	// The request mapping template.
	RequestMappingTemplate *string

	// The resolver Amazon Resource Name (ARN).
	ResolverArn *string

	// The response mapping template.
	ResponseMappingTemplate *string

	// Describes a runtime used by an Amazon Web Services AppSync pipeline resolver or
	// Amazon Web Services AppSync function. Specifies the name and version of the
	// runtime to use. Note that if a runtime is specified, code must also be
	// specified.
	Runtime *AppSyncRuntime

	// The SyncConfig for a resolver attached to a versioned data source.
	SyncConfig *SyncConfig

	// The resolver type name.
	TypeName *string

	noSmithyDocumentSerde
}

// Describes a Sync configuration for a resolver. Specifies which Conflict
// Detection strategy and Resolution strategy to use when the resolver is invoked.
type SyncConfig struct {

	// The Conflict Detection strategy to use.
	//   - VERSION: Detect conflicts based on object versions for this resolver.
	//   - NONE: Do not detect conflicts when invoking this resolver.
	ConflictDetection ConflictDetectionType

	// The Conflict Resolution strategy to perform in the event of a conflict.
	//   - OPTIMISTIC_CONCURRENCY: Resolve conflicts by rejecting mutations when
	//   versions don't match the latest version at the server.
	//   - AUTOMERGE: Resolve conflicts with the Automerge conflict resolution
	//   strategy.
	//   - LAMBDA: Resolve conflicts with an Lambda function supplied in the
	//   LambdaConflictHandlerConfig .
	ConflictHandler ConflictHandlerType

	// The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler.
	LambdaConflictHandlerConfig *LambdaConflictHandlerConfig

	noSmithyDocumentSerde
}

// Describes a type.
type Type struct {

	// The type Amazon Resource Name (ARN).
	Arn *string

	// The type definition.
	Definition *string

	// The type description.
	Description *string

	// The type format: SDL or JSON.
	Format TypeDefinitionFormat

	// The type name.
	Name *string

	noSmithyDocumentSerde
}

// Describes an Amazon Cognito user pool configuration.
type UserPoolConfig struct {

	// The Amazon Web Services Region in which the user pool was created.
	//
	// This member is required.
	AwsRegion *string

	// The action that you want your GraphQL API to take when a request that uses
	// Amazon Cognito user pool authentication doesn't match the Amazon Cognito user
	// pool configuration.
	//
	// This member is required.
	DefaultAction DefaultAction

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// A regular expression for validating the incoming Amazon Cognito user pool app
	// client ID. If this value isn't set, no filtering is applied.
	AppIdClientRegex *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
