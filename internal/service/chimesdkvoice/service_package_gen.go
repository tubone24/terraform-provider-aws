// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package chimesdkvoice

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	chimesdkvoice_sdkv1 "github.com/aws/aws-sdk-go/service/chimesdkvoice"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceGlobalSettings,
			TypeName: "aws_chimesdkvoice_global_settings",
		},
		{
			Factory:  ResourceSipMediaApplication,
			TypeName: "aws_chimesdkvoice_sip_media_application",
			Name:     "Sip Media Application",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceVoiceProfileDomain,
			TypeName: "aws_chimesdkvoice_voice_profile_domain",
			Name:     "Voice Profile Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ChimeSDKVoice
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*chimesdkvoice_sdkv1.ChimeSDKVoice, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return chimesdkvoice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
