# Processor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Assembly** | [**IdRef**](idRef.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**InstructionSet** | [**InstructionSet**](InstructionSet.md) |  | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Location** | [**Location2**](Location_2.md) |  | [optional] 
**Manufacturer** | **string** | The processor manufacturer. | [optional] 
**MaxSpeedMHz** | **int32** | The maximum clock speed of the processor. | [optional] 
**Model** | **string** | The product model number of this device. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Oem extension object. | [optional] 
**ProcessorArchitecture** | [**ProcessorArchitecture**](ProcessorArchitecture.md) |  | [optional] 
**ProcessorId** | [**ProcessorId**](ProcessorId.md) |  | [optional] 
**ProcessorType** | [**ProcessorType**](ProcessorType.md) |  | [optional] 
**Socket** | **string** | The socket or location of the processor. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**SubProcessors** | [**IdRef**](idRef.md) |  | [optional] 
**TotalCores** | **int32** | The total number of cores contained in this processor. | [optional] 
**TotalThreads** | **int32** | The total number of execution threads supported by this processor. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


