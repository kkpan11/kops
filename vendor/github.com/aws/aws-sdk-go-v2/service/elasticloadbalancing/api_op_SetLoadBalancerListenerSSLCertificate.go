// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the certificate that terminates the specified listener's SSL connections.
// The specified certificate replaces any prior certificate that was used on the
// same load balancer and port.
//
// For more information about updating your SSL certificate, see [Replace the SSL Certificate for Your Load Balancer] in the Classic
// Load Balancers Guide.
//
// [Replace the SSL Certificate for Your Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-update-ssl-cert.html
func (c *Client) SetLoadBalancerListenerSSLCertificate(ctx context.Context, params *SetLoadBalancerListenerSSLCertificateInput, optFns ...func(*Options)) (*SetLoadBalancerListenerSSLCertificateOutput, error) {
	if params == nil {
		params = &SetLoadBalancerListenerSSLCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetLoadBalancerListenerSSLCertificate", params, optFns, c.addOperationSetLoadBalancerListenerSSLCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetLoadBalancerListenerSSLCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The port that uses the specified SSL certificate.
	//
	// This member is required.
	LoadBalancerPort int32

	// The Amazon Resource Name (ARN) of the SSL certificate.
	//
	// This member is required.
	SSLCertificateId *string

	noSmithyDocumentSerde
}

// Contains the output of SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetLoadBalancerListenerSSLCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerListenerSSLCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerListenerSSLCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetLoadBalancerListenerSSLCertificate"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpSetLoadBalancerListenerSSLCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerListenerSSLCertificate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetLoadBalancerListenerSSLCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetLoadBalancerListenerSSLCertificate",
	}
}
