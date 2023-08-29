// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to update the configuration of an existing Amazon FSx file
// system. You can update multiple properties in a single request. For FSx for
// Windows File Server file systems, you can update the following properties:
//   - AuditLogConfiguration
//   - AutomaticBackupRetentionDays
//   - DailyAutomaticBackupStartTime
//   - SelfManagedActiveDirectoryConfiguration
//   - StorageCapacity
//   - StorageType
//   - ThroughputCapacity
//   - DiskIopsConfiguration
//   - WeeklyMaintenanceStartTime
//
// For FSx for Lustre file systems, you can update the following properties:
//   - AutoImportPolicy
//   - AutomaticBackupRetentionDays
//   - DailyAutomaticBackupStartTime
//   - DataCompressionType
//   - LogConfiguration
//   - LustreRootSquashConfiguration
//   - StorageCapacity
//   - WeeklyMaintenanceStartTime
//
// For FSx for ONTAP file systems, you can update the following properties:
//   - AddRouteTableIds
//   - AutomaticBackupRetentionDays
//   - DailyAutomaticBackupStartTime
//   - DiskIopsConfiguration
//   - FsxAdminPassword
//   - RemoveRouteTableIds
//   - StorageCapacity
//   - ThroughputCapacity
//   - WeeklyMaintenanceStartTime
//
// For FSx for OpenZFS file systems, you can update the following properties:
//   - AddRouteTableIds
//   - AutomaticBackupRetentionDays
//   - CopyTagsToBackups
//   - CopyTagsToVolumes
//   - DailyAutomaticBackupStartTime
//   - DiskIopsConfiguration
//   - RemoveRouteTableIds
//   - StorageCapacity
//   - ThroughputCapacity
//   - WeeklyMaintenanceStartTime
func (c *Client) UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) {
	if params == nil {
		params = &UpdateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFileSystem", params, optFns, c.addOperationUpdateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the UpdateFileSystem operation.
type UpdateFileSystemInput struct {

	// The ID of the file system that you are updating.
	//
	// This member is required.
	FileSystemId *string

	// A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent
	// updates. This string is automatically filled on your behalf when you use the
	// Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The configuration object for Amazon FSx for Lustre file systems used in the
	// UpdateFileSystem operation.
	LustreConfiguration *types.UpdateFileSystemLustreConfiguration

	// The configuration updates for an Amazon FSx for NetApp ONTAP file system.
	OntapConfiguration *types.UpdateFileSystemOntapConfiguration

	// The configuration updates for an FSx for OpenZFS file system.
	OpenZFSConfiguration *types.UpdateFileSystemOpenZFSConfiguration

	// Use this parameter to increase the storage capacity of an FSx for Windows File
	// Server, FSx for Lustre, FSx for OpenZFS, or FSx for ONTAP file system. Specifies
	// the storage capacity target value, in GiB, to increase the storage capacity for
	// the file system that you're updating. You can't make a storage capacity increase
	// request if there is an existing storage capacity increase request in progress.
	// For Lustre file systems, the storage capacity target value can be the following:
	//
	//   - For SCRATCH_2 , PERSISTENT_1 , and PERSISTENT_2 SSD deployment types, valid
	//   values are in multiples of 2400 GiB. The value must be greater than the current
	//   storage capacity.
	//   - For PERSISTENT HDD file systems, valid values are multiples of 6000 GiB for
	//   12-MBps throughput per TiB file systems and multiples of 1800 GiB for 40-MBps
	//   throughput per TiB file systems. The values must be greater than the current
	//   storage capacity.
	//   - For SCRATCH_1 file systems, you can't increase the storage capacity.
	// For more information, see Managing storage and throughput capacity (https://docs.aws.amazon.com/fsx/latest/LustreGuide/managing-storage-capacity.html)
	// in the FSx for Lustre User Guide. For FSx for OpenZFS file systems, the storage
	// capacity target value must be at least 10 percent greater than the current
	// storage capacity value. For more information, see Managing storage capacity (https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/managing-storage-capacity.html)
	// in the FSx for OpenZFS User Guide. For Windows file systems, the storage
	// capacity target value must be at least 10 percent greater than the current
	// storage capacity value. To increase storage capacity, the file system must have
	// at least 16 MBps of throughput capacity. For more information, see Managing
	// storage capacity (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-storage-capacity.html)
	// in the Amazon FSxfor Windows File Server User Guide. For ONTAP file systems, the
	// storage capacity target value must be at least 10 percent greater than the
	// current storage capacity value. For more information, see Managing storage
	// capacity and provisioned IOPS (https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-storage-capacity.html)
	// in the Amazon FSx for NetApp ONTAP User Guide.
	StorageCapacity *int32

	// Specifies the file system's storage type.
	StorageType types.StorageType

	// The configuration updates for an Amazon FSx for Windows File Server file system.
	WindowsConfiguration *types.UpdateFileSystemWindowsConfiguration

	noSmithyDocumentSerde
}

// The response object for the UpdateFileSystem operation.
type UpdateFileSystemOutput struct {

	// A description of the file system that was updated.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateFileSystemResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addIdempotencyToken_opUpdateFileSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystem(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpUpdateFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateFileSystemInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateFileSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "UpdateFileSystem",
	}
}

type opUpdateFileSystemResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateFileSystemResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateFileSystemResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "fsx"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "fsx"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("fsx")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateFileSystemResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateFileSystemResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
