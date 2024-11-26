# {{classname}}

All URIs are relative to *https://api.satstream.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMiningInfo**](MiningApi.md#GetMiningInfo) | **Get** /mining/info | Get mining information
[**GetNetworkHashps**](MiningApi.md#GetNetworkHashps) | **Post** /mining/networkhashps | Get network hash per second

# **GetMiningInfo**
> InlineResponse20025 GetMiningInfo(ctx, )
Get mining information

Returns a json object containing mining-related information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkHashps**
> InlineResponse20013 GetNetworkHashps(ctx, body)
Get network hash per second

Returns the estimated network hashes per second based on the last n blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestsGetNetworkHashPsRequest**](RequestsGetNetworkHashPsRequest.md)| Network hash rate parameters | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
