# MemoryChunks2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AddressRangeType** | [**AddressRangeType**](AddressRangeType.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**InterleaveSets** | [**[]InterleaveSet**](InterleaveSet.md) | This is the interleave sets for the memory chunk. | [optional] 
**IsMirrorEnabled** | **bool** | Mirror Enabled status. | [optional] 
**IsSpare** | **bool** | Spare enabled status. | [optional] 
**MemoryChunkSizeMiB** | **int32** | Size of the memory chunk measured in mebibytes (MiB). | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Oem extension object. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


