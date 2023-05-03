// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the cluster configuration of the specified Amazon OpenSearch Service
// domain.
func (c *Client) UpdateDomainConfig(ctx context.Context, params *UpdateDomainConfigInput, optFns ...func(*Options)) (*UpdateDomainConfigOutput, error) {
	if params == nil {
		params = &UpdateDomainConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainConfig", params, optFns, c.addOperationUpdateDomainConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to the UpdateDomain operation.
type UpdateDomainConfigInput struct {

	// The name of the domain that you're updating.
	//
	// This member is required.
	DomainName *string

	// Identity and Access Management (IAM) access policy as a JSON-formatted string.
	AccessPolicies *string

	// Key-value pairs to specify advanced configuration options. The following
	// key-value pairs are supported:
	//   - "rest.action.multi.allow_explicit_index": "true" | "false" - Note the use of
	//   a string rather than a boolean. Specifies whether explicit references to indexes
	//   are allowed inside the body of HTTP requests. If you want to configure access
	//   policies for domain sub-resources, such as specific indexes and domain APIs, you
	//   must disable this property. Default is true.
	//   - "indices.fielddata.cache.size": "80" - Note the use of a string rather than
	//   a boolean. Specifies the percentage of heap space allocated to field data.
	//   Default is unbounded.
	//   - "indices.query.bool.max_clause_count": "1024" - Note the use of a string
	//   rather than a boolean. Specifies the maximum number of clauses allowed in a
	//   Lucene boolean query. Default is 1,024. Queries with more than the permitted
	//   number of clauses result in a TooManyClauses error.
	// For more information, see Advanced cluster parameters (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options)
	// .
	AdvancedOptions map[string]string

	// Options for fine-grained access control.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Options for Auto-Tune.
	AutoTuneOptions *types.AutoTuneOptions

	// Changes that you want to make to the cluster configuration, such as the
	// instance type and number of EC2 instances.
	ClusterConfig *types.ClusterConfig

	// Key-value pairs to configure Amazon Cognito authentication for OpenSearch
	// Dashboards.
	CognitoOptions *types.CognitoOptions

	// Additional options for the domain endpoint, such as whether to require HTTPS
	// for all traffic.
	DomainEndpointOptions *types.DomainEndpointOptions

	// This flag, when set to True, specifies whether the UpdateDomain request should
	// return the results of a dry run analysis without actually applying the change. A
	// dry run determines what type of deployment the update will cause.
	DryRun *bool

	// The type of dry run to perform.
	//   - Basic only returns the type of deployment (blue/green or dynamic) that the
	//   update will cause.
	//   - Verbose runs an additional check to validate the changes you're making. For
	//   more information, see Validating a domain update (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-configuration-changes#validation-check)
	//   .
	DryRunMode types.DryRunMode

	// The type and size of the EBS volume to attach to instances in the domain.
	EBSOptions *types.EBSOptions

	// Encryption at rest options for the domain.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// Options to publish OpenSearch logs to Amazon CloudWatch Logs.
	LogPublishingOptions map[string]types.LogPublishingOption

	// Node-to-node encryption options for the domain.
	NodeToNodeEncryptionOptions *types.NodeToNodeEncryptionOptions

	// Off-peak window options for the domain.
	OffPeakWindowOptions *types.OffPeakWindowOptions

	// Option to set the time, in UTC format, for the daily automated snapshot.
	// Default value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// Service software update options for the domain.
	SoftwareUpdateOptions *types.SoftwareUpdateOptions

	// Options to specify the subnets and security groups for a VPC endpoint. For more
	// information, see Launching your Amazon OpenSearch Service domains using a VPC (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html)
	// .
	VPCOptions *types.VPCOptions

	noSmithyDocumentSerde
}

// The results of an UpdateDomain request. Contains the status of the domain being
// updated.
type UpdateDomainConfigOutput struct {

	// The status of the updated domain.
	//
	// This member is required.
	DomainConfig *types.DomainConfig

	// The status of the dry run being performed on the domain, if any.
	DryRunProgressStatus *types.DryRunProgressStatus

	// Results of the dry run performed in the update domain request.
	DryRunResults *types.DryRunResults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateDomainConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainConfig(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDomainConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpdateDomainConfig",
	}
}
