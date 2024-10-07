// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Classic Load Balancer.
//
// You can add listeners, security groups, subnets, and tags when you create your
// load balancer, or you can add them later using CreateLoadBalancerListeners, ApplySecurityGroupsToLoadBalancer, AttachLoadBalancerToSubnets, and AddTags.
//
// To describe your current load balancers, see DescribeLoadBalancers. When you are finished with a
// load balancer, you can delete it using DeleteLoadBalancer.
//
// You can create up to 20 load balancers per region per account. You can request
// an increase for the number of load balancers for your account. For more
// information, see [Limits for Your Classic Load Balancer]in the Classic Load Balancers Guide.
//
// [Limits for Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html
func (c *Client) CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancer", params, optFns, c.addOperationCreateLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancer.
type CreateLoadBalancerInput struct {

	// The listeners.
	//
	// For more information, see [Listeners for Your Classic Load Balancer] in the Classic Load Balancers Guide.
	//
	// [Listeners for Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html
	//
	// This member is required.
	Listeners []types.Listener

	// The name of the load balancer.
	//
	// This name must be unique within your set of load balancers for the region, must
	// have a maximum of 32 characters, must contain only alphanumeric characters or
	// hyphens, and cannot begin or end with a hyphen.
	//
	// This member is required.
	LoadBalancerName *string

	// One or more Availability Zones from the same region as the load balancer.
	//
	// You must specify at least one Availability Zone.
	//
	// You can add more Availability Zones after you create the load balancer using EnableAvailabilityZonesForLoadBalancer.
	AvailabilityZones []string

	// The type of a load balancer. Valid only for load balancers in a VPC.
	//
	// By default, Elastic Load Balancing creates an Internet-facing load balancer
	// with a DNS name that resolves to public IP addresses. For more information about
	// Internet-facing and Internal load balancers, see [Load Balancer Scheme]in the Elastic Load Balancing
	// User Guide.
	//
	// Specify internal to create a load balancer with a DNS name that resolves to
	// private IP addresses.
	//
	// [Load Balancer Scheme]: https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html#load-balancer-scheme
	Scheme *string

	// The IDs of the security groups to assign to the load balancer.
	SecurityGroups []string

	// The IDs of the subnets in your VPC to attach to the load balancer. Specify one
	// subnet per Availability Zone specified in AvailabilityZones .
	Subnets []string

	// A list of tags to assign to the load balancer.
	//
	// For more information about tagging your load balancer, see [Tag Your Classic Load Balancer] in the Classic Load
	// Balancers Guide.
	//
	// [Tag Your Classic Load Balancer]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the output for CreateLoadBalancer.
type CreateLoadBalancerOutput struct {

	// The DNS name of the load balancer.
	DNSName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLoadBalancer"); err != nil {
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
	if err = addOpCreateLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLoadBalancer",
	}
}
