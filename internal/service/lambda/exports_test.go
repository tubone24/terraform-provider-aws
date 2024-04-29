// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambda

// Exports for use in tests only.
var (
	ResourceAlias                     = resourceAlias
	ResourceCodeSigningConfig         = resourceCodeSigningConfig
	ResourceEventSourceMapping        = resourceEventSourceMapping
	ResourceFunction                  = resourceFunction
	ResourceFunctionEventInvokeConfig = resourceFunctionEventInvokeConfig
	ResourceFunctionURL               = resourceFunctionURL
	ResourceInvocation                = resourceInvocation
	ResourceLayerVersion              = resourceLayerVersion

	FindAliasByTwoPartKey                     = findAliasByTwoPartKey
	FindCodeSigningConfigByARN                = findCodeSigningConfigByARN
	FindEventSourceMappingByID                = findEventSourceMappingByID
	FindFunctionByName                        = findFunctionByName
	FindFunctionEventInvokeConfigByTwoPartKey = findFunctionEventInvokeConfigByTwoPartKey
	FindFunctionURLByTwoPartKey               = findFunctionURLByTwoPartKey
	FindLayerVersionByTwoPartKey              = findLayerVersionByTwoPartKey
	FunctionEventInvokeConfigParseResourceID  = functionEventInvokeConfigParseResourceID
	LayerVersionParseResourceID               = layerVersionParseResourceID
	SignerServiceIsAvailable                  = signerServiceIsAvailable
)
