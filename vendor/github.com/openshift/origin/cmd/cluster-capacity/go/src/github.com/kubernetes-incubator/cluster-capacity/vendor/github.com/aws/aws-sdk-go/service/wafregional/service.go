// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package wafregional

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// This is the AWS WAF Regional API Reference for using AWS WAF with Elastic
// Load Balancing (ELB) Application Load Balancers. The AWS WAF actions and
// data types listed in the reference are available for protecting Application
// Load Balancers. You can use these actions and data types by means of the
// endpoints listed in AWS Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#waf_region).
// This guide is for developers who need detailed information about the AWS
// WAF API actions, data types, and errors. For detailed information about AWS
// WAF features and an overview of how to use the AWS WAF API, see the AWS WAF
// Developer Guide (http://docs.aws.amazon.com/waf/latest/developerguide/).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28
type WAFRegional struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "waf-regional" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName    // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the WAFRegional client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a WAFRegional client from just a session.
//     svc := wafregional.New(mySession)
//
//     // Create a WAFRegional client with additional configuration
//     svc := wafregional.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *WAFRegional {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *WAFRegional {
	svc := &WAFRegional{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-11-28",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSWAF_Regional_20161128",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a WAFRegional operation and runs any
// custom request initialization.
func (c *WAFRegional) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
