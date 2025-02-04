// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This request creates a logical container to where backups may be copied. This
// request includes a name, the Region, the maximum number of retention days, the
// minimum number of retention days, and optionally can include tags and a creator
// request ID. Do not include sensitive data, such as passport numbers, in the name
// of a backup vault.
func (c *Client) CreateLogicallyAirGappedBackupVault(ctx context.Context, params *CreateLogicallyAirGappedBackupVaultInput, optFns ...func(*Options)) (*CreateLogicallyAirGappedBackupVaultOutput, error) {
	if params == nil {
		params = &CreateLogicallyAirGappedBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLogicallyAirGappedBackupVault", params, optFns, c.addOperationCreateLogicallyAirGappedBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLogicallyAirGappedBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLogicallyAirGappedBackupVaultInput struct {

	// This is the name of the vault that is being created.
	//
	// This member is required.
	BackupVaultName *string

	// This is the setting that specifies the maximum retention period that the vault
	// retains its recovery points. If this parameter is not specified, Backup does not
	// enforce a maximum retention period on the recovery points in the vault (allowing
	// indefinite storage). If specified, any backup or copy job to the vault must have
	// a lifecycle policy with a retention period equal to or shorter than the maximum
	// retention period. If the job retention period is longer than that maximum
	// retention period, then the vault fails the backup or copy job, and you should
	// either modify your lifecycle settings or use a different vault.
	//
	// This member is required.
	MaxRetentionDays *int64

	// This setting specifies the minimum retention period that the vault retains its
	// recovery points. If this parameter is not specified, no minimum retention period
	// is enforced. If specified, any backup or copy job to the vault must have a
	// lifecycle policy with a retention period equal to or longer than the minimum
	// retention period. If a job retention period is shorter than that minimum
	// retention period, then the vault fails the backup or copy job, and you should
	// either modify your lifecycle settings or use a different vault.
	//
	// This member is required.
	MinRetentionDays *int64

	// These are the tags that will be included in the newly-created vault.
	BackupVaultTags map[string]string

	// This is the ID of the creation request.
	CreatorRequestId *string

	noSmithyDocumentSerde
}

type CreateLogicallyAirGappedBackupVaultOutput struct {

	// This is the ARN (Amazon Resource Name) of the vault being created.
	BackupVaultArn *string

	// The name of a logical container where backups are stored. Logically air-gapped
	// backup vaults are identified by names that are unique to the account used to
	// create them and the Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	BackupVaultName *string

	// The date and time when the vault was created. This value is in Unix format,
	// Coordinated Universal Time (UTC), and accurate to milliseconds. For example, the
	// value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	CreationDate *time.Time

	// This is the current state of the vault.
	VaultState types.VaultState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLogicallyAirGappedBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLogicallyAirGappedBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLogicallyAirGappedBackupVault{}, middleware.After)
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
	if err = addCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLogicallyAirGappedBackupVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLogicallyAirGappedBackupVault(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLogicallyAirGappedBackupVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "CreateLogicallyAirGappedBackupVault",
	}
}

type opCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "backup"
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
				signingName = "backup"
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
				v4aScheme.SigningName = aws.String("backup")
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

func addCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateLogicallyAirGappedBackupVaultResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
